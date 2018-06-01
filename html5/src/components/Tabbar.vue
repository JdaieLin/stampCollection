<template>
  <div class="tabbar">
    <div class="tab-item" :class="currentTab=='dig'?'active':null" @click="tabClick('dig')">
      <div class="tabicon dig"></div>
      <span>挖一挖</span>
    </div>
    <div class="tab-item" :class="currentTab=='album'?'active':null"  @click="tabClick('album')">
      <div class="tabicon album"></div>
      <span>集邮册</span>
    </div>
    <div class="tab-item" :class="currentTab=='slide'?'active':null"  @click="tabClick('slide')">
      <div class="tabicon slide center"></div>
      <span>划一划</span>
    </div>
    <div class="tab-item" :class="currentTab=='trade'?'active':null"  @click="tabClick('trade')">
      <div class="tabicon trade"></div>
      <span>交易所</span>
    </div>
    <div class="tab-item" :class="currentTab=='user'?'active':null"  @click="tabClick('user')">
      <div class="tabicon user"></div>
      <span>我的号</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tabbar',
  data () {
    return {
      currentTab: 'slide'
    }
  },
  methods: {
    tabClick (name) {
      this.currentTab = name
      let path = name === 'slide' ? '/' : name
      this.$router.push({ path: path })
    }
  },
  mounted () {
    this.currentTab = this.$router.history.current.name
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.tabbar{
  position: fixed;
  width: 100%;
  height: 60px;
  left: 0;
  bottom: 0;
  background-color: #fff;
  box-shadow: 0 0 20px 5px #0000000d;
  display: flex;
  align-items: center;
  justify-content: center;
  .tab-item{
    flex: 1;
    display: block;
    position: relative;
    height: 100%;
    text-align: center;
    .tabicon{
      width: 80%;
      height: 50%;
      margin: 5px auto 8px auto;
      background-size: contain;
      background-repeat: no-repeat;
      background-position: center;
      &.center{
        width: 100%;
        height: 60px;
        margin-top: -20px;
        margin-bottom: 3px;
      }
      &.dig{
        background-image: url(/static/ui/menuDig.disactive.png);
      }
      &.album{
        background-image: url(/static/ui/menuAlbum.disactive.png);
      }
      &.slide{
        background-image: url(/static/ui/menuSlide.disactive.png);
      }
      &.trade{
        background-image: url(/static/ui/menuTrade.disactive.png);
      }
      &.user{
        background-image: url(/static/ui/menuUser.disactive.png);
      }
    }
    &.active{
      .dig{
        background-image: url(/static/ui/menuDig.active.png);
      }
      .album{
        background-image: url(/static/ui/menuAlbum.active.png);
      }
      .slide{
        background-image: url(/static/ui/menuSlide.active.png);
      }
      .trade{
        background-image: url(/static/ui/menuTrade.active.png);
      }
      .user{
        background-image: url(/static/ui/menuUser.active.png);
      }
      span{
        color: #128045;
      }
    }
    span{
      display: block;
      width: 100%;
      line-height: 12px;
      font-size: 12px;
      text-align: center;
      color: #555;
    }
  }
}
</style>
