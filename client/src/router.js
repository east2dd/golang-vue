import VueRouter from 'vue-router';
import { routes } from './routes';

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
  let jwt = $cookies.get('jwt') || localStorage.getItem('jwt');

  if(to.matched.some(record => record.meta.requiresAuth)) {
      if (jwt == null) {
          next({
              path: '/signin',
              params: { nextUrl: to.fullPath }
          })
      } else {
          let user = JSON.parse(localStorage.getItem('user'))
          if(to.matched.some(record => record.meta.is_admin)) {
              if(user.is_admin == 1){
                  next()
              }
              else{
                  next({ name: 'categories'})
              }
          }else {
              next()
          }
      }
  } else if(to.matched.some(record => record.meta.guest)) {
      if(jwt == null){
          next()
      }
      else{
          next({ name: 'categories'})
      }
  }else {
      next() 
  }
})

export default router