import * as authActions from '../actions/auth'
import * as userActions from '../actions/user'
import axios from '../../axios-auth'

const state = {
  token: window.$cookies.get('user-token') || '',
  status: '',
}

const getters = {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status,
};

const mutations = {
  [authActions.AUTH_REQUEST]: (state) => {
    state.status = 'loading'
  },
  [authActions.AUTH_SUCCESS]: (state, resp) => {
    state.status = 'success'
    state.token = resp
    state.hasLoadedOnce = true
  },
  [authActions.AUTH_ERROR]: (state) => {
    state.status = 'error'
    state.hasLoadedOnce = true
  },
  [authActions.AUTH_LOGOUT]: (state) => {
    state.token = ''
  }
};

const actions = {
  [authActions.AUTH_REQUEST]: ({commit, dispatch}, user) => {
    return new Promise((resolve, reject) => {
      commit(authActions.AUTH_REQUEST)
      axios({url: '/api/user/login', data: user, method: 'POST' })
      .then(resp => {
          const token = resp.data.data.Token
          window.$cookies.set('user-token', token)
          axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
          // Add the following line:
          commit(authActions.AUTH_SUCCESS, token)
          dispatch(userActions.USER_REQUEST)
          resolve(resp)
      })
      .catch(err => {
          commit(authActions.AUTH_ERROR, err)
          window.$cookies.remove('user-token')
          reject(err)
      })
    })
  },
  [authActions.AUTH_SIGNUP]: ({commit, dispatch}, user) => {
    return new Promise((resolve, reject) => {
      commit(authActions.AUTH_REQUEST)
      axios({url: '/api/user/new', data: user, method: 'POST' })
      .then(resp => {
          const token = resp.data.data.Token
          window.$cookies.set('user-token', token)
          axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
          // Add the following line:
          commit(authActions.AUTH_SUCCESS, token)
          dispatch(userActions.USER_REQUEST)
          resolve(resp)
      })
      .catch(err => {
          commit(authActions.AUTH_ERROR, err)
          window.$cookies.remove('user-token')
          reject(err)
      })
    })
  },
  [authActions.AUTH_LOGOUT]: ({commit, dispatch}) => {
    return new Promise((resolve, reject) => {
        commit(authActions.AUTH_LOGOUT)
        window.$cookies.remove('user-token')
        resolve()
    })
  }
};

export default {
    state,
    mutations,
    actions,
    getters
}