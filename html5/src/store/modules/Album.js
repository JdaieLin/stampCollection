const state = {
  collectionList: [],
  boughtList: [],
  rest: []
}

// 需要接口获取序列的详情
// http post : /api/stamp/list
// http post : /api/stamp/get   detail

const mutations = {
  DECREMENT_MAIN_COUNTER (state) {
    state.main--
  },
  INCREMENT_MAIN_COUNTER (state, payload) {
    state.main += payload
  }
}

const actions = {
  someAsyncTask ({ commit }, payload) {
    // do something async
    commit('INCREMENT_MAIN_COUNTER', payload)
  }
}

export default {
  state,
  mutations,
  actions
}
