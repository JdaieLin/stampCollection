import Vue from 'vue'
import apiHost from '../host'

const state = {
  tradeList: []
}

// http post : /api/deal/list
// http post : /api/deal/create
// http post : /api/deal/accept

const mutations = {
  UPDATE_TRADE_LIST (state, newList) {
    state.tradeList = newList
  }
}

const actions = {
  getTrade ({ commit, rootState }) {
    Vue.http.post(apiHost + '/api/deal/list', {
      user_id: rootState.User.user_id
    }).then((response) => {
      if (response.data.success) {
        console.log(response.data)
        let tradeList = response.data.result.deals
        if (tradeList) {
          commit('UPDATE_TRADE_LIST', tradeList)
        }
      }
    })
  }
}

export default {
  state,
  mutations,
  actions
}
