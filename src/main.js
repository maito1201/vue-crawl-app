import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vue2TouchEvents from 'vue2-touch-events'

Vue.config.productionTip = false
Vue.use(
  VueAxios,
  axios
)
Vue.use(Vue2TouchEvents, {swipeTolerance: 60})

new Vue({
  render: h => h(App),
}).$mount('#app')
