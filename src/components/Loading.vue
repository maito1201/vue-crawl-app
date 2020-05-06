 <template>
    <div v-if="isLoading">
       <h4>loading...</h4>
       <vue-loading type="bubbles" color="#28a745" :size="{ width: '50px', height: '50px' }"></vue-loading>    
    </div>
 </template>
  <script>
  import { VueLoading } from 'vue-loading-template'
  export default {
    name: 'Loading',
    components: {
      VueLoading
    },
    props: {
      isLoading: Boolean
    },
  }
  </script>
  