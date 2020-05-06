<template>
  <div class='main'>
    <div class="row mr-2 ml-2">
      <div class="input-group col-xl-3 col-md-6 p-2">
        <input type="text" class="form-control" v-model="keyword">
        <div class="input-group-append">
          <button type="button" class="btn btn-outline-success" @click="getItems">Search</button>
        </div>
      </div>
    </div>
    <Loading :isLoading="isLoading"/>
    <div class="col-12">
      <div class="alert alert-secondary" v-if="items.length == 0">
        <p class="m-0" >一致する商品はありません</p>
        <p class="m-0" v-if="searched.length">キーワード: {{searched}}</p>
      </div>
    </div>
    <div class="row mr-2 ml-2">
      <div v-for="item in items" :key="item.id" class="col-md-3 col-sm-6 col-6 p-2">
        <a :href="item.Url" target="_blank">
          <div class="card p-0 pb-3 shadow">
            <img :src="item.Image" class="card-img-top">
            <p class="card-title">{{ item.Name }}</p>
            <p class="card-text font-weight-bolder">{{ item.Price }}</p>
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import Loading from './Loading.vue'
export default {
  name: 'Main',
  components: {
    Loading
  },
  data: function (){
    return {
      items: [],
      keyword: "",
      searched: "",
      searchPath: "/api/digimart?keyword=",
      isLoading: false
    }
  },
  methods: {
    getItems(){
      if(this.isLoading) { return }
      this.isLoading = true;
      this.axios.get(
        `${this.searchPath}${this.keyword}`, {
          mode: 'no-cors',
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          },
          withCredentials: true,
          credentials: 'same-origin'
        }
      ).then((response) => {
        this.searched = this.keyword;
        this.items = response.data;
      }).catch((e) => {
        alert(e);
      }).finally(() => {
        this.isLoading = false;
      });
    }
  }
}
</script>
