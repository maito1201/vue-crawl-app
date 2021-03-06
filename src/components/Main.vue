<template>
  <div class="main" style="height: 80vh">
    <div v-for="condition in conditions" :key="condition.id">
      <div class="row mr-2 ml-2" v-if="current == condition.id">
        <div class="input-group col-xl-3 col-md-6 p-2">
          <select class="form-control" v-model="condition.searchPath">
            <option v-for="option in selectOptions" v-bind:key="option.text" v-bind:value="option.value">
              {{ option.text }}
            </option>
          </select>
          <input type="text" class="form-control" v-model="condition.keyword">
          <div class="input-group-append">
            <button type="button" class="btn btn-outline-success" @click="getItems(condition)">Search</button>
          </div>
        </div>
      </div>
    </div>
    <carousel
      :per-page="1"
      :mouse-drag="false"
      :minSwipeDistance="50"
      :paginationPosition="'top'"
      :paginationPadding="2"
      :paginationSize="8"
      ref="carousel"
      v-touch:swipe.bottom="reloadItems"
      @page-change="updateCurrent"
    >
      <slide v-for="condition in conditions" :key="condition.id">
        <Loading :isLoading="condition.isLoading"/>
        <div style="height: 80vh; overflow: scroll" ref="itemBody">
          <div class="col-12" v-if="condition.items.length == 0">
            <div class="alert alert-secondary">
              <p class="m-0" >一致する商品はありません</p>
              <p class="m-0" v-if="condition.searched.length">キーワード: {{condition.searched}}</p>
            </div>
          </div>
          <div class="row mr-2 ml-2">
            <div v-for="item in condition.items" :key="item.id" class="col-md-3 col-sm-6 col-6 p-2">
              <ItemCard :item="item"/>
            </div>
          </div>
        </div>
      </slide>
    </carousel>
  </div>
</template>

<script>
import ItemCard from "./ItemCard.vue"
import Loading from "./Loading.vue"
import { Carousel, Slide } from "vue-carousel";
export default {
  name: "Main",
  components: {
    ItemCard,
    Loading,
    Carousel,
    Slide
  },
  data: function (){
    return {
      current: 0,
      storedConditions: [],
      conditions: [
        {
          id: 0,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword",
          isLoading: false
        },
        {
          id: 1,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false
        },
        {
          id: 2,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false
        },
        {
          id: 3,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false
        },
        {
          id: 4,
          items: [],
          keyword: "",
          searched: "",
          searchPath: "/api/digimart?keyword=",
          isLoading: false
        }
      ],
      selectOptions: [
        { text: "ヤフオク", value: "/api/yahoo?keyword=" },
        { text: "メルカリ", value: "/api/mercari?keyword=" },
        { text: "デジマート", value: "/api/digimart?keyword=" },
        { text: "reverb", value: "/api/reverb?keyword=" }
      ]
    }
  },
  methods: {
    // アイテム読み込み
    getItems(condition, loaded){
      if (typeof(loaded) == "function") { loaded("done"); }
      if(condition.isLoading) { return }
      condition.isLoading = true;
      this.axios.get(
        `${condition.searchPath}${condition.keyword}`, {
          mode: "no-cors",
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
          },
          withCredentials: true,
          credentials: "same-origin"
        }
      ).then((response) => {
        this.setDatas(response, condition)
      }).catch((e) => {
        alert(e);
      }).finally(() => {
        condition.isLoading = false;
      });
    },
    // 下にスワイプしてリロード
    reloadItems(loaded){
      const index = this.currentCarousel();
      if(this.$refs.itemBody[index].scrollTop > 30) { return }
      this.getItems(this.conditions[index], loaded);
    },
    // ページ再訪、リロード時にcookieの情報をもとに再検索
    resumeItems(loaded){
      const func = this.getItems;
      this.conditions.forEach(function(condition){
        if(condition.keyword != "" && condition.searchPath != ""){
          func(condition, loaded);
        }
      });
    },
    setDatas(response, condition){
      condition.searched = condition.keyword;
      condition.items = response.data;
      this.storedConditions = this.conditions.map(function(condition) {
        return { keyword : condition.keyword, searchPath : condition.searchPath }
      });
      localStorage.setItem("storedConditions", JSON.stringify(this.storedConditions));
    },
    currentCarousel(){
      return this.$refs.carousel ? this.$refs.carousel.currentPage : 0;
    },
    updateCurrent(){
      this.current = this.currentCarousel();
    }
  },
  mounted() {
    if (localStorage.getItem("storedConditions")) {
      try {
        this.storedConditions = JSON.parse(localStorage.getItem("storedConditions"));
        for( var i=0; i < this.conditions.length; i++) {
          this.conditions[i].keyword = this.storedConditions[i].keyword;
          this.conditions[i].searchPath = this.storedConditions[i].searchPath;
        }
        this.resumeItems();
      } catch(e) {
        localStorage.removeItem("storedConditions");
      }
    }
  }
}
</script>
