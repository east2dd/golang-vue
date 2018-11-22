import Vue from 'vue';
import Vuex from 'vuex';
import counter from './modules/counter';

import actions from './actions';
import getters from './getters';
import mutations from './mutations';

Vue.use(Vuex);

export const store = new Vuex.Store({
    state: {
        page_title: "Home"
    },
    getters,
    mutations,
    actions,
    modules: {
        counter
    }
});