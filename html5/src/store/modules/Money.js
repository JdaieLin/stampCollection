import apiHost from '../host'
import Vue from 'vue'

const state = {
  coins: 0,
  ingots: 0,
  chests: [
    {
      hour: 0,
      ingots: 0,
      opened: true
    }
  ]
}

// http post : /api/chest/list
// http post : /api/chest/sync

const mutations = {
  INCREMENT_COINS (state, payload) {
    if (state.coins + payload > 999999) {
      state.coins = 999999
    } else {
      state.coins += payload
    }
  },
  REDUCE_COINS (state, payload) {
    state.coins -= payload
  },
  UPDATE_COINS (state, coins) {
    if (coins > 999999) {
      coins = 999999
    }
    state.coins = coins
  },
  INCREMENT_INGOTS (state, ingots) {
    state.ingots += ingots
  },
  UPDATE_INGOTS (state, ingots) {
    state.ingots = ingots
  },
  UPDATE_CHESTS (state, chests) {
    state.chests = chests
  }
}

const actions = {
  refreshCoin ({commit}) {
    setTimeout(function () {
      commit('INCREMENT_COINS', 1)
    }, 500)
    setTimeout(function () {
      commit('REDUCE_COINS', 1)
    }, 501)
  },
  coinIncrease ({ commit }, payload) {
    commit('INCREMENT_COINS', payload)
  },
  ingotIncrease ({ commit }, payload) {
    commit('INCREMENT_INGOTS', payload)
  },
  listChests ({commit, rootState}, payload) {
    let data = {
      user_id: rootState.User.user_id
    }
    Vue.http.post(apiHost + '/api/chest/list', data).then((response) => {
      if (response.data.success) {
        commit('UPDATE_CHESTS', response.data.result.chests)
        commit('UPDATE_INGOTS', response.data.result.ingots)
      }
    })
  },
  syncChest ({commit, rootState}, payload) {
    let data = {
      user_id: rootState.User.user_id,
      chests: [payload]
    }
    Vue.http.post(apiHost + '/api/chest/sync', data).then((response) => {
      console.log(response.data)
    })
  }
}

export default {
  state,
  mutations,
  actions
}
