import Vue from 'vue';
import Vuex from 'vuex';
import auth from './modules/auth';
import actions from './actions';
import getters from './getters';
import mutations from './mutations';
import * as types from './types';

Vue.use(Vuex);

export const store = new Vuex.Store({
    state: {
        page_title: "Home"
    },
    getters,
    mutations,
    actions,
    modules: {
        auth
    }
});

store.dispatch(types.UPDATE_USER, $cookies.get('user'))