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
    <CoinCounter :number="$store.state.Money.coins" />
    <div class="stack-wrapper"
         @touchmove.stop.prevent="touchmove"
         @touchstart.stop.prevent="touchstart"
         @touchend.stop.prevent="touchend"
         @touchcancel.stop.prevent="touchend">
      <stack ref="stack" :pages="currentLoopList" :stackinit="stackinit" @slideChange="slideChange"></stack>
    </div>
    <div class="controls">
      <button @click="onReverse" :class="$store.state.Slide.canReverseTime>0?'active':null" class="button reverse small"></button>
      <button @click="onPressBad" class="button bad"></button>
      <button @click="onPressGood" class="button good"></button>
      <button @click="onCollect" :class="$store.state.Slide.currentStamp.colleted?'active':null" class="button favor small"></button>
    </div>
  </div>
</template>

<script>
import stack from '../components/Stack'
import CoinCounter from '../components/CoinCounter'

export default {
  name: 'Slide',
  data () {
    return {
      currentLoopList: [],
      stackinit: {
        visible: 3
      },
      canReverseTime: 3,
      isFavor: false,
      currentPage: 0,
      currentPathArray: [],
      pathFade: false,
      drawing: false
    }
  },
  computed: {
    currentPath () {
      if (!this.currentPathArray.length) {
        return ''
      } else {
        return this.currentPathArray.map(i => i.join(' ')).join(',')
      }
    }
  },
  mounted () {
    this.currentLoopList = this.$store.state.Slide.currentLoopList
    let preloadList = this.$store.state.Slide.preloadList
    for (let i in preloadList) {
      let preloadImg = new Image()
      preloadImg.src = '/static/ui/' + preloadList[i]
    }
    this.$store.dispatch('changeLoopIndex', 0)
  },
  components: {
    stack,
    CoinCounter
  },
  methods: {
    onReverse () {
      this.$emit('reverseSlide')
      this.$store.dispatch('reverseSlide')
    },
    onPressBad () {
      this.$refs.stack.$emit('prev')
      this.$store.dispatch('setBad')
    },
    onPressGood () {
      this.$refs.stack.$emit('next')
      this.$store.dispatch('setGood')
    },
    onCollect () {
      this.isFavor = !this.isFavor
      this.$store.dispatch('setCollect')
      // this.$refs.stack.$emit('next')
      this.$emit('favor')
    },
    onBuy () {
      this.$refs.stack.$emit('next')
    },
    slideChange (index) {
      this.currentPage = index
      this.coinNumber += 10
      this.$store.dispatch('changeLoopIndex', index)
      this.$store.dispatch('coinIncreaseSlide')
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
      background-image: url(/static/ui/opReverse.disactive.png);
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
