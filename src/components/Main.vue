<template>
  <div class='main'>
    <p>{{data}}</p>
    <input @click="getData" type="button" value="取得">
    <div class="row mr-2 ml-2">
      <div v-for="item in items" :key="item.id" class="col-md-3 col-sm-6 p-2">
        <div class="card p-0 pb-3 shadow">
          <!-- TODO: 取得した画像を表示 <img :src="item.image" class="card-img-top"> -->
          <img src="@/assets/dummy.jpg" class="card-img-top">
          <p class="card-title">{{ item.name }}</p>
          <p class="card-text font-weight-bolder">¥{{ item.price }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// TODO: 良い感じにクローリングしてデータを成型する関数かAPIを作る
function getItems(){
  return [
    { id: 1, name: "商品名", price: 1000, image: "画像URL" },
    { id: 2, name: "商品名", price: 1000, image: "画像URL" },
    { id: 3, name: "商品名", price: 1000, image: "画像URL" },
    { id: 4, name: "商品名", price: 1000, image: "画像URL" },
    { id: 5, name: "商品名", price: 1000, image: "画像URL" },
    { id: 6, name: "商品名", price: 1000, image: "画像URL" },
    { id: 7, name: "商品名", price: 1000, image: "画像URL" },
    { id: 8, name: "商品名", price: 1000, image: "画像URL" }
  ]
}

let items = getItems();
export default {
  name: 'Main',
  data: function (){
    return {
      items: items,
      data: "test",
      searchPath: "https://google.com"
    }
  },
  methods: {
    getData(){
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
        this.data = response.data;
      }).catch((e) => {
        alert(e);
      });
    }
  }
}
</script>
