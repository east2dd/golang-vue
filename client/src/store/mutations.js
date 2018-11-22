import * as types from './types';

export default {
    [types.MUTATE_UPDATE_PAGE_TITLE]: (state, payload) => {
        state.page_title = payload;
    }
};