<template>
  <a :href="item.Url" target="_blank">
    <div class="card p-0 pb-3 shadow">
      <img :src="item.Image" class="card-img-top">
      <p class="card-title">{{ item.Name }}</p>
      <p class="card-text font-weight-bolder">{{ item.Price }}</p>
    </div>
  </a>
</template>
<script>
export default {
  name: 'ItemCard',
  props: {
    item: Object
  },
}
</script>
  
