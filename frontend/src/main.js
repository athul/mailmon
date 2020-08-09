import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueFormulate from '@braid/vue-formulate'


Vue.use(VueFormulate)
Vue.config.productionTip = false

Vue.prototype.$http = axios

new Vue({
  render: h => h(App),
}).$mount('#app')
