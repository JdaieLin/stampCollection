const state = {
  main: 0
}

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
