<template>
  <div class="mid-center">
    <div class="pathSpace" :class="pathFade?'fade':null">
      <svg width="100%" height="100%" version="1.1"
           xmlns="http://www.w3.org/2000/svg">
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#ffffff33;stroke-width:12px"/>
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#ffffffaa;stroke-width:10px"/>
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#ffffffcc;stroke-width:8px"/>
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#ffffeccc;stroke-width:6px"/>
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#ffffffff;stroke-width:4px"/>
        <polyline :points="currentPath"
                  style="fill:transparent;stroke:#23ce7b;stroke-width:2px"/>
      </svg>
    </div>
    <CoinCounter :number="coinNumber" />
    <div class="stack-wrapper"
         @touchmove.stop.prevent="touchmove"
         @touchstart.stop.prevent="touchstart"
         @touchend.stop.prevent="touchend"
         @touchcancel.stop.prevent="touchend">
      <stack ref="stack" :pages="someList" :stackinit="stackinit" :updateTrigger="updateTrigger" @slideChange="slideChange"></stack>
    </div>
    <div class="controls">
      <button @click="onReverse" :class="canReverseTime>3?'active':null" class="button reverse small"></button>
      <button @click="onBad" class="button bad"></button>
      <button @click="onGood" class="button good"></button>
      <button @click="onFavor" :class="isFavor?'active':null" class="button favor small"></button>
    </div>
  </div>
</template>

<script>
import stack from '../components/Stack'
import CoinCounter from '../components/CoinCounter'

const preloadList = [
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

export default {
  name: 'Slide',
  data () {
    return {
      someList: [],
      stackinit: {
        visible: 3
      },
      coinNumber: 0,
      updateTrigger: false,
      canReverseTime: 3,
      isFavor: false,
      currentPage: 0,
      currentPathArray: [],
      pathFade: false,
      drawing: false
    }
  },
  computed: {
    count () {
      return this.$store.state.Counter.main
    },
    currentPath () {
      if (!this.currentPathArray.length) {
        return ''
      } else {
        return this.currentPathArray.map(i => i.join(' ')).join(',')
      }
    }
  },
  mounted () {
    let that = this
    that.someList = [
      {
        title: '价值连城的邮票',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/1.png'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo2.jpg'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo3.jpg'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo4.jpg'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo5.jpg'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo6.jpg'
      },
      {
        title: 'stamp name',
        age: '2018',
        grade: 98,
        total: 300,
        rest: 23,
        imgSrc: '../../static/img/demo7.jpg'
      }
    ]
    for (let i in preloadList) {
      let preloadImg = new Image()
      preloadImg.src = '/static/ui/' + preloadList[i]
    }
  },
  components: {
    stack,
    CoinCounter
  },
  methods: {
    onReverse () {
      this.$emit('reverse')
      console.log('insert page before ', this.currentPage)
    },
    onBad () {
      this.$refs.stack.$emit('prev')
    },
    onGood () {
      this.$refs.stack.$emit('next')
    },
    onFavor () {
      this.isFavor = !this.isFavor
      this.$emit('favor')
    },
    slideChange (index) {
      console.log('pageChange to', index)
      this.currentPage = index
      this.coinNumber += 10
      // if (index !== 0) {
      //   this.someList[index - 1].read = true
      // } else {
      //   this.someList[this.someList.length - 1].read = true
      // }
    },
    touchstart (e) {
      this.drawing = true
      this.pathFade = false
      this.currentPathArray = []
      this.currentPathArray.push([e.changedTouches[0].pageX, e.changedTouches[0].pageY])
    },
    touchmove (e) {
      this.drawing = true
      this.pathFade = false
      this.currentPathArray.push([e.changedTouches[0].pageX, e.changedTouches[0].pageY])
    },
    touchend (e) {
      this.drawing = false
      this.pathFade = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
  h1, h2 {
    font-weight: normal;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }
  .mid-center{
    display: flex;
    position: relative;
    width: 100%;
    height: calc(100% - 70px);
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .stack-wrapper{
    margin: 0 auto;
    position: relative;
    z-index: 1000;
    width: 320px;
    height: 320px;
    padding: 0;
    list-style: none;
    pointer-events: none;
  }
  .controls{
    position: relative;
    width: 100%;
    text-align: center;
    display:flex;
    justify-content:center;
    margin: 0 auto;
    margin-top: 40px
  }
  .controls .button {
    border: none;
    background: none;
    position: relative;
    display: inline-block;
    cursor: pointer;
    font-size: 16px;
    width: 75px;
    height: 75px;
    z-index: 100;
    -webkit-tap-highlight-color: rgba(0,0,0,0);
    border-radius: 50%;
    background: #fff;
    margin: 0 8px;
    background-size: contain;
    outline: none;
    &.small{
      width: 48px;
      height: 48px;
      margin-top: 14px;
    }
    &.reverse{
      background-image: url(/static/ui/opReverse.active.png);
      &:active{
        background-image: url(/static/ui/opReverse.active.pressed.png);
      }
      &.active{
        background-image: url(/static/ui/opReverse.active.png);
        &:active{
          background-image: url(/static/ui/opReverse.disactive.pressed.png);
        }
      }
    }
    &.bad{
      background-image: url(/static/ui/opBad.disactive.png);
      &:active{
        background-image: url(/static/ui/opBad.active.pressed.png);
      }
    }
    &.good{
      background-image: url(/static/ui/opGood.disactive.png);
      &:active{
        background-image: url(/static/ui/opGood.active.pressed.png);
      }
    }
    &.favor{
      background-image: url(/static/ui/opFavor.disactive.png);
      &:active{
        background-image: url(/static/ui/opFavor.disactive.pressed.png);
      }
      &.active{
        background-image: url(/static/ui/opFavor.active.png);
        &:active{
          background-image: url(/static/ui/opFavor.disactive.pressed.png);
        }
      }
    }
  }
  .controls .text-hidden {
    position: absolute;
    overflow: hidden;
    width: 0;
    height: 0;
    color: transparent;
    display: block;
  }
  .pathSpace{
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    pointer-events: none;
    z-index: 10000;
    opacity: 1;
    &.fade{
      transition: all 1s ease-out;
      opacity: 0;
    }
  }
</style>
