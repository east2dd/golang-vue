import Vue from 'vue'
import VueRouter from 'vue-router';
import App from './App.vue'
import { routes } from './routes';
import { store } from './store/store';
import VueMaterial from 'vue-material'

import DefaultLayout from './layouts/Default'
import UserLayout from './layouts/User'

Vue.use(VueMaterial)
Vue.use(VueRouter);

Vue.component('default-layout', DefaultLayout)
Vue.component('user-layout', UserLayout)

const router = new VueRouter({
  routes,
  mode: 'history',
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    if (to.hash) {
      return { selector: to.hash };
    }
    return {x: 0, y: 0};
  }
});

router.beforeEach((to, from, next) => {
  next();
});

new Vue({
  store: store,
  router: router,
  render: h => h(App)
}).$mount('#app')