import Vue from 'vue'
import Vuex from 'vuex'
import auth from './modules/auth'
import user from './modules/user'

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    page_title: "Home"
  },
  modules: {
    auth,
    user
  }
});

export default store;