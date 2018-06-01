<template>
  <div class="stamp-sizer">
    <figure class="stamp-wrap" ref="figure" :class="[frame?null:'no-frame', type, lock?'lock': null]">
      <img class="stamp-img" :src="imgSrc" alt="">
      <div class="stamp-frame top"></div>
      <div class="stamp-frame bottom"></div>
      <div class="stamp-frame right"></div>
      <div class="stamp-frame left"></div>
      <div v-if="frame" class="material-frame gold">
        <div class="top"></div>
        <div class="bottom"></div>
        <div class="left"></div>
        <div class="right"></div>
      </div>
    </figure>
    <div class="sizer" ref="sizer"></div>
  </div>
</template>
<script>
/**
 * type
 * String: large/thumb/list
 */
export default {
  name: 'StampWrap',
  props: {
    imgSrc: String,
    level: Number,
    frame: Boolean,
    type: String,
    padding: Number,
    vPadding: Number,
    lock: {
      type: Boolean,
      value: false
    }
  },
  data () {
    return {
      ratio: 1,
      width: 0,
      height: 0
    }
  },
  mounted () {
    let that = this
    let img = new Image()
    img.onload = function () {
      that.ratio = this.width / this.height
      that.setPosition()
    }
    img.src = this.imgSrc
    if (img.width && img.height) {
      that.ratio = this.width / this.height
      that.setPosition()
    }
  },
  methods: {
    setPosition () {
      let padding = this.padding ? this.padding : 16
      let vPadding = this.vPadding ? this.vPadding : padding
      let cw = this.$refs.sizer.clientWidth - padding * 2
      let ch = this.$refs.sizer.clientHeight - vPadding * 2
      let figure = this.$refs.figure
      figure.style.margin = 'auto'
      figure.style.position = 'relative'
      if (this.ratio > cw / ch) {
        figure.style.width = cw + 'px'
      } else {
        figure.style.width = ch * this.ratio + 'px'
      }
      setTimeout(function () {
        figure.style.transition = 'all 0.1s ease-in-out'
        figure.style.opacity = '1'
      }, 10)
    }
  }
}
</script>
<style lang="less">
  @material-weight: 6px;
  @material-weight-list: 3px;
  .sizer{
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
  .stamp-wrap{
    position: absolute;
    top: 0;
    left: 0;
    -webkit-filter: drop-shadow(0px 0px 10px rgba(0,0,0,0.06));
    background-color: #fff;
    opacity: 0;
    transform: scale3d(1,1,1);
    &.lock{
      filter: grayscale(100%);
    }
    &.large{
      padding: @material-weight;
      box-shadow: 0 0 10px 8px #00000045;
    }
    &.list{
      padding: @material-weight-list;
      box-shadow: 0 0 2px 5px #00000015;
    }
    &.thumb{
      /*box-shadow: 0 0 2px 2px #00000055;*/
    }
    &.no-frame{
      padding: 0;
    }
    .material-frame{
      position: absolute;
      width: 100%;
      height: 100%;
      top: 0;
      left: 0;
      z-index: 20;
      &.gold div{
        background-image: url("/static/ui/material.gold.jpg");
      }
      .top{
        position: absolute;
        width: 100%;
        height: @material-weight;
        top: 0;
        left: 0;
        background-position: top left;
        background-size: cover;
      }
      .bottom{
        position: absolute;
        width: 100%;
        height: @material-weight;
        bottom: 0;
        left: 0;
        background-position: bottom left;
        background-size: cover;
      }
      .left{
        position: absolute;
        height: 100%;
        width: @material-weight;
        top: 0;
        left: 0;
        background-position: top left;
        background-size: cover;
      }
      .right{
        position: absolute;
        height: 100%;
        width: @material-weight;
        top: 0;
        right: 0;
        background-position: top right;
        background-size: cover;
      }
    }
    &.list:before{
      content: '';
      box-sizing: border-box;
      position: absolute;
      width: calc(100% + 2px);
      height: calc(100% + 2px);
      top: -1px;
      left: -1px;
      z-index: 10;
      border: 4px solid white;
    }
    &.list .material-frame{
      .top, .bottom{
        height: @material-weight-list;
      }
      .left, .right{
        width: @material-weight-list;
      }
    }
    &.large .stamp-frame{
      position: absolute;
      background-color: #fff;
      background: radial-gradient(
        transparent 0px,
        transparent 2px,
        white 2px,
        white
      );
      background-size: 8px 8px;
      &.top{
        width: calc(100% + 10px);
        height: 5px;
        bottom: 100%;
        left: -5px;
        background-position: 2px -3px;
      }
      &.bottom{
        width: calc(100% + 10px);
        height: 5px;
        top: 100%;
        left: -5px;
        background-position: 2px 0px;
      }
      &.right{
        width: 5px;
        height: 100%;
        top: 0;
        left: 100%;
        background-position: 0px 0px;
      }
      &.left{
        width: 6px;
        height: 100%;
        top: 0;
        right: 100%;
        background-position: -3px 0px;
      }
    }
    &.thumb .stamp-frame {
      position: absolute;
      background-color: #fff;
      background: radial-gradient(
        transparent 0px,
        transparent 1.5px,
        white 1.5px,
        white
      );
      background-size: 5px 5px;
      &.top {
        width: calc(100% + 6px);
        height: 3px;
        bottom: 100%;
        left: -3px;
        background-position: 2px 2px;
      }
      &.bottom {
        width: calc(100% + 6px);
        height: 3px;
        top: 100%;
        left: -3px;
        background-position: 2px -4px;
      }
      &.right {
        width: 3px;
        height: 100%;
        top: 0;
        left: 100%;
        background-position: -9px 0px;
      }
      &.left {
        width: 3px;
        height: 100%;
        top: 0;
        right: 100%;
        background-position: 7px 0px;
      }
    }
    &.list .stamp-frame {
      position: absolute;
      background-color: #fff;
      background: radial-gradient(
        transparent 0px,
        transparent 1.4px,
        white 1.4px,
        white
      );
      background-size: 4px 4px;
      z-index: 10;
      &.top {
        width: calc(100% + 6px);
        height: 3px;
        bottom: 100%;
        left: -3px;
        background-position: 2px -2px;
      }
      &.bottom {
        width: calc(100% + 6px);
        height: 3px;
        top: 100%;
        left: -3px;
        background-position: 2px 1px;
      }
      &.right {
        width: 3px;
        height: 100%;
        top: 0;
        left: 100%;
        background-position: -3px 0px;
      }
      &.left {
        width: 3px;
        height: 100%;
        top: 0;
        right: 100%;
        background-position: 2px 0px;
      }
    }
    .stamp-img{
      max-width: 100%;
      max-height: 100%;
      line-height: 0;
      display: block;
      vertical-align: top;
      transform: scale(1.055);
    }
  }
</style>
