import Vue from 'vue'
import apiHost from '../host'

const state = {
  collectList: [],
  collectBoughtList: [],
  boughtList: [],
  tradingList: [],
  finalBoughtList: [],
  collectBook: [],
  boughtBook: [],
  yearBook: [],
  allStamp: []
}

// 需要接口获取序列的详情
// http post : /api/stamp/list  100.未收藏。101.已收藏。102.收藏认购。103.交易认购。104.交易中。
// http post : /api/stamp/get   detail

const mutations = {
  UPDATE_YEAR_BOOK (state, newList) {
    let year = ''
    state.allStamp = []
    state.yearBook = []
    for (let i in newList) {
      newList[i].set_name = newList[i].name
      newList[i].set = newList[i].Atlas[0].set
      newList[i].image = newList[i].Atlas[0].image
      newList[i].year = newList[i].Atlas[0].year
      if (newList[i].year !== year) {
        year = newList[i].year
        newList[i].yearStart = newList[i].year
      }
      state.yearBook.push(newList[i])
      state.yearBook[newList[i].set] = newList[i]
      for (let j in newList[i].Atlas) {
        state.allStamp.push(newList[i].Atlas[j])
      }
    }
  },
  CHANGE_LOCK_YEARBOOK (state, newList) {
    let unlockIds = newList.map(i => i.serial_id)
    for (let i = 0; i < state.yearBook.length; i++) {
      if (unlockIds.indexOf(state.yearBook[i].id) < 0) {
        state.yearBook[i].lock = true
      }
    }
  },
  UPDATE_BOOKS (state) {
    // 更新收藏邮册
    state.collectBook = []
    let collectBook = state.collectList.filter((i, k, a) => a.map(j => j.set).indexOf(i.set) === k).sort((a, b) => b.year - a.year)
    for (let i in collectBook) {
      collectBook[i].lock = false
      state.collectBook.push(collectBook[i])
    }
    // 更新购买邮票总列表
    state.boughtList = state.boughtList.map(i => {
      i.boughtType = 'trade'
      return i
    })
    state.collectBoughtList = state.collectBoughtList.map(i => {
      i.boughtType = 'collect'
      return i
    })
    state.tradingList = state.tradingList.map(i => {
      i.trading = true
      return i
    })
    let finalBoughtList = [...state.boughtList, ...state.collectBoughtList, ...state.tradingList].sort((a, b) => b.year - a.year)
    state.finalBoughtList = finalBoughtList
    // 更新购买邮册
    state.boughtBook = []
    let boughtBook = []
    for (let i in finalBoughtList) {
      let set = finalBoughtList[i].set
      finalBoughtList[i].lock = false
      if (boughtBook[set] == null) {
        boughtBook.push(finalBoughtList[i])
        boughtBook[set] = finalBoughtList[i]
      }
    }
    state.boughtBook = boughtBook
    // 集合多品相单张 collectList
    let newCollectList = {}
    for (let i in state.collectList) {
      let atlas = state.collectList[i].atlas_id
      if (newCollectList[atlas] == null) {
        newCollectList[atlas] = [state.collectList[i]]
      } else {
        newCollectList[atlas].push(state.collectList[i])
      }
    }
    state.collectList = []
    for (let i in newCollectList) {
      if (newCollectList[i].length === 1) {
        state.collectList.push(newCollectList[i][0])
      } else {
        state.collectList.push({
          multiple: true,
          set: newCollectList[i][0].set,
          serial_id: newCollectList[i][0].serial_id,
          list: newCollectList[i]
        })
      }
    }
    // 集合多品相单张 finalBoughtList
    let newBuyList = {}
    for (let i in state.finalBoughtList) {
      let atlas = state.finalBoughtList[i].atlas_id
      if (newBuyList[atlas] == null) {
        newBuyList[atlas] = [state.finalBoughtList[i]]
      } else {
        newBuyList[atlas].push(state.finalBoughtList[i])
      }
    }
    finalBoughtList = []
    for (let i in newBuyList) {
      if (newBuyList[i].length === 1) {
        finalBoughtList.push(newBuyList[i][0])
      } else {
        finalBoughtList.push({
          multiple: true,
          set: newBuyList[i][0].set,
          serial_id: newBuyList[i][0].serial_id,
          list: newBuyList[i]
        })
      }
    }
    state.finalBoughtList = finalBoughtList
  },
  UPDATE_COLLECT_LIST (state, newList) {
    if (!newList) return
    state.collectList.length = 0
    for (let i in newList) {
      state.collectList.push(newList[i])
    }
  },
  UPDATE_BOUGHT_LIST (state, newList) {
    if (!newList) return
    state.boughtList.length = 0
    for (let i in newList) {
      state.boughtList.push(newList[i])
    }
  },
  UPDATE_COLLECT_BOUGHT_LIST (state, newList) {
    if (!newList) return
    state.collectBoughtList.length = 0
    for (let i in newList) {
      state.collectBoughtList.push(newList[i])
    }
  },
  UPDATE_TRADING_LIST (state, newList) {
    if (!newList) return
    state.tradingList.length = 0
    for (let i in newList) {
      state.tradingList.push(newList[i])
    }
  },
  UNLOCK_SINGLE_BOOK (state, id) {
    for (let i in state.yearBook) {
      if (state.yearBook[i].id === id) {
        state.yearBook[i].lock = false
      }
    }
  }
}

const actions = {
  getAlbum ({ commit, rootState }) {
    let fin = 4
    Vue.http.post(apiHost + '/api/stamp/list', {
      user_id: rootState.User.user_id,
      status: [101] // 已收藏
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_COLLECT_LIST', response.data.result.Stamps)
        if (!--fin) commit('UPDATE_BOOKS')
      }
    })
    Vue.http.post(apiHost + '/api/stamp/list', {
      user_id: rootState.User.user_id,
      status: [102] // 收藏认购
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_COLLECT_BOUGHT_LIST', response.data.result.Stamps)
        if (!--fin) commit('UPDATE_BOOKS')
      }
    })
    Vue.http.post(apiHost + '/api/stamp/list', {
      user_id: rootState.User.user_id,
      status: [103] // 交易认购
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_BOUGHT_LIST', response.data.result.Stamps)
        if (!--fin) commit('UPDATE_BOOKS')
      }
    })
    Vue.http.post(apiHost + '/api/stamp/list', {
      user_id: rootState.User.user_id,
      status: [104] // 交易中
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_TRADING_LIST', response.data.result.Stamps)
        if (!--fin) commit('UPDATE_BOOKS')
      }
    })
  },
  getYearBook ({ commit, dispatch, rootState }) {
    Vue.http.post(apiHost + '/api/base/all', {
      user_id: rootState.User.user_id
    }).then((response) => {
      if (response.data.success) {
        commit('UPDATE_YEAR_BOOK', response.data.result.All)
        dispatch('getUnlockBook')
      }
    })
  },
  getUnlockBook ({ commit, rootState }) {
    Vue.http.post(apiHost + '/api/slot/list', {
      user_id: rootState.User.user_id
    }).then((response) => {
      if (response.data.success) {
        commit('CHANGE_LOCK_YEARBOOK', response.data.result.slots)
      }
    })
  },
  unlockBook ({ commit, rootState, dispatch }, serialId) {
    Vue.http.post(apiHost + '/api/slot/unlock', {
      user_id: rootState.User.user_id,
      serial_id: serialId
    }).then((response) => {
      if (response.data.success) {
        commit('UNLOCK_SINGLE_BOOK', serialId)
      }
    })
  },
  albumBuyStamp ({ commit, rootState, dispatch }, stamp) {
    let data = {
      user_id: rootState.User.user_id,
      stamp_id: stamp.id
    }
    Vue.http.post(apiHost + '/api/stamp/buy', data).then((response) => {
      if (response.data.success) {
        dispatch('getAlbum')
        dispatch('getTrade')
      }
    })
  },
  albumRecycleStamp ({ commit, rootState, dispatch }, stamp) {
    let data = {
      user_id: rootState.User.user_id,
      stamp_id: stamp.id
    }
    Vue.http.post(apiHost + '/api/stamp/recycle', data).then((response) => {
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
