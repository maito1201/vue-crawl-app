<template>
  <div class='main'>
    <input @click="getItems" type="button" value="取得">
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

export default {
  name: 'Main',
  data: function (){
    return {
      items: [],
      searchPath: "/api/digimart?keyword=dingwall"
    }
  },
  methods: {
    getItems(){
      this.axios.get(
        this.searchPath, {
          mode: 'no-cors',
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          },
          withCredentials: true,
          credentials: 'same-origin'
        }
      ).then((response) => {
        this.items = response.data;
      }).catch((e) => {
        alert(e);
      });
    }
  }
}
</script>
