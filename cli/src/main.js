import httpc from './api'
import Vue from 'vue'
import router from './router'
import App from './App.vue'
import store from './store'
import echarts from "echarts"
import "bulma/css/bulma.css"

Vue.config.productionTip = false
Vue.prototype.$echarts = echarts
Vue.prototype.$httpc = httpc

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')

