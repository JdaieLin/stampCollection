import Vue from 'vue'
import apiHost from '../host'

const state = {
  tradeList: [],
  sellListSingle: [],
  sellListAlbum: [],
  buyListSingle: [],
  buyListAlbum: [],
  hotStamps: []
}

// http post : /api/deal/list
// http post : /api/deal/create
// http post : /api/deal/accept

const mutations = {
  UPDATE_TRADE_LIST (state, { tradeList, userId }) {
    tradeList = tradeList.filter(i => i.stamps).sort((a, b) => b.atlas_id - a.atlas_id).filter(i => i.status === 1)
    state.tradeList = tradeList
    state.sellListSingle = tradeList.filter(i => i.create_user_id === userId && i.stamps.length === 1)
    state.sellListAlbum = tradeList.filter(i => i.create_user_id === userId && i.stamps.length > 1)
    state.buyListSingle = tradeList.filter(i => i.create_user_id !== userId && i.stamps.length === 1)
    state.buyListAlbum = tradeList.filter(i => i.create_user_id !== userId && i.stamps.length > 1)
    let hotStampList = {}
    state.hotStamps = []
    for (let i in state.buyListSingle) {
      if (hotStampList[state.buyListSingle[i].atlas_id] == null) {
        hotStampList[state.buyListSingle[i].atlas_id] = state.buyListSingle[i]
      }
    }
    for (let i in hotStampList) {
      state.hotStamps.push(hotStampList[i])
    }
  }
}

const actions = {
  getTrade ({ commit, rootState }) {
    let userId = rootState.User.user_id
    Vue.http.post(apiHost + '/api/deal/list', {
      user_id: userId
    }).then((response) => {
      if (response.data.success) {
        let tradeList = response.data.result.deals
        if (tradeList) {
          commit('UPDATE_TRADE_LIST', { tradeList, userId })
        }
      }
    })
  },
  createDeal ({ commit, rootState, dispatch }, { stampIds, price }) {
    let data = {
      user_id: rootState.User.user_id,
      stamp_ids: stampIds,
      price: price
    }
    Vue.http.post(apiHost + '/api/deal/create', data).then((response) => {
      if (response.data.success) {
        dispatch('getAlbum')
        dispatch('getTrade')
      }
    })
  },
  cancelTrade ({ commit, rootState, dispatch }, tradeId) {
    let data = {
      user_id: rootState.User.user_id,
      id: tradeId
    }
    Vue.http.post(apiHost + '/api/deal/destroy', data).then((response) => {
      if (response.data.success) {
        dispatch('getAlbum')
        dispatch('getTrade')
      }
    })
  },
  acceptTrade ({ commit, rootState, dispatch }, tradeId) {
    let data = {
      user_id: rootState.User.user_id,
      deal_id: tradeId
    }
    Vue.http.post(apiHost + '/api/deal/accept', data).then((response) => {
      if (response.data.success) {
        dispatch('getAlbum')
        dispatch('getTrade')
      }
    })
  }
}

export default {
  state,
  mutations,
  actions
}
