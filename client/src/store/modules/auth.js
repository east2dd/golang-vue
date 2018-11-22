import * as types from '../types';

const state = {
    user: null
};

const getters = {
    [types.USER]: state => {
        return state.user
    },
    isLoggedIn: state => {
        return state.user != null
    }
};

const mutations = {
    [types.MUTATE_UPDATE_USER]: (state, payload) => {
        state.user = payload;
    }
};

const actions = {
    [types.UPDATE_USER]: ({ commit }, payload) => {
        commit(types.MUTATE_UPDATE_USER, payload);
    },
    [types.LOGIN_USER]: ({ commit }, payload) => {
        $cookies.set("jwt", payload.token)
        $cookies.set("user", payload)
        commit(types.MUTATE_UPDATE_USER, payload);
    },
    [types.LOGOUT_USER]: ({ commit }, payload) => {
        $cookies.remove("jwt")
        $cookies.remove("user")
        commit(types.MUTATE_UPDATE_USER, null);
    },

};

export default {
    state,
    mutations,
    actions,
    getters
}