import mojoapi from './api'
import Vue from 'vue'
import VueClipboard from 'vue-clipboard2'
import router from './router'
import App from './App.vue'
import store from './store'
import echarts from "echarts"
import "bulma/css/bulma.css"

Vue.config.productionTip = false
Vue.use(VueClipboard)
Vue.prototype.$echarts = echarts
Vue.prototype.$mojoapi = mojoapi
new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')

