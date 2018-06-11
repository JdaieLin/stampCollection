import Vue from 'vue'
import apiHost from '../host'

const state = {
  collectList: [],
  collectOwnList: [],
  boughtList: [],
  rest: []
}

// 需要接口获取序列的详情
// http post : /api/stamp/list  100.未收藏。101.已收藏。102.收藏认购。103.交易认购。104.交易中。
// http post : /api/stamp/get   detail

const mutations = {
  UPDATE_COLLECT (state, newList) {
    state.collectList.length = 0
    for (let i in newList) {
      state.collectList.push(newList[i])
    }
  }
}

const actions = {
  getAlbum ({ commit, rootState }) {
    Vue.http.post(apiHost + '/api/stamp/list', {
      user_id: rootState.User.user_id,
      status: [101]
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_COLLECT', response.data.result.stamps)
      }
    })
  }
}

export default {
  state,
  mutations,
  actions
}
