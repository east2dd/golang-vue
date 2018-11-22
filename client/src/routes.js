import Home from './components/Home.vue';
import Header from './components/Header.vue';
import DashboardPage from './components/dashboard/dashboard.vue'
import SignupPage from './components/auth/signup.vue'
import SigninPage from './components/auth/signin.vue'

const User = resolve => {
  require.ensure(['./components/user/User.vue'], () => {
    resolve(require('./components/user/User.vue'));
  });
};
const UserList = resolve => {
  require.ensure(['./components/user/UserList.vue'], () => {
    resolve(require('./components/user/UserList.vue'));
  });
};
const UserEdit = resolve => {
  require.ensure(['./components/user/UserEdit.vue'], () => {
    resolve(require('./components/user/UserEdit.vue'));
  });
};
const UserDetail = resolve => {
  require.ensure(['./components/user/UserDetail.vue'], () => {
    resolve(require('./components/user/UserDetail.vue'));
  });
};

const Category = resolve => {
  require.ensure(['./components/category/Category.vue'], () => {
    resolve(require('./components/category/Category.vue'));
  });
};
const CategoryList = resolve => {
  require.ensure(['./components/category/CategoryList.vue'], () => {
    resolve(require('./components/category/CategoryList.vue'));
  });
};

const CategoryDetail = resolve => {
  require.ensure(['./components/category/CategoryDetail.vue'], () => {
    resolve(require('./components/category/CategoryDetail.vue'));
  });
};

const CategoryEdit = resolve => {
  require.ensure(['./components/category/CategoryEdit.vue'], () => {
    resolve(require('./components/category/CategoryEdit.vue'));
  });
};

const Product = resolve => {
  require.ensure(['./components/product/Product.vue'], () => {
    resolve(require('./components/product/Product.vue'));
  });
};
const ProductList = resolve => {
  require.ensure(['./components/product/ProductList.vue'], () => {
    resolve(require('./components/product/ProductList.vue'));
  });
};

const ProductDetail = resolve => {
  require.ensure(['./components/product/ProductDetail.vue'], () => {
    resolve(require('./components/product/ProductDetail.vue'));
  });
};

const ProductEdit = resolve => {
  require.ensure(['./components/product/ProductEdit.vue'], () => {
    resolve(require('./components/product/ProductEdit.vue'));
  });
};

const NotFound = resolve => {
  require.ensure(['./components/404.vue'], () => {
    resolve(require('./components/404.vue'));
  });
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
      'header-top': Header
    }, 
    meta: { layout: 'user' }, 
    children: [
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
      'header-top': Header
    },
    meta: { layout: 'user', requiresAuth: true }, 
    children: [
      { path: '', component: CategoryList, name: 'categoryList' },
      {
        path: ':id', component: CategoryDetail, beforeEnter: (to, from, next) => {
          next();
        }
      },
      { path: ':id/edit', component: CategoryEdit, name: 'categoryEdit' }
    ]
  },
  {
    path: '/products', components: {
      default: Product,
      'header-top': Header
    },
    meta: { layout: 'user', requiresAuth: true }, 
    children: [
      { path: '', component: ProductList },
      { path: ':id', component: ProductDetail },
      { path: ':id/edit', component: ProductEdit, name: 'productEdit' }
    ]
  },
  { path: '/redirect-me', redirect: { name: 'home' } },
  { path: '/404', component: NotFound, name: 'notFound' },
  { path: '/signup', component: SignupPage },
  { path: '/signin', component: SigninPage },
  { path: '/dashboard', component: DashboardPage },
  { path: '*', redirec: '/' }
];