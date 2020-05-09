<template>
  <a :href="item.Url" target="_blank">
    <div class="card p-0 pb-3 shadow">
      <img :src="item.Image" class="card-img-top">
      <p class="card-title">{{ item.Name | truncate }}</p>
      <p class="card-text font-weight-bolder">{{ item.Price }}</p>
    </div>
  </a>
</template>

<script>
String.prototype.bytes = function () {
  var length = 0;
  for (var i = 0; i < this.length; i++) {
    var c = this.charCodeAt(i);
    if ((c >= 0x0 && c < 0x81) || (c === 0xf8f0) || (c >= 0xff61 && c < 0xffa0) || (c >= 0xf8f1 && c < 0xf8f4)) {
      length += 1;
    } else {
      length += 2;
    }
  }
  return length;
};

export default {
  name: 'ItemCard',
  props: {
    item: Object
  },
  filters: {
    truncate: function (str) {
      const len = 20;
      if (!str) return ''
      let truncated = str.substr(0, len);
      if (truncated.bytes() > 23) truncated = str.substr(0, 12)
      return str.length <= len ? str : (truncated + "...");
    }
  }
}
</script>

<style scoped>
.card-title{
  font-size: 12px;
  color: black;
}
.card-text{
  color: #B12704;
}
.card-img-top{
  height: 161.5px;
}
</style>
