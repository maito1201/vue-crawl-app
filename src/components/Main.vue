<template>
  <div class='main'>
    <h1>Hello My Vue</h1>
    <p class="text-danger">{{ msg }}</p>
  </div>
</template>

<script>
export default {
  name: 'Main',
  props: {
    msg: String
  }
}
</script>
