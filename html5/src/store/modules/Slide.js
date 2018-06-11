import apiHost from '../host'
import Vue from 'vue'

const slideLength = 15

const state = {
  stamps: [],
  currentStamp: null,
  currentLoopList: [],
  currentLoopIndex: 0,
  canReverseTime: localStorage.getItem('reverseCountDate_' + new Date().toLocaleDateString()) || 3,
  start: false,
  slideResultList: [],
  preloadList: [
    'coin.png',
    'coin.w-shadow.png',
    'coinCount.back.png',
    'coinCount.shadow.png',
    'counter.number.jpg',
    'material.gold.jpg',
    'menuAlbum.active.png',
    'menuAlbum.disactive.png',
    'menuDig.active.png',
    'menuDig.disactive.png',
    'menuSlide.active.png',
    'menuSlide.disactive.png',
    'menuTrade.active.png',
    'menuTrade.disactive.png',
    'menuUser.active.png',
    'menuUser.disactive.png',
    'opBad.active.png',
    'opBad.active.pressed.png',
    'opBad.disactive.png',
    'opBad.disactive.pressed.png',
    'opFavor.active.png',
    'opFavor.active.pressed.png',
    'opFavor.disactive.png',
    'opFavor.disactive.pressed.png',
    'opGood.active.png',
    'opGood.active.pressed.png',
    'opGood.disactive.png',
    'opGood.disactive.pressed.png',
    'opReverse.active.png',
    'opReverse.active.pressed.png',
    'opReverse.disactive.png',
    'opReverse.disactive.pressed.png'
  ]
}

// http post : /api/sweep/list
// http post : /api/sweep/sync

// http post : /api/stamp/buy
// http post : /api/stamp/collect
// http post : /api/stamp/recycle

const mutations = {
  START (state) {
    state.start = true
  },
  UPDATE_LOOP_LIST (state, newList) {
    if (state.currentLoopList.length === 0) {
      for (let i in newList) {
        state.currentLoopList.push(newList[i])
      }
      state.currentLoopIndex = 0
    } else {
      let updateIndex
      if (state.currentLoopIndex > 9) {
        updateIndex = 0
      } else if (state.currentLoopIndex > 4) {
        updateIndex = 10
      } else {
        updateIndex = 5
      }
      for (let i in newList) {
        state.currentLoopList[(updateIndex + i) % slideLength] = newList[i]
      }
    }
    state.currentStamp = state.currentLoopList[state.currentLoopIndex]
  },
  CHANGE_LOOP_INDEX (state, index) {
    if (state.currentStamp) {
      // save slide result
      state.slideResultList.push({
        stamp_id: state.currentStamp.id,
        action: state.currentStamp.action
      })
    }
    state.currentLoopIndex = index
    state.currentStamp = state.currentLoopList[state.currentLoopIndex]
  },
  SET_BAD (state) {
    state.currentStamp.action = 0
  },
  SET_GOOD (state) {
    state.currentStamp.action = 1
  },
  SET_FAVOR (state) {
    state.currentStamp.action = 2
    if (state.currentStamp.favor != null) {
      state.currentStamp.favor = !state.currentStamp.favor
    } else {
      state.currentStamp.favor = true
    }
  },
  SET_BUY (state) {
    state.currentStamp.action = 3
  },
  REVERSE_SLIDE (state) {
    if (parseInt(state.canReverseTime) <= 0) return
    let stamp = state.currentLoopList.pop()
    state.currentLoopList.unshift(stamp)
    state.currentStamp = state.currentLoopList[state.currentLoopIndex]
    if (state.canReverseTime) {
      state.canReverseTime--
      localStorage.setItem('reverseCountDate_' + new Date().toLocaleDateString(), state.canReverseTime)
    }
  },
  CLEAR_SLIDE_RESULT (state) {
    state.slideResultList = []
  }
}

const actions = {
  getLoopSlideInit ({commit, rootState}) {
    let data = {
      user_id: rootState.User.user_id,
      max: slideLength
    }
    Vue.http.post(apiHost + '/api/sweep/list', data).then((response) => {
      if (response.data.success) {
        commit('UPDATE_LOOP_LIST', response.data.result.stamps)
      }
    })
  },
  changeLoopIndex ({commit, rootState}, index) {
    commit('CHANGE_LOOP_INDEX', index)
    if (!state.start) {
      commit('START')
    } else {
      if (index === 0 || index === 5 || index === 10) {
        let data = {
          user_id: rootState.User.user_id,
          max: 5
        }
        Vue.http.post(apiHost + '/api/sweep/list', data).then((response) => {
          if (response.data.success) {
            console.log('get 5 more slides')
            commit('UPDATE_LOOP_LIST', response.data.result.stamps)
          }
        })
      }
    }
  },
  setGood ({commit}) {
    commit('SET_GOOD')
  },
  setBad ({commit}) {
    commit('SET_BAD')
  },
  setBuy ({commit}) {
    commit('SET_BUY')
  },
  setFavor ({commit}) {
    commit('SET_FAVOR')
  },
  reverseSlide ({commit}) {
    commit('REVERSE_SLIDE')
  },
  syncSlideReult ({commit, rootState}) {
    let data = {
      user_id: rootState.User.user_id,
      stamps: state.slideResultList
    }
    commit('CLEAR_SLIDE_RESULT')
    Vue.http.post(apiHost + '/api/sweep/sync', data).then((response) => {
      if (response.data.success) {
        console.log('slide sync success')
      }
    })
  }
}

export default {
  state,
  mutations,
  actions
}
