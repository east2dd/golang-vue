import Vue from 'vue'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import VueMaterial from 'vue-material'
import DefaultLayout from './layouts/Default'
import UserLayout from './layouts/User'
import VueCookies from 'vue-cookies'
import App from './App.vue'
import { router } from './router'
import store from './store'

Vue.use(VueCookies);
Vue.use(VueMaterial);
Vue.use(VueRouter);
Vue.use(Vuex);

Vue.component('default-layout', DefaultLayout);
Vue.component('user-layout', UserLayout);

new Vue({
  store: store,
  router: router,
  render: h => h(App)
}).$mount('#app')