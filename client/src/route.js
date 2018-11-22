import Home from './components/Home.vue';
import Header from './components/Header.vue';

const User = resolve => {
  require.ensure(['./components/user/User.vue'], () => {
    resolve(require('./components/user/User.vue'));
  }, 'user');
};
const UserList = resolve => {
  require.ensure(['./components/user/UserList.vue'], () => {
    resolve(require('./components/user/UserList.vue'));
  }, 'user');
};
const UserEdit = resolve => {
  require.ensure(['./components/user/UserEdit.vue'], () => {
    resolve(require('./components/user/UserEdit.vue'));
  }, 'user');
};
const UserDetail = resolve => {
  require.ensure(['./components/user/UserDetail.vue'], () => {
    resolve(require('./components/user/UserDetail.vue'));
  }, 'user');
};

export const routes = [
  {
    path: '', name: 'home', components: {
      default: Home,
      'header-top': Header
    }
  },
  {
    path: '/users', components: {
      default: User,
      'header-bottom': Header
    }, children: [
      { path: '', component: UserList },
      {
        path: ':id', component: UserDetail, beforeEnter: (to, from, next) => {
          next();
        }
      },
      { path: ':id/edit', component: UserEdit, name: 'userEdit' }
    ]
  },
  { path: '/redirect-me', redirect: { name: 'home' } },
  { path: '*', redirect: '/' }
];