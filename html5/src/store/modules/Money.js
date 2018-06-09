const state = {
  coins: 0,
  ingot: 0,
  chests: [
    {
      hour: 0,
      ingot: 0,
      opened: true
    }
  ]
}

// http post : /api/chest/list
// http post : /api/chest/sync

const mutations = {
  INCREMENT_COIN (state) {
    state.coins += 10
  }
}

const actions = {
  coinIncreaseSlide ({ commit }) {
    commit('INCREMENT_COIN', 10)
  }
}

export default {
  state,
  mutations,
  actions
}
