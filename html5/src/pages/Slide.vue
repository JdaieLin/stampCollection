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
      <button @click="onPressBad" class="button bad small"></button>
      <button @click="onReverse" :class="$store.state.Slide.canReverseTime>0?'active':null" class="button reverse"></button>
      <button @click="onFavor" :class="isFavor?'active':null" class="button favor"></button>
      <button @click="onPressGood" class="button good small"></button>
    </div>
    <div v-if="askModal" class="ask-modal" @click="closeAskModal">
      <div class="modal-center" @click="stopPropagation">
        <div class="favor-icon"></div>
        <div class="col">
          <div class="text">花费元宝</div>
          <div class="text"><span class="icon ingot-icon"></span> x 20</div>
          <div class="action-btn theme-btn white" @click="actionCollect">收藏五天</div>
        </div>
        <div class="col">
          <div class="text">花费以太币</div>
          <div class="text"><span class="icon eth-icon"></span> x 10</div>
          <div class="action-btn theme-btn green" @click="actionBuy">直接购买</div>
        </div>
        <div class="close-btn" @click="closeAskModal"></div>
      </div>
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
        visible: 3,
        currentPage: this.$store.state.Slide.currentLoopIndex || 0
      },
      canReverseTime: 3,
      isFavor: false,
      currentPage: 0,
      currentPathArray: [],
      pathFade: false,
      drawing: false,
      syncCount: 0,
      askModal: false
    }
  },
  computed: {
    currentPath () {
      if (!this.currentPathArray.length) {
        return ''
      } else {
        return this.currentPathArray.map(i => i.join(' ')).join(',')
      }
    },
    isStoreFavor () {
      return this.$store.state.Slide.currentStamp.favor
    }
  },
  mounted () {
    this.currentLoopList = this.$store.state.Slide.currentLoopList
    let preloadList = this.$store.state.Slide.preloadList
    for (let i in preloadList) {
      let preloadImg = new Image()
      preloadImg.src = '/static/ui/' + preloadList[i]
    }
    this.$store.dispatch('refreshCoin')
    this.$store.dispatch('changeLoopIndex', 0)
    let that = this
    setInterval(function () {
      if (that.syncCount) {
        if (that.syncCount === 1) {
          that.$store.dispatch('syncSlideReult')
        }
        that.syncCount--
      }
    }, 1000)
  },
  components: {
    stack,
    CoinCounter
  },
  methods: {
    onReverse () {
      this.$store.dispatch('reverseSlide')
      this.$emit('reverseSlide')
    },
    onPressBad () {
      this.$store.dispatch('setBad')
      this.$refs.stack.$emit('prev')
    },
    onPressGood () {
      this.$store.dispatch('setGood')
      this.$refs.stack.$emit('next')
    },
    onFavor () {
      this.isFavor = !this.isFavor
      if (this.isFavor) {
        this.askModal = true
      }
      this.$store.dispatch('setFavor', this.isFavor)
      this.$emit('favor')
    },
    stopPropagation (e) {
      e.stopPropagation()
    },
    actionCollect () {
      console.log('collect', this.$store.state.Slide.currentStamp)
    },
    actionBuy () {
      console.log('buy', this.$store.state.Slide.currentStamp)
    },
    closeAskModal () {
      this.askModal = false
    },
    slideChange (index) {
      this.currentPage = index
      this.$store.dispatch('changeLoopIndex', index)
      this.$store.dispatch('coinIncrease', 10)
      this.syncCount = 4
      this.isFavor = false
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
  @import "../less/common.less";
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
    height: 100%;
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
    margin: 40px auto 60px;
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
  .ask-modal{
    position: fixed;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
    background-color: #000000aa;
    z-index: 1000;
    .modal-center{
      position: relative;
      width: 260px;
      height: 200px;
      background-color: #fff;
      border-radius: 8px;
      box-shadow: 0 0 10px 5px #0000005e;
      display: flex;
      flex-wrap: wrap;
      padding: 20px;
      .favor-icon{
        width: 100%;
        height: 36px;
        background-image: url(/static/ui/opFavor.small.png);
        background-size: contain;
        background-repeat: no-repeat;
        background-position: center;
      }
      .col{
        position: relative;
        flex: 1;
        font-size: 13px;
        margin-top: -40px;
        .text{
          .icon{
            display: inline-block;
            position: relative;
            top: 7px;
            width: 22px;
            height: 22px;
            background-size: contain;
            background-repeat: no-repeat;
            background-position: center;
            &.ingot-icon{
              background-image: url(/static/ui/ignotMini.png);
            }
            &.eth-icon{
              background-image: url(/static/ui/ether.png);
            }
          }
        }
        .action-btn{
          position: absolute;
          width: 80%;
          height: 30px;
          left: 10%;
          bottom: 10px;
          border-radius: 15px;
        }
      }
      .close-btn{
        position: absolute;
        width: 58px;
        height: 58px;
        bottom: -30px;
        left: 50%;
        margin-top: 40px;
        margin-left: -29px;
        border-radius: 50%;
        background-image: url(/static/ui/single.close.png);
        background-size: cover;
        &:active{
          transform: scale(0.9);
        }
      }
    }
  }
</style>
