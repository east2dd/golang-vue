import VueRouter from 'vue-router';
import { routes } from './routes';
import store from './store'

export const router = new VueRouter({
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
  let isAuthenticated = store.getters.isAuthenticated

  if(to.matched.some(record => record.meta.requiresAuth)) {
      if (!isAuthenticated) {
          next({
              path: '/signin',
              params: { nextUrl: to.fullPath }
          })
      } else {
        next()
      }
  } else {
      next() 
  }
})

export default router