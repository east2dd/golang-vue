import * as types from '../types';

const state = {
    counter: 0
};

const getters = {
    [types.DOUBLE_COUNTER]: state => {
        return state.counter * 2;
    }
};

const mutations = {
    [types.MUTATE_INCREMENT_COUNTER]: (state, payload) => {
        state.counter += payload;
    }
};

const actions = {
    [types.COUNTER_INCREMENT]: ({ commit }, payload) => {
        commit(types.MUTATE_INCREMENT_COUNTER, payload);
    }
};

export default {
    state,
    mutations,
    actions,
    getters
}