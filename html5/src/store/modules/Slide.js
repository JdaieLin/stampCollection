const state = {
  stamps: [
    {
      stamp_id: 1,
      brightness: 1.1,
      crack: 1.1,
      stain: 1.1,
      postmark: 2,
      image_small: '',
      heat: 1.1// 此处需要发行量数据
    }
  ],
  currentStamp: {},
  currentLoopList: [
    {
      title: '价值连城的邮票',
      age: '2018',
      level: 98,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/1.png'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 72,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo2.jpg'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 68,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo3.jpg'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 89,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo4.jpg'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 76,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo5.jpg'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 40,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo6.jpg'
    },
    {
      title: 'stamp name',
      age: '2018',
      level: 96,
      total: 300,
      rest: 23,
      imgSrc: '../../static/img/demo7.jpg'
    }
  ],
  currentLoopIndex: 0,
  canReverseTime: 3,
  slideResultList: [
    {
      stamp_id: 0,
      action: 0
    }
  ],
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
  CHANGE_LOOP_INDEX (state, index) {
    state.currentLoopIndex = index
    state.currentStamp = state.currentLoopList[state.currentLoopIndex]
  },
  SET_BAD (state) {
    state.currentStamp.action = 0
  },
  SET_GOOD (state) {
    state.currentStamp.action = 1
  },
  SET_COLLECT (state) {
    state.currentStamp.action = 2
    state.currentStamp.collected = true
    console.log(state.currentStamp)
  },
  SET_BUY (state) {
    state.currentStamp.action = 3
  },
  REVERSE_SLIDE (state) {
    if (!state.canReverseTime) return
    let stamp = state.currentLoopList.pop()
    state.currentLoopList.unshift(stamp)
    state.currentStamp = state.currentLoopList[state.currentLoopIndex]
    if (state.canReverseTime) {
      state.canReverseTime--
    }
  }
}

const actions = {
  changeLoopIndex ({commit}, index) {
    commit('CHANGE_LOOP_INDEX', index)
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
  setCollect ({commit}) {
    commit('SET_COLLECT')
  },
  reverseSlide ({commit}) {
    commit('REVERSE_SLIDE')
  }
}

export default {
  state,
  mutations,
  actions
}
