import * as types from './types';

export default {
    [types.UPDATE_PAGE_TITLE]: ({commit}, payload) => {
        commit(types.MUTATE_UPDATE_PAGE_TITLE, payload)
    }
};