<template>
  <div class="album-item" ref="album" :class="lock?'lock':null" @click="click">
    <img class="album-cover-bg" :src="coverUrl" alt="">
    <div class="img-wrap">
      <StampWrap :imgSrc="coverUrl" :type="'thumb'" :padding="padding"></StampWrap>
    </div>
    <div class="album-name">{{name}}</div>
    <div class="album-year">{{year}}</div>
    <div class="album-complete" :class="type" v-if="completeCollected">
      <div class="cp-tag">齐</div>
      <div class="cp-tag">全</div>
      <div class="cp-tag">套</div>
    </div>
  </div>
</template>
<script>
import StampWrap from './StampWarp'

export default {
  name: 'AlbumWrap',
  components: {
    StampWrap
  },
  data () {
    return {
      padding: 22
    }
  },
  props: {
    year: String,
    name: String,
    completeCollected: Boolean,
    lock: Boolean,
    coverUrl: String,
    type: String
  },
  mounted () {
    if (this.type === 'trade') {
      this.padding = 14
    }
  },
  methods: {
    click () {
      this.$emit('click')
    }
  }
}
</script>
<style lang="less">
  @themeColor: #01ce7e;
  @rowHeight: 140px;
  @albumRowHeight: 120px;
  .year-folder{
    clear: both;
    text-align: left;
    width: 100%;
    height: 20px;
    font-size: 13px;
    margin-bottom: 10px;
    &:before{
      content: "";
      display: inline-block;
      width: 10px;
      height: 12px;
      margin-right: 5px;
      background-image: url(/static/ui/triangle.png);
      background-size: cover;
      transition: all 0.3s ease-in-out;
      transform: rotate(90deg)translate(1px,0);
    }
    &.close:before{
      transform: rotate(0)translate(1px,0);
    }
  }
  .album-item{
    position: absolute;
    width: 100%;
    height: 100%;
    color: #aaa;
    border-radius: 4px;
    overflow: hidden;
    box-shadow: 0 0 1px 2px #0000001e;
    transform: scale3d(1,1,1);
    transition: all 0.1s ease-in-out;
    &.lock{
      .img-wrap{
        filter: grayscale(100%);
      }
      .album-cover-bg{
        filter: blur(15px)grayscale(100%);
      }
      &:before{
        content: "";
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        background-color: #00000099;
        z-index: 100;
      }
      &:after{
        content: "";
        position: absolute;
        width: 39px;
        height: 45px;
        top: 50%;
        left: 50%;
        margin-left: -19.5px;
        margin-top: -27.5px;
        background-image: url(/static/ui/bookLocked.icon.png);
        background-repeat: no-repeat;
        background-size: contain;
        z-index: 120;
      }
    }
    .img-wrap{
      position: absolute;
      display: flex;
      align-items: center;
      justify-content: center;
      position: relative;
      width: 100%;
      height: 70px;
      top: 0;
      left: 0;
    }
    .album-cover-bg{
      position: absolute;
      top: -30%;
      left: -30%;
      width: 160%;
      height: 160%;
      object-fit: cover;
      filter: blur(15px)brightness(70%);
      transform: scale3d(1,1,1);
      opacity: 1;
      pointer-events: none;
    }
    .album-name{
      position: absolute;
      width: calc(100% - 10px);
      left: 5px;
      bottom: 21px;
      font-size: 12px;
      color: #ffffff;
      transform: scale(0.8);
    }
    .album-year{
      position: absolute;
      width: 100%;
      height: 20px;
      bottom: 5px;
      font-size: 12px;
      color: #ffffff;
      transform: scale(0.7);
    }
    .album-complete{
      position: absolute;
      top: 0;
      left: 6px;
      padding: 1px;
      padding-top: 14px;
      background-color: #ffdb19;
      border-radius: 0 0 7px 7px;
      .cp-tag{
        width: 12px;
        height: 10px;
        line-height: 0;
        font-size: 12px;
        transform: scale(0.72);
        color: #000;
      }
      &.trade{
        top: 60px;
        left: 0;
        width: 70%;
        height: 8px;
        border-radius: 0;
        padding: 0 0 0 6px;
        background-color: transparent;
        &:before{
          content: '';
          display: block;
          position: absolute;
          width: 100%;
          height: 100%;
          top: 0;
          left: 0;
          background-color: #ffdb19;
          transform: skewX(-25deg)translateX(-5px);
        }
        &:after{
          content: '';
          display: block;
          position: absolute;
          width: 100%;
          height: 100%;
          top: 100%;
          left: 0;
          background-color: #ffdb19;
          transform: skewX(25deg)translateX(-5px);
        }
        .cp-tag{
          position: relative;
          float: left;
          width: 9px;
          height: 20px;
          font-size: 12px;
          margin-top: 6px;
          transform: scale(0.72);
          color: #000;
          z-index: 10;
        }
      }
    }
  }
</style>
