// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuex from 'vuex'
import store from './store'
import axios from 'axios'

Vue.config.productionTip = false
Vue.http = Vue.prototype.$http = axios

Vue.use(Vuex)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})

// 登录测试账号
store.dispatch('login', {
  login_id: 'balabala',
  login_password: 'aabbcc'
})

document.body.addEventListener('touchstart', function () { })
