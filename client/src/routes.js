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

const Category = resolve => {
  require.ensure(['./components/category/Category.vue'], () => {
    resolve(require('./components/category/Category.vue'));
  }, 'category');
};
const CategoryList = resolve => {
  require.ensure(['./components/category/CategoryList.vue'], () => {
    resolve(require('./components/category/CategoryList.vue'));
  }, 'category');
};

const CategoryDetail = resolve => {
  require.ensure(['./components/category/CategoryDetail.vue'], () => {
    resolve(require('./components/category/CategoryDetail.vue'));
  }, 'category');
};

const CategoryEdit = resolve => {
  require.ensure(['./components/category/CategoryEdit.vue'], () => {
    resolve(require('./components/category/CategoryEdit.vue'));
  }, 'category');
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
  {
    path: '/categories', components: {
      default: Category,
      'header-bottom': Header
    }, children: [
      { path: '', component: CategoryList },
      {
        path: ':id', component: CategoryDetail, beforeEnter: (to, from, next) => {
          next();
        }
      },
      { path: ':id/edit', component: CategoryEdit, name: 'categoryEdit' }
    ]
  },
  { path: '/redirect-me', redirect: { name: 'home' } },
  { path: '*', redirec: '/' }
];