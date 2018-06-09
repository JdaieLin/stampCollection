<template>
  <div class="mid-center">
    <div class="album-page" v-if="pageLevel == 'album'">
      <div class="tabs">
        <div class="tab-item" @click="switchTab('collected')" :class="currentAlbumTab == 'collected' ? 'active' : null">已收藏</div>
        <div class="tab-item" @click="switchTab('bought')" :class="currentAlbumTab == 'bought' ? 'active' : null">已购买</div>
        <div class="tab-item" @click="switchTab('year')" :class="currentAlbumTab == 'year' ? 'active' : null">年鉴</div>
      </div>
      <div class="album-page--wrapper">
        <div v-if="currentAlbumTab != 'year'" class="sort">
          <div class="sort-option" :class="sort=='year'?'active':null" @click="swithSort('year')">发布年份</div>
          <div class="sort-option" :class="sort=='recent'?'active':null" @click="swithSort('recent')">最近查看</div>
        </div>
        <div class="album-page--scroll">
          <template v-for="(album, index) in currentAlbum" v-if="!(currentAlbumTab != 'year' && album.lock)">
            <div class="year-folder" v-if="album.yearStart && currentAlbumTab=='year'" :key="'time_'+index">{{album.year}}</div>
            <div class="album-wrap" :key="'album_'+index">
              <AlbumWrap :name="album.name"
                         :year="album.year"
                         :lock="album.lock"
                         :coverUrl="album.src"
                         @click="checkAlbum"></AlbumWrap>
            </div>
          </template>
        </div>
      </div>
    </div>
    <div class="serial-page" v-if="pageLevel == 'serial'">
      <div class="back" @click="backToAlbum"></div>
      <div class="title">飞行系列</div>
      <div class="serial-page--wrapper">
        <div class="serial-page--shelf">
          <div class="stamp-item" v-for="(stamp, index) in currentSerial" :key="'stamp_' + index">
            <div v-if="stamp.multiple" class="multiple-wrap" :key="'stamp_img_' + index">
              <div class="img-wrap" v-for="(stp, index) in stamp.list" :key="'stp_img_' + index" @click="multipleClick">
                <StampWrap :imgSrc="stp.src"
                           :type="'list'"
                           :frame="true"
                           :level="stp.level"></StampWrap>
              </div>
            </div>
            <div v-else class="img-wrap" @click="openSingle">
              <StampWrap :imgSrc="stamp.src" :type="'list'" :frame="true" :key="'stamp_img_' + index"></StampWrap>
            </div>
            <div v-if="stamp.expire" class="expire">{{stamp.expire}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="multiple-page" v-if="multiple && !single" @click="closeMultiple">
      <div class="stamp-transformer" ref="mItem" v-for="(stamp, index) in currentMultiple" :key="'multiple_' + index">
        <div class="stamp-item">
          <div class="img-wrap" @click="openSingle">
            <StampWrap :imgSrc="stamp.src" :type="'list'" :frame="true"></StampWrap>
          </div>
          <div v-if="stamp.expire" class="expire">{{stamp.expire}}</div>
        </div>
      </div>
    </div>
    <div class="single-page" v-if="single" @click="closeSingle">
      <div class="single-wrap">
        <div class="img-warp">
          <StampWrap :imgSrc="currentStamp.src" :type="'large'" :frame="true" :padding="60" :v-padding="30"></StampWrap>
        </div>
        <div class="single-detail">
          <div class="stamp-title">{{currentStamp.name}}, {{currentStamp.year}}</div>
          <div class="stamp-counts">
            <span class="level">品相：{{currentStamp.level}}</span>
            <span class="type">{{currentStamp.type}}</span>
            <span class="num">剩余量：{{currentStamp.rest}}/{{currentStamp.total}}</span>
          </div>
          <div class="stamp-desc">{{currentStamp.desc}}</div>
          <div class="stamp-expire" v-if="currentStamp.expire">
            <span class="expire-btn buy" @click="buyStamp">购买</span>
            <span class="expire-btn extend" @click="extendExpire">延期</span>
            <span class="expire-text">{{currentStamp.expire}}</span>
          </div>
        </div>
        <div class="single-close" @click="closeSingle"></div>
      </div>
    </div>
  </div>
</template>

<script>
import StampWrap from '../components/StampWarp'
import AlbumWrap from '../components/AlbumWrap'

export default {
  name: 'Album',
  components: {
    StampWrap,
    AlbumWrap
  },
  data () {
    return {
      currentAlbumTab: 'collected',
      pageLevel: 'album',
      sort: 'recent',
      multiple: false,
      single: false,
      currentStamp: {
        src: '/static/img/demo7.jpg',
        name: '飞天一号',
        year: '2018',
        desc: '《中国恐龙》特种邮票则是中国邮政首次发行以恐龙为主题的邮票。该套邮票以在我国境内发掘的恐龙化石复原图为主要内容，整体画面富有动感，兼具科学性和艺术性。',
        level: 89,
        total: 40000,
        rest: 23083,
        type: '特种邮票',
        expire: '2018.5.31 到期'
      },
      currentMultiple: [
        {
          src: '/static/img/demo1.jpg',
          name: 'falcon',
          year: '2017',
          expire: '2018.5.31 到期'
        },
        {
          src: '/static/img/demo1.jpg',
          name: 'falcon',
          year: '2017'
        },
        {
          src: '/static/img/demo1.jpg',
          name: 'falcon',
          year: '2017'
        }
      ],
      currentSerial: [{
        src: '/static/img/demo7.jpg',
        name: '飞天一号',
        year: '2018'
      },
      {
        src: '/static/img/demo4.jpg',
        name: '飞天二号珍藏版',
        year: '2018',
        expire: '2018.5.31 到期'
      },
      {
        multiple: true,
        list: [
          {
            src: '/static/img/demo1.jpg',
            name: 'falcon',
            year: '2017',
            expire: '2018.5.31 到期'
          },
          {
            src: '/static/img/demo1.jpg',
            name: 'falcon',
            year: '2017'
          },
          {
            src: '/static/img/demo1.jpg',
            name: 'falcon',
            year: '2017'
          }
        ]
      },
      {
        src: '/static/img/1.png',
        name: 'falcon',
        year: '2017',
        expire: '2018.5.31 到期'
      },
      {
        src: '/static/img/3.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/5.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/1.png',
        name: 'falcon',
        year: '2017',
        expire: '2018.5.31 到期'
      },
      {
        src: '/static/img/3.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/5.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/1.png',
        name: 'falcon',
        year: '2017',
        expire: '2018.5.31 到期'
      },
      {
        src: '/static/img/3.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/5.png',
        name: 'falcon',
        year: '2017'
      },
      {
        src: '/static/img/6.png',
        name: 'falcon',
        year: '2016',
        expire: '2018.5.31 到期'
      },
      {
        src: '/static/img/2.png',
        name: 'falcon',
        year: '2018'
      },
      {
        src: '/static/img/7.png',
        name: 'falcon',
        year: '2018'
      },
      {
        src: '/static/img/4.png',
        name: 'falcon',
        year: '2018'
      }],
      currentAlbum: [{
        src: '/static/img/demo1.jpg',
        name: '邮册名称',
        year: '2018',
        yearStart: '2018'
      },
      {
        src: '/static/img/demo2.jpg',
        name: '香港回归祖国二十周年',
        year: '2017',
        yearStart: '2017',
        completeCollected: true
      },
      {
        src: '/static/img/demo3.jpg',
        name: 'stamp2',
        year: '2017'
      },
      {
        src: '/static/img/demo4.jpg',
        name: 'stamp2lock',
        year: '2017',
        lock: true
      },
      {
        src: '/static/img/demo5.jpg',
        name: 'stamp2',
        year: '2015',
        yearStart: '2015'
      },
      {
        src: '/static/img/demo7.jpg',
        name: 'stamp2',
        year: '2015'
      },
      {
        src: '/static/img/demo1.jpg',
        name: 'stamp2',
        year: '2015'
      },
      {
        src: '/static/img/demo4.jpg',
        name: 'stamp2',
        year: '2014',
        yearStart: '2014',
        lock: true
      },
      {
        src: '/static/img/demo5.jpg',
        name: 'stamp2',
        year: '2014',
        yearStart: '2015'
      },
      {
        src: '/static/img/demo7.jpg',
        name: 'stamp2',
        year: '2014'
      },
      {
        src: '/static/img/demo5.jpg',
        name: 'stamp2',
        year: '2014',
        yearStart: '2015'
      },
      {
        src: '/static/img/demo7.jpg',
        name: 'stamp2',
        year: '2014'
      },
      {
        src: '/static/img/demo5.jpg',
        name: 'stamp2',
        year: '2014',
        yearStart: '2015'
      },
      {
        src: '/static/img/demo7.jpg',
        name: 'stamp2',
        year: '2014'
      }],
      boughtAlbums: [],
      yearAlbums: []
    }
  },
  mounted () {
  },
  methods: {
    switchTab (tab) {
      this.currentAlbumTab = tab
    },
    checkAlbum (e) {
      this.pageLevel = 'serial'
    },
    backToAlbum (e) {
      this.pageLevel = 'album'
    },
    swithSort (sort) {
      this.sort = sort
    },
    multipleClick (e) {
      let that = this
      this.openMultiple(e)
      let pageX = e.pageX
      let pageY = e.pageY
      setTimeout(function () {
        let elems = that.$refs.mItem
        let offsetHeight = elems[0].offsetHeight
        let offsetWidth = elems[0].offsetWidth
        for (let i in elems) {
          let offsetTop = elems[i].getBoundingClientRect().y
          let offsetLeft = elems[i].getBoundingClientRect().x
          let ox = (pageX - offsetWidth / 2) - offsetLeft
          let oy = (pageY - offsetHeight / 2) - offsetTop
          elems[i].style.transform = 'translate(' + ox + 'px,' + oy + 'px)scale(0.7)'
          elems[i].style.webkitTransform = 'translate(' + ox + 'px,' + oy + 'px)scale(0.7)'
        }
        setTimeout(function () {
          for (let i in elems) {
            elems[i].style.transition = 'all 0.3s ease-out'
            elems[i].style.webkitTransition = 'all 0.3s ease-out'
            elems[i].style.transform = 'translate(0,0)'
            elems[i].style.webkitTransform = 'translate(0,0)'
            elems[i].style.opacity = '1'
          }
        }, 0)
      }, 300)
    },
    openMultiple (e) {
      this.multiple = true
    },
    openSingle (e) {
      this.single = true
      e.stopPropagation()
    },
    closeMultiple (e) {
      this.multiple = false
    },
    closeSingle (e) {
      this.single = false
      let that = this
      setTimeout(function () {
        if (that.multiple) {
          let elems = that.$refs.mItem
          for (let i in elems) {
            elems[i].style.transition = 'all 0.3s ease-out'
            elems[i].style.webkitTransition = 'all 0.3s ease-out'
            elems[i].style.opacity = '1'
          }
        }
      }, 200)
    },
    buyStamp (e) {
      e.stopPropagation()
    },
    extendExpire (e) {
      e.stopPropagation()
    }
  }
}
</script>

<style scoped lang="less">
@import '../less/common.less';
.mid-center{
  width: 100%;
  height: calc(100% - 60px);
}
.album-page{
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: hidden;
  &:before{
    content: "";
    position: absolute;
    top: 60px;
    left: 16px;
    width: calc(100% - 32px);
    height: 4px;
    background-color: @themeColor;
    z-index: 100;
    border-radius: 8px 8px 0 0;
  }
  .tabs{
    position: absolute;
    top: 30px;
    left: 32px;
    z-index: 10;
    .tab-item{
      float: left;
      display: flex;
      position: relative;
      align-items: center;
      justify-content: center;
      width: 80px;
      height: 25px;
      margin-top: 5px;
      margin-right: 5px;
      background-color: #cdcdcd;
      border-radius: 4px 4px 0 0;
      font-size: 12px;
      color: #666;
      overflow: hidden;
      &:after{
        content: '';
        position: absolute;
        width: 100%;
        height: 1px;
        bottom: -1px;
        left: 0;
        box-shadow: 0 0 10px 5px #0000001e;
      }
      &.active{
        background-color: @themeColor;
        color: white;
        height: 30px;
        margin-top: 0;
        z-index: 20;
        &:after{
          display: none;
        }
      }
    }
  }
  .album-page--wrapper{
    position: absolute;
    width: calc(100% - 32px);
    height: calc(100% - 60px);
    top: 60px;
    left: 16px;
    padding: 16px 6px 80px 16px;
    border-radius: 8px 8px 0 0;
    background-color: #fff;
    overflow-y: scroll;
    box-shadow: 0 0 10px 3px #0000002e;
    -webkit-overflow-scrolling: touch;
    .sort{
      position: relative;
      width: 100%;
      height: 24px;
      margin-bottom: 10px;
      .sort-option{
        float: right;
        width: 80px;
        height: 24px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 12px;
        color: @themeColor;
        border: 1px solid @themeColor;
        border-radius: 12px;
        margin-right: 10px;
        &.active{
          background-color: @themeColor;
          color: #fff;
        }
      }
    }
    .album-page--scroll{
      width: 100%;
      display: flex;
      flex-wrap:wrap;
    }
    .album-wrap{
      position: relative;
      width: calc(33.3% - 10px);
      height: @albumRowHeight;
      margin-right: 10px;
      margin-bottom: 10px;
    }
  }
}
.serial-page{
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: hidden;
  .back{
    position: absolute;
    width: 20px;
    height: 20px;
    top: 18px;
    left: 16px;
    z-index: 10;
    background-image: url(/static/ui/back.png);
    background-size: contain;
    background-repeat: no-repeat;
  }
  .title{
    position: absolute;
    width: 100%;
    height: 30px;
    top: 16px;
    color: #333;
  }
  .serial-page--wrapper{
    position: absolute;
    width: calc(100% - 32px);
    height: calc(100% - 50px);
    top: 50px;
    left: 16px;
    padding: 0 6px 80px 16px;
    border-radius: 8px 8px 0 0;
    background-color: #fff;
    overflow: auto;
    box-shadow: 0 0 10px 3px #0000003e;
    -webkit-overflow-scrolling: touch;
    .serial-page--shelf{
      position: relative;
      width: calc(100% + 22px);
      min-height: 100%;
      margin-left: -16px;
      position: relative;
      border-radius: 8px 8px 0 0;
      clear: both;
      &:before{
        content: '';
        display: block;
        position: absolute;
        width: 100%;
        height: 100%;
        pointer-events: none;
        background: linear-gradient(0deg, black 35px, transparent 37px);
        background-size: 100% @stampRowHeight;
        opacity: 0.1;
        z-index: 10;
      }
      &:after{
        content: "";
        display: block;
        clear: both;
      }
      .stamp-item{
        position: relative;
        float: left;
        width: 33.3%;
        height: @stampRowHeight;
        transform: scale3d(1,1,1);
        .img-wrap{
          display: flex;
          flex-direction: column-reverse;
          position: absolute;
          width: 100%;
          height: 110px;
          bottom: 6px;
        }
        .expire{
          position: absolute;
          bottom: -20px;
          width: 100%;
          font-size: 12px;
          color: #666;
          transform: scale(0.8);
        }
        .multiple-wrap{
          transition: all 0.3s ease;
          .img-wrap{
            opacity: 0;
            transition: all 0.3s ease;
            &:nth-child(1){
              opacity: 1;
              z-index: 10;
            }
            &:nth-child(2){
              opacity: 1;
              transform: scale(0.9)translateY(-7px);
              z-index: 9;
            }
            &:nth-child(3){
              opacity: 1;
              transform: scale(0.8)translateY(-14px);
              z-index: 8;
            }
          }
        }
      }
    }
  }
}
.single-page{
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 2000;
  background-color: #000000cc;
  .single-wrap{
    position: absolute;
    top: 8%;
    left: 10px;
    width: calc(100% - 20px);
    height: 80%;
    border-radius: 6px;
    background-color: #fff;
    box-shadow: 0 0 10px 5px #0000008e;
  }
  .img-warp{
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    top: 0;
    left: 0;
    width: 100%;
    height: calc(100% - 200px);
    background-color: #e3e3e3;
    border-radius: 6px 6px 0 0;
  }
  .single-detail{
    position: absolute;
    width: 100%;
    height: 200px;
    bottom: 0;
    left: 0;
    padding: 16px;
    .stamp-title{
      width: 100%;
      height: 22px;
      font-size: 16px;
      font-weight: bold;
      text-align: left;
    }
    .stamp-counts{
      font-size: 12px;
      color: #666;
      .level{
        float: left;
      }
      .num{
        float: right;
      }
    }
    .stamp-desc{
      width: 100%;
      height: 90px;
      padding: 10px 0;
      font-size: 12px;
      text-align: left;
    }
    .stamp-expire{
      color: #666;
      width: 100%;
      font-size: 12px;
      span{
        float: right;
        line-height: 24px;
      }
      .expire-btn{
        display: flex;
        justify-content: center;
        align-items: center;
        width: 50px;
        height: 24px;
        border-radius: 12px;
        border: 1px solid @themeColor;
        margin-left: 6px;
        &.extend{
          color: @themeColor;
          &:active{
            background-color: darken(#fff, 6%);
          }
        }
        &.buy{
          background-color: @themeColor;
          &:active{
            background-color: darken(@themeColor, 6%);
          }
          color: #fff;
        }
      }
    }
  }
  .single-close{
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
.multiple-page{
  position: fixed;
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  padding-bottom: 80px;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  background-color: #000000cc;
  .stamp-transformer{
    position: relative;
    width: 30%;
    height: 80px;
    opacity: 0;
  }
  .stamp-item{
    transform: scale3d(1,1,1);
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    .img-wrap{
      display: flex;
      flex-direction: column-reverse;
      position: absolute;
      width: 100%;
      height: 110px;
      bottom: 6px;
    }
    .expire{
      position: absolute;
      bottom: -20px;
      width: 100%;
      font-size: 12px;
      color: #aaa;
      transform: scale(0.8);
    }
  }
}
</style>
