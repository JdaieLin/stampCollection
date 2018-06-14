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
              <AlbumWrap :name="album.set_name"
                         :year="album.year"
                         :coverUrl="album.image"
                         :lock="album.lock"
                         @click="checkAlbum(album.set, album.lock)"></AlbumWrap>
            </div>
          </template>
        </div>
      </div>
    </div>
    <div class="serial-page" v-if="pageLevel == 'serial'">
      <div class="back" @click="backToAlbum"></div>
      <div class="action-btn" v-if="currentAlbumTab=='bought'" @click="openSellModal">出售全套</div>
      <div class="title">{{setName}}</div>
      <div class="serial-page--wrapper">
        <div class="serial-page--shelf">
          <div class="stamp-item" v-for="(stamp, index) in currentSerial" :key="'stamp_' + index">
            <div v-if="stamp.multiple" class="multiple-wrap" :key="'stamp_img_' + index">
              <div class="img-wrap" v-for="(stp, index) in stamp.list" :key="'stp_img_' + index" @click="multipleClick($event, stamp.list)">
                <StampWrap :imgSrc="stp.image"
                           :type="'list'"
                           :frame="true"
                           :level="stp.score"
                           :stamp="stamp"></StampWrap>
              </div>
            </div>
            <div v-else class="img-wrap" @click="openSingle(stamp)">
              <StampWrap :imgSrc="stamp.image"
                         :type="'list'"
                         :frame="true"
                         :key="'stamp_img_' + index"
                         :level="stamp.score"
                         :stamp="stamp"></StampWrap>
            </div>
            <div v-if="stamp.expire" class="expire">{{stamp.expire}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="multiple-page" v-if="multiple && !single" @click="closeMultiple">
      <div class="stamp-transformer" ref="mItem" v-for="(stamp, index) in currentMultiple" :key="'multiple_' + index">
        <div class="stamp-item">
          <div class="img-wrap" @click="openSingle(stamp)">
            <StampWrap :imgSrc="stamp.image"
                       :type="'list'"
                       :frame="true"
                       :level="stamp.score"
                       :stamp="stamp"></StampWrap>
          </div>
          <div v-if="stamp.expire" class="expire">{{stamp.expire}}</div>
        </div>
      </div>
    </div>
    <div class="single-page" v-if="single" @click="closeSingle">
      <div class="single-wrap">
        <div class="img-warp">
          <StampWrap :imgSrc="currentStamp.image"
                     :level="currentStamp.score"
                     :type="'large'"
                     :frame="true"
                     :padding="60"
                     :v-padding="30"
                     :stamp="currentStamp"></StampWrap>
        </div>
        <div class="single-detail">
          <div class="stamp-title">{{currentStamp.name!='-'?currentStamp.name:currentStamp.set_name}}, {{currentStamp.year}}</div>
          <div class="stamp-counts">
            <span class="level" v-if="currentAlbumTab!='year'">品相：{{currentStamp.score}}</span>
            <span class="type">{{currentStamp.type}}</span>
            <span class="num">剩余量：{{currentStamp.remain}}/{{currentStamp.volume}}</span>
          </div>
          <div class="stamp-desc">{{currentStamp.desc}}</div>
          <template v-if="currentAlbumTab!='year'">
            <div class="stamp-action" v-if="currentStamp.boughtType">
              <span class="action-btn green" @click="openSellModal">交易</span>
              <span class="action-btn white wide" @click="openRecycleModal">出售给平台</span>
              <span class="action-text green" v-if="currentStamp.boughtType == 'collect'">已获得（收藏满一个月）</span>
              <span class="action-text green" v-if="currentStamp.boughtType == 'buy'">已于{{currentStamp.buyTime}}购买</span>
            </div>
            <div class="stamp-action" v-else-if="currentStamp.trading">
              <span class="action-text">交易中</span>
            </div>
            <div class="stamp-action" v-else>
              <span class="action-btn green" @click="openBuyModal">购买</span>
              <!--<span class="action-btn white" @click="extendExpire">延期</span>-->
              <span class="action-text">{{currentStamp.expire}}</span>
            </div>
          </template>
        </div>
        <div class="single-close" @click="closeSingle"></div>
      </div>
    </div>
    <div class="modal unlock-modal" v-if="unlockModal">
      <div class="modal-center">
        <div class="modal-text">
          <p>将花费铜钱</p>
          <p><span class="coins-icon"></span> x 1000</p>
          <p>解锁该邮册</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closeUnlockModal">取消</div>
          <div class="theme-btn action-btn green" @click="unlockAlbum">确定</div>
        </div>
      </div>
    </div>
    <div class="modal buy-modal" v-if="buyModal">
      <div class="modal-center">
        <div class="modal-text">
          <p>将花费以太币</p>
          <p><span class="eth-icon"></span> x {{Math.ceil(currentStamp.floor*1000)/1000}}</p>
          <p>购买该邮票</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closebuyModal">取消</div>
          <div class="theme-btn action-btn green" @click="buyStamp">确定</div>
        </div>
      </div>
    </div>
    <div class="modal recycle-modal" v-if="recycleModal">
      <div class="modal-center">
        <div class="modal-text">
          <p>出售给平台获得元宝</p>
          <p><span class="ingot-icon"></span> x {{Math.ceil(currentStamp.floor)}}</p>
          <p class="tip">出售给平台只能获得元宝</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closeRecycleModal">取消</div>
          <div class="theme-btn action-btn green" @click="recycleStamp">确定</div>
        </div>
      </div>
    </div>
    <div class="modal recycle-modal" v-if="sellModal">
      <div class="modal-center">
        <div class="modal-text">
          <p>将以下列价格出售该邮票</p>
          <p><span class="eth-icon"></span> x <input type="text" :value="sellPrice" @change="priceChange"></p>
          <p class="tip">平台将收取售价1%作为交易手续费</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closeSellModal">取消</div>
          <div class="theme-btn action-btn green" @click="sellStamp">确定</div>
        </div>
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
      set: null,
      currentStamp: {},
      currentMultiple: [],
      unlockModal: false,
      unlockId: 0,
      buyModal: false,
      recycleModal: false,
      sellIds: [],
      sellPrice: 0,
      sellDefaultPrice: 100,
      sellModal: false
    }
  },
  mounted () {
  },
  computed: {
    currentAlbum () {
      if (this.currentAlbumTab === 'year') {
        return this.$store.state.Album.yearBook
      } else if (this.currentAlbumTab === 'bought') {
        return this.$store.state.Album.boughtBook
      } else if (this.currentAlbumTab === 'collected') {
        return this.$store.state.Album.collectBook
      } else {
        return []
      }
    },
    currentSerial () {
      if (this.set == null) return []
      let serialArray = []
      let allStamp
      if (this.currentAlbumTab === 'year') {
        allStamp = this.$store.state.Album.allStamp
      } else if (this.currentAlbumTab === 'collected') {
        allStamp = this.$store.state.Album.collectList
      } else {
        allStamp = this.$store.state.Album.finalBoughtList
      }
      for (let i in allStamp) {
        if (allStamp[i].set === this.set) {
          serialArray.push(allStamp[i])
        }
      }
      return serialArray
    }
  },
  methods: {
    switchTab (tab) {
      this.currentAlbumTab = tab
    },
    checkAlbum (set, lock) {
      if (lock) {
        this.unlockModal = true
        for (let i in this.$store.state.Album.yearBook) {
          if (this.$store.state.Album.yearBook[i].set === set) {
            this.unlockId = this.$store.state.Album.yearBook[i].id
          }
        }
      } else {
        this.set = set
        this.setName = this.$store.state.Album.yearBook[this.set].set_name
        this.pageLevel = 'serial'
      }
    },
    backToAlbum (e) {
      this.pageLevel = 'album'
    },
    swithSort (sort) {
      this.sort = sort
    },
    multipleClick (e, list) {
      this.currentMultiple = list
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
    openSingle (stamp) {
      this.currentStamp = stamp
      this.single = true
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
    closeUnlockModal () {
      this.unlockModal = false
    },
    unlockAlbum () {
      // 检测铜钱数量
      this.$store.dispatch('unlockBook', this.unlockId)
      this.unlockModal = false
      this.actionRefresh()
    },
    actionRefresh () {
      let tab = this.currentAlbumTab
      let that = this
      this.currentAlbumTab = ''
      setTimeout(function () {
        that.currentAlbumTab = tab
      }, 200)
    },
    openBuyModal (e) {
      this.buyModal = true
      e.stopPropagation()
    },
    closebuyModal () {
      this.buyModal = false
    },
    buyStamp (e) {
      this.buyModal = false
      this.$store.dispatch('albumBuyStamp', this.currentStamp)
      e.stopPropagation()
    },
    openRecycleModal (e) {
      this.recycleModal = true
      e.stopPropagation()
    },
    closeRecycleModal () {
      this.recycleModal = false
    },
    recycleStamp (e) {
      this.recycleModal = false
      this.$store.dispatch('albumRecycleStamp', this.currentStamp)
      e.stopPropagation()
    },
    openSellModal (e) {
      this.sellPrice = this.sellDefaultPrice
      this.sellModal = true
      let minPrice = 0
      if (this.single) {
        this.sellIds = [this.currentStamp.id]
        minPrice += this.currentStamp.floor
      } else {
        for (let i in this.currentSerial) {
          let item = this.currentSerial[i]
          if (item.multiple) {
            for (let j in item.list) {
              this.sellIds.push(item.list[j].id)
              minPrice += item.list[j].floor
            }
          } else {
            this.sellIds.push(item.id)
            minPrice += item.floor
          }
        }
      }
      this.sellPrice = Math.ceil(minPrice)
      e.stopPropagation()
    },
    closeSellModal () {
      this.sellModal = false
    },
    priceChange (e) {
      this.sellPrice = parseInt(e.target.value)
    },
    sellStamp () {
      this.$store.dispatch('createDeal', {
        stampIds: this.sellIds,
        price: this.sellPrice
      })
      this.actionRefresh()
      this.sellModal = false
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
  .action-btn{
    position: absolute;
    top: 16px;
    right: 16px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    width: 60px;
    height: 24px;
    border-radius: 12px;
    color: white;
    border: 1px solid @themeColor;
    background-color: @themeColor;
    z-index: 10;
    &:active{
      background-color: darken(@themeColor, 6%);
    }
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
    .stamp-action{
      color: #666;
      width: 100%;
      font-size: 12px;
      span{
        float: right;
        line-height: 24px;
      }
      .action-btn{
        display: flex;
        justify-content: center;
        align-items: center;
        width: 50px;
        height: 24px;
        border-radius: 12px;
        border: 1px solid @themeColor;
        margin-left: 6px;
        &.wide{
          width: 80px;
        }
        &.white{
          color: @themeColor;
          &:active{
            background-color: darken(#fff, 6%);
          }
        }
        &.green{
          background-color: @themeColor;
          &:active{
            background-color: darken(@themeColor, 6%);
          }
          color: #fff;
        }
      }
      .action-text{
        &.green{
          color: @themeColor;
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
.modal{
  position: fixed;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  background-color: #000000aa;
  z-index: 3000;
  .modal-center{
    position: relative;
    top: -30px;
    width: 240px;
    height: 180px;
    background-color: #fff;
    border-radius: 4px;
    .modal-text{
      position: relative;
      top: 10px;
      p{
        font-size: 14px;
        line-height: 14px;
        vertical-align: center;
        &.tip{
          font-size: 12px;
        }
        input{
          width: 40px;
          border: 1px solid #ccc;
          border-radius: 4px;
          padding: 4px;
        }
      }
      .coins-icon{
        display: inline-block;
        width: 24px;
        height: 24px;
        background-image: url(/static/ui/coin.png);
        background-size: contain;
        background-repeat: no-repeat;
      }
      .eth-icon{
        display: inline-block;
        width: 24px;
        height: 24px;
        background-image: url(/static/ui/ether.png);
        background-size: contain;
        background-repeat: no-repeat;
      }
      .ingot-icon{
        display: inline-block;
        width: 30px;
        height: 18px;
        background-image: url(/static/ui/ignotMini.png);
        background-size: contain;
        background-position: center;
        background-repeat: no-repeat;
      }
    }
    .btn-wrap{
      position: absolute;
      width: 100%;
      bottom: 20px;
      display: flex;
      align-items: center;
      justify-content: center;
      .action-btn{
        width: 90px;
        height: 30px;
        margin: 4px;
        display: flex;
        border-radius: 15px;
      }
    }
  }
}
</style>
