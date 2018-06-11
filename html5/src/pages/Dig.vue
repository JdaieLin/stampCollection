<template>
  <div class="mid-center">
    <div v-if="shouldShowBox && !openingBox" class="boxes">
      <div class="box-item" @click="openBox"></div>
      <div class="box-item" @click="openBox"></div>
      <div class="box-item" @click="openBox"></div>
    </div>
    <div v-if="!shouldShowBox && !openingBox" class="digging">
      <div class="digging-modal">
        <div class="digging-text">还在挖掘中<br>请耐心等候哦</div>
        <div class="digging-timer">{{timer}}</div>
      </div>
    </div>
    <div v-if="openingBox" class="opening">
      <div class="opening-modal">
        <div class="opening-text">恭喜<br>获得{{currentBox.ingot}}个元宝</div>
        <div class="opening-count"><span class="ingot-icon"></span> x {{currentBox.ingot}}</div>
        <div class="reward-btn" @click="getReward">领取奖励</div>
      </div>
    </div>
  </div>
</template>

<script>
const waitDuration = 3600
export default {
  name: 'Dig',
  data () {
    return {
      lastOpenTime: 0,
      lastOpenIndex: 0,
      currentTime: 0,
      openingBox: false,
      currentBox: null
    }
  },
  computed: {
    shouldShowBox () {
      return (this.currentTime - this.lastOpenTime > waitDuration * 1000)
    },
    timer () {
      if (!this.lastOpenTime) {
        return `00:00`
      }
      let duration = waitDuration - Math.floor((this.currentTime - this.lastOpenTime) / 1000)
      let min = Math.floor(duration / 60)
      let sec = duration % 60
      if (min > 60) {
        min = 0
      }
      if (min < 10) {
        min = '0' + min
      }
      if (sec < 10) {
        sec = '0' + sec
      }
      return `${min}:${sec}`
    }
  },
  mounted () {
    let lastOpenTime = localStorage.getItem('lastOpenTime')
    if (lastOpenTime) {
      this.lastOpenTime = parseInt(lastOpenTime)
    }
    if (!this.lastOpenTime) {
      this.lastOpenTime = new Date().getTime()
      localStorage.setItem('lastOpenTime', this.lastOpenTime)
    }
    let that = this
    setInterval(function () {
      that.currentTime = new Date().getTime()
    }, 500)
  },
  methods: {
    openBox () {
      let hour = new Date().getHours()
      let chests = this.$store.state.Money.chests
      this.currentBox = chests[hour]
      this.openingBox = true
    },
    getReward () {
      this.lastOpenTime = new Date().getTime()
      localStorage.setItem('lastOpenTime', this.lastOpenTime)
      this.$store.dispatch('syncChest', this.currentBox)
      this.$store.dispatch('ingotIncrease', this.currentBox.ingot)
      this.openingBox = false
    }
  }
}
</script>

<style scoped lang="less">
  @import "../less/common.less";
  @keyframes show {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
  .mid-center{
    width: 100%;
    height: 100%;
    background-image: url(/static/ui/Dig_Main.png);
    background-size: cover;
    .boxes{
      width: 100%;
      height: 100%;
      display: flex;
      padding: 60px 40px;
      flex-direction: column;
      align-items: center;
      .box-item{
        /*position: absolute;*/
        width: 178px;
        height: 152px;
        background-image: url(/static/ui/box1.close.png);
        background-size: contain;
        background-repeat: no-repeat;
        margin-top: 6px;
        &:nth-child(1){
          position: relative;
          left: -60px;
          &:active{
            transform: scale(0.95);
          }
        }
        &:nth-child(2){
          position: relative;
          left: 60px;
          transform: rotate(40deg);
          &:active{
            transform: rotate(40deg)scale(0.95);
          }
        }
        &:nth-child(3){
          position: relative;
          left: -60px;
          transform: rotate(5deg);
          &:active{
            transform: rotate(5deg)scale(0.95);
          }
        }
      }
    }
    .digging{
      position: fixed;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 100%;
      height: 100%;
      background-color: #000000aa;
      animation: show 0.3s ease-in-out;
      .digging-modal{
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        width: 220px;
        height: 180px;
        background-color: #fff;
        margin-top: -80px;
        border-radius: 8px;
        box-shadow: 0 0 10px 5px #0000005e;
        &:before{
          content: '';
          display: block;
          position: absolute;
          top: -60px;
          width: 60px;
          height: 120px;
          background-image: url(/static/ui/shove.png);
          background-repeat: no-repeat;
          background-size: contain;
          background-position: center;
        }
        .digging-text{
          font-size: 14px;
          margin-top: 40px;
        }
        .digging-timer{
          margin-top: 20px;
          font-weight: bold;
          font-size: 32px;
          color: @themeColor;
        }
      }
    }
    .opening{
      position: fixed;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 100%;
      height: 100%;
      background-color: #000000aa;
      .opening-modal{
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        width: 220px;
        height: 180px;
        background-color: #fff;
        margin-top: -80px;
        border-radius: 8px;
        box-shadow: 0 0 10px 5px #0000005e;
        &:before{
          content: '';
          display: block;
          position: absolute;
          top: -70px;
          width: 150px;
          height: 120px;
          background-image: url(/static/ui/Box.open.png);
          background-repeat: no-repeat;
          background-size: contain;
          background-position: center;
        }
        .opening-text{
          font-size: 14px;
        }
        .opening-count{
          margin-top: 6px;
          .ingot-icon{
            display: inline-block;
            position: relative;
            top: 3px;
            width: 21px;
            height: 15px;
            background-image: url(/static/ui/ignotMini.png);
            background-repeat: no-repeat;
            background-size: contain;
          }
        }
        .reward-btn{
          position: absolute;
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: 14px;
          width: 75%;
          height: 28px;
          bottom: 16px;
          background-color: @themeColor;
          color: #fff;
          border-radius: 14px;
          &:active{
            background-color: darken(@themeColor, 6%);
          }
        }
      }
    }
  }
</style>
