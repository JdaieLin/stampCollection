const state = {
  main: 0,
  name: '傅红雪',
  phone: '13500000000',
  mail: 'user@qq.com',
  login_id: 'abc',
  login_password: '123',
  network: 101
}

// http post : /api/user/[register/update]

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
