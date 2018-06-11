import apiHost from '../host'
import Vue from 'vue'

const state = {
  name: '',
  phone: '',
  mail: '',
  user_id: 0,
  login_id: '',
  login_password: '',
  network: 0
}

// http post : /api/user/[register/update]

const mutations = {
  LOGIN (state, userInfo) {
    state.user_id = userInfo.id
    state.name = userInfo.name
    state.phone = userInfo.phone
    state.mail = userInfo.mail
    state.login_id = userInfo.login_id
    state.login_password = userInfo.login_password
  }
}

const actions = {
  login ({ commit, dispatch }, payload) {
    Vue.http.post(apiHost + '/api/user/login', payload).then((response) => {
      if (response.data.success) {
        let data = response.data.result
        commit('LOGIN', data)
        // get money
        commit('UPDATE_COINS', data.coins)
        commit('UPDATE_INGOTS', data.ingots)
        // get chests
        dispatch('listChests')
        // get slide
        dispatch('getLoopSlideInit')
        // get album
        dispatch('getAlbum')
      }
    })
  }
}

export default {
  state,
  mutations,
  actions
}
