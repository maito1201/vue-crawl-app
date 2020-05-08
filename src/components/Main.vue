<template>
  <div class='main'>
    <carousel
      :per-page="1"
      :mouse-drag="false"
      :minSwipeDistance="50"
      :paginationPosition="'top'"
      :paginationPadding="2"
      :paginationSize="8"
      ref="carousel"
      v-touch:swipe.bottom="reloadItems"
    >
      <slide v-for="condition in conditions" :key="condition.id">
        <Loading :isLoading="condition.isLoading"/>
        <div class="row mr-2 ml-2">
          <div class="input-group col-xl-3 col-md-6 p-2">
            <input type="text" class="form-control" v-model="condition.keyword">
            <div class="input-group-append">
              <button type="button" class="btn btn-outline-success" @click="getItems(condition)">Search</button>
            </div>
          </div>
        </div>
        <div class="col-12">
          <div class="alert alert-secondary" v-if="condition.items.length == 0">
            <p class="m-0" >一致する商品はありません</p>
            <p class="m-0" v-if="condition.searched.length">キーワード: {{condition.searched}}</p>
          </div>
        </div>
        <div class="row mr-2 ml-2">
          <div v-for="item in condition.items" :key="item.id" class="col-md-3 col-sm-6 col-6 p-2">
            <ItemCard :item="item"/>
          </div>
        </div>
      </slide>
    </carousel>
  </div>
</template>

<script>
import ItemCard from './ItemCard.vue'
import Loading from './Loading.vue'
import { Carousel, Slide } from 'vue-carousel';
export default {
  name: 'Main',
  components: {
    ItemCard,
    Loading,
    Carousel,
    Slide
  },
  data: function (){
    return {
      conditions: [
        {
          id: 0,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false,
          cssClass: "carousel-item active"
        },
        {
          id: 1,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false,
          cssClass: "carousel-item"
        },
        {
          id: 3,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false,
          cssClass: "carousel-item"
        },
        {
          id: 4,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false,
          cssClass: "carousel-item"
        },
        {
          id: 5,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false,
          cssClass: "carousel-item"
        }
      ],
      isLoading: false
    }
  },
  methods: {
    getItems(condition, loaded){
      if (typeof(loaded) == "function") { loaded('done'); }
      if(condition.isLoading) { return }
      condition.isLoading = true;
      this.axios.get(
        `${condition.searchPath}${condition.keyword}`, {
          mode: 'no-cors',
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
          },
          withCredentials: true,
          credentials: 'same-origin'
        }
      ).then((response) => {
        condition.searched = condition.keyword;
        condition.items = response.data;
      }).catch((e) => {
        alert(e);
      }).finally(() => {
        condition.isLoading = false;
      });
    },
    reloadItems(loaded){
      if(window.scrollY != 0) { return }
      const index = this.$refs.carousel ? this.$refs.carousel.currentPage : 0;
      this.getItems(this.conditions[index], loaded);
    },
    Test(){
      let val = window.scrollY;
      console.log(val)
    }
  }
}
</script>
