<template>
  <div class="mid-center">
    <div class="trade-tab">
      <div class="tab-item" :class="currentTab=='buy'?'active':null" @click="switchTab('buy')">购买</div>
      <div class="tab-item" :class="currentTab=='sell'?'active':null" @click="switchTab('sell')">出售中</div>
    </div>
    <div class="trade-page">
      <div class="category">
        <div v-if="currentTab=='buy'" class="group">
          <div class="group-label">热门</div>
          <div class="group-item" :class="currentTag=='hot_stamp'?'active':null" @click="clickTag('hot_stamp')">热门单张</div>
          <div class="group-item" :class="currentTag=='hot_album'?'active':null" @click="clickTag('hot_album')">热门套票</div>
        </div>
        <div v-if="currentTab=='buy'" class="group">
          <div class="group-label">系列</div>
          <div class="group-item set"
               v-for="(item, key) in serialTags"
               :key="'tag_'+key"
               :class="currentTag==item.set?'active':null"
               @click="clickTag(item.set)">{{item.set_name}}</div>
        </div>
        <div v-if="currentTab=='sell'" class="group">
          <div class="group-label">出售中</div>
          <div class="group-item" :class="currentTag=='onsell_stamp'?'active':null" @click="clickTag('onsell_stamp')">单张</div>
          <div class="group-item" :class="currentTag=='onsell_album'?'active':null" @click="clickTag('onsell_album')">套票</div>
        </div>
        <!--<div v-if="currentTab=='sell'" class="group">-->
          <!--<div class="group-label">已出售</div>-->
          <!--<div class="group-item" :class="currentTag=='sold_stamp'?'active':null" @click="clickTag('sold_stamp')">单张</div>-->
          <!--<div class="group-item" :class="currentTag=='sold_album'?'active':null" @click="clickTag('sold_album')">套票</div>-->
        <!--</div>-->
      </div>
      <div class="trade-wrap">
        <div class="trade-wrap--scroll">
          <template v-if="currentTag.indexOf('_album')>0">
            <!--邮册排列-->
            <div v-for="(item, index) in tradeDisplayAlbum" :key="'trade_' + index" class="trade-item list" @click="toggleMultiple(true, item.stamps)">
              <div class="image-wrap">
                <AlbumWrap :coverUrl="item.image"
                           :type="'trade'"
                           :completeCollected="item.completeCollected"></AlbumWrap>
              </div>
              <div class="stamp-detail">
                <div class="name">{{item.set_name}}</div>
                <div class="type">套票</div>
                <div class="price">{{item.price}}</div>
              </div>
              <div class="seller-detail" v-if="!currentTag.indexOf('sold_')">
                <div class="seller-head"></div>
                <div class="seller-name">seller</div>
              </div>
              <div v-if="currentTab=='buy'" class="btn buy" @click="openBuyModal($event, { tradeId: item.trade_id, price: item.price })">购买</div>
              <div v-if="currentTag.indexOf('onsell_') == 0" class="btn buy" @click="openCancelModal($event,item.trade_id)">下架</div>
              <div v-if="currentTag.indexOf('sold_') == 0" class="btn sold">已出售</div>
            </div>
          </template>
          <template v-else-if="currentTag=='hot_stamp'">
            <!--三张邮票排列-->
            <div v-for="(item, index) in tradeDisplayStamp" :key="'trade_' + index" class="trade-item stamp" @click="clickTag(item.set)">
              <div class="image-wrap album">
                <StampWrap :imgSrc="item.image"
                           :type="'list'"
                           :frame="true"
                           :level="item.score"
                           :padding="2"></StampWrap>
              </div>
            </div>
          </template>
          <template v-else>
            <!--邮票列表排列-->
            <div v-for="(item, index) in tradeDisplayStamp" :key="'trade_' + index" class="trade-item list">
              <div class="stamp-list-item" @click="toggleSingle(true, item)">
                <div class="image-wrap">
                  <StampWrap :imgSrc="item.image"
                             :type="'list'"
                             :frame="true"
                             :level="item.score"
                             :padding="6"
                             :v-padding="1"
                             :stamp="item"></StampWrap>
                </div>
                <div class="stamp-detail">
                  <div class="name">{{item.name}}</div>
                  <div class="level">品相:{{item.score}}</div>
                  <div class="type">邮册:{{item.set_name}}</div>
                  <div class="price">{{item.price}}</div>
                </div>
                <div class="seller-detail" v-if="!currentTag.indexOf('sold_')">
                  <div class="seller-head"></div>
                  <div class="seller-name">seller</div>
                </div>
                <div v-if="currentTab=='buy'" class="btn buy" @click="openBuyModal($event, { tradeId: item.trade_id, price: item.price })">购买</div>
                <div v-if="currentTag.indexOf('onsell_') == 0" class="btn buy" @click="openCancelModal($event, item.trade_id)">下架</div>
                <div v-if="currentTag.indexOf('sold_') == 0" class="btn sold">已出售</div>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>
    <div class="single-warp" v-if="showSingle" @click="toggleSingle(false)">
      <div class="single-center">
        <StampWrap :imgSrc="tradeDisplaySingle.image"
                   :type="'large'"
                   :frame="true"
                   :level="tradeDisplaySingle.score"
                   :padding="2"
                   :stamp="tradeDisplaySingle"></StampWrap>
      </div>
    </div>
    <div class="multiple-warp" v-if="showMultiple" @click="toggleMultiple(false)">
      <div class="multiple-center">
        <div class="multiple-content">
          <div v-for="(item, index) in tradeDisplayMulti" :key="'multi_' + index" class="multi-item" >
            <div class="image-wrap">
              <StampWrap :imgSrc="item.image"
                       :type="'large'"
                       :frame="true"
                       :level="item.score"
                       :padding="2"
                       :stamp="item"></StampWrap>
            </div>
            <div class="stamp-detail">
              <div class="name">{{item.name}}</div>
              <div class="level">品相：{{item.score}}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="modal cancel-modal" v-if="cancelModal">
      <div class="modal-center">
        <div class="modal-text">
          <br>
          <p>确定要下架这笔交易吗</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closeCancelModal">返回</div>
          <div class="theme-btn action-btn green" @click="cancelTrade">确定</div>
        </div>
      </div>
    </div>
    <div class="modal buy-modal" v-if="buyModal">
      <div class="modal-center">
        <div class="modal-text">
          <br>
          <p>将用以下价格进行购买</p>
          <p><span class="eth-icon"></span> x {{buyItem.price}}</p>
        </div>
        <div class="btn-wrap">
          <div class="theme-btn action-btn white" @click="closeBuylModal">取消</div>
          <div class="theme-btn action-btn green" @click="acceptTrade">确定</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import StampWrap from '../components/StampWarp'
import AlbumWrap from '../components/AlbumWrap'

export default {
  name: 'Trade',
  components: {
    StampWrap,
    AlbumWrap
  },
  data () {
    return {
      tradeDisplayMulti: [],
      tradeDisplaySingle: {},
      currentTab: 'buy',
      currentTag: 'hot_stamp',
      showSingle: false,
      showMultiple: false,
      buyTags: [],
      hotTags: [],
      hotStampFilter: null,
      cancelModal: false,
      cancelId: 0,
      buyModal: false,
      buyItem: {
        tradeId: 0,
        price: 0
      }
    }
  },
  computed: {
    tradeDisplayStamp () {
      if (this.currentTab === 'sell' && this.currentTag === 'onsell_stamp') {
        // 出售中邮票详情
        let stamps = this.$store.state.Trade.sellListSingle.map(i => {
          let stamp = i.stamps[0]
          stamp.price = i.price
          stamp.trade_id = i.id
          return stamp
        })
        return stamps
      } else if (this.currentTab === 'buy' && this.currentTag === 'hot_stamp') {
        // 购买热门邮票
        let stamps = this.$store.state.Trade.buyListSingle.map(i => {
          let stamp = i.stamps[0]
          stamp.price = i.price
          stamp.trade_id = i.id
          return stamp
        }).filter((i, k, a) => {
          return k === a.map(j => j.atlas_id).indexOf(i.atlas_id)
        })
        return stamps
      } else if (this.currentTab === 'buy' && this.currentTag.indexOf('hot_') < 0 && !this.hotStampFilter) {
        // tag筛选邮票详情
        let stamps = this.$store.state.Trade.buyListSingle.map(i => {
          let stamp = i.stamps[0]
          stamp.price = i.price
          stamp.trade_id = i.id
          return stamp
        }).filter(i => i.set === this.currentTag)
        return stamps
      }
      return []
    },
    tradeDisplayAlbum () {
      if (this.currentTab === 'sell' && this.currentTag === 'onsell_album') {
        let album = this.$store.state.Trade.sellListAlbum.map(i => {
          let album = i.stamps[0]
          album.price = i.price
          album.stamps = i.stamps
          album.trade_id = i.id
          return album
        })
        return album
      } else if (this.currentTab === 'buy' && this.currentTag === 'hot_album') {
        let album = this.$store.state.Trade.buyListAlbum.map(i => {
          let album = i.stamps[0]
          album.price = i.price
          album.stamps = i.stamps
          album.trade_id = i.id
          return album
        })
        return album
      }
      return []
    },
    serialTags () {
      let tags = this.$store.state.Trade.buyListSingle.map(i => {
        let stamp = i.stamps[0]
        return { set: stamp.set, set_name: stamp.set_name }
      }).filter((i, k, a) => a.map(j => j.set).indexOf(i.set) === k)
      return tags
    }
  },
  methods: {
    switchTab (name) {
      this.currentTab = name
      if (name === 'buy') {
        this.currentTag = 'hot_stamp'
      } else {
        this.currentTag = 'onsell_stamp'
      }
    },
    clickTag (set) {
      this.currentTag = set
    },
    toggleSingle (bool, stamp) {
      if (stamp) this.tradeDisplaySingle = stamp
      this.showSingle = bool
    },
    toggleMultiple (bool, stamps) {
      if (stamps) this.tradeDisplayMulti = stamps
      this.showMultiple = bool
    },
    closeCancelModal () {
      this.cancelModal = false
    },
    openCancelModal (e, tradeId) {
      e.stopPropagation()
      this.cancelId = tradeId
      this.cancelModal = true
    },
    cancelTrade () {
      this.$store.dispatch('cancelTrade', this.cancelId)
      this.cancelModal = false
    },
    openBuyModal (e, { tradeId, price }) {
      e.stopPropagation()
      console.log(tradeId, price)
      this.buyItem = {
        tradeId: tradeId,
        price: price
      }
      this.buyModal = true
    },
    closeBuylModal () {
      this.buyModal = false
    },
    acceptTrade () {
      this.$store.dispatch('acceptTrade', this.buyItem.tradeId)
      this.buyModal = false
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
  .trade-tab{
    position: fixed;
    top: 20px;
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
  .trade-page{
    position: fixed;
    top: 50px;
    left: 0;
    width: 100%;
    height: calc(100% - 110px);
    background-color: #fff;
    box-shadow: 0 0 10px 5px #0000001e;
    &:before{
      content: '';
      position: absolute;
      width: 100%;
      height: 3px;
      left: 0;
      background-color: @themeColor;
      z-index: 10;
    }
    .category{
      position: absolute;
      width: 28%;
      height: 100%;
      font-size: 12px;
      text-align: left;
      color: #444;
      overflow-y: auto;
      -webkit-overflow-scrolling: touch;
      .group{
        padding: 10px 10px 0 20px;
        .group-label{
          position: relative;
          font-size: 14px;
          margin: 5px 0;
          &:before{
            content: '';
            position: absolute;
            width: 10px;
            height: 10px;
            top: 5px;
            left: -14px;
            background-image: url(/static/ui/triangle.png);
            background-size: contain;
            background-repeat: no-repeat;
            transform: rotate(90deg);
          }
        }
        .group-item{
          margin: 5px 0;
          &.set{
            margin-bottom: 10px;
          }
          &.active{
            font-weight: bold;
            color: @themeColor;
          }
        }
      }
    }
    .trade-wrap{
      position: absolute;
      width: 72%;
      height: 100%;
      top: 0;
      right: 0;
      padding: 16px;
      overflow-y: auto;
      -webkit-overflow-scrolling: touch;
      .trade-wrap--scroll{
        display: flex;
        flex-wrap:wrap;
        width: 100%;
      }
      .trade-item{
        .stamp-list-item{
          height: @stampRowHeightTrade;
        }
        &.stamp{
          position: relative;
          width: calc(33% - 10px);
          height: @stampRowHeightTrade;
          margin-right: 10px;
          margin-bottom: 20px;
          .image-wrap{
            position: relative;
            display: flex;
            align-items: center;
            justify-content: center;
            width: 100%;
            height: 100%;
          }
        }
        &.list{
          position: relative;
          width: 100%;
          height: 110px;
          margin-bottom: 10px;
          background-color: #eee;
          border-radius: 5px;
          .image-wrap{
            position: relative;
            display: flex;
            align-items: center;
            justify-content: center;
            top: 10px;
            left: 10px;
            width: 80px;
            height: 80%;
          }
          .stamp-detail{
            position: absolute;
            width: calc(100% - 90px);
            left: 100px;
            top: 30px;
            text-align: left;
            font-size: 12px;
            .name{
              margin-bottom: 5px;
              height: 18px;
              width: calc(100% - 18px);
              overflow: hidden;
            }
            .level, .type{
              margin-top: -5px;
              transform-origin: left;
              transform: scale(0.8);
              height: 18px;
              overflow: hidden;
              text-overflow: ellipsis;
            }
            .price{
              position: relative;
              padding-left: 22px;
              margin-top: 4px;
              &:before{
                content: '';
                position: absolute;
                width: 18px;
                height: 18px;
                left: 0;
                top: -2px;
                border-radius: 50%;
                background-color: #aaa;
                background-image: url(/static/ui/ether.png);
                background-size: contain;
              }
            }
          }
          .seller-detail{
            position: absolute;
            width: 100px;
            top: 8px;
            right: 10px;
            .seller-head{
              float: right;
              width: 20px;
              height: 20px;
              background-color: #aaa;
              border-radius: 50%;
              margin-left: 5px;
            }
            .seller-name{
              padding-top: 1px;
              text-align: right;
              font-size: 12px;
            }
          }
          .btn{
            position: absolute;
            width: 50px;
            height: 20px;
            bottom: 8px;
            right: 10px;
            background-color: @themeColor;
            color: #fff;
            font-size: 12px;
            display: flex;
            align-items: center;
            justify-content: center;
            border-radius: 10px;
            &.buy:active{
              background-color: darken(@themeColor, 6%);
            }
            &.sold{
              background-color: #ccc;
            }
          }
        }
        &.album{
          position: relative;
          width: 100%;
          height: @albumRowHeightTrade;
          margin-right: 10px;
          margin-bottom: 20px;
          .album-wrap{
            position: absolute;
            width: 80px;
            height: 86%;
            top: 7%;
            left: 0;
          }
        }
        .album-detail{
          position: absolute;
          width: 50%;
          top: 0;
          right: -10%;
          font-size: 12px;
          text-align: left;
          transform-origin: top;
          transform: scale(0.8);
          .name{
            margin-top: 6px;
            margin-bottom: 6px;
          }
        }
      }
    }
  }
  .single-warp{
    position: fixed;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
    background-color: #000000cc;
    z-index: 10;
    .single-center{
      position: relative;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 80%;
      height: 80%;
    }
  }
  .multiple-warp{
    position: fixed;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
    background-color: #000000cc;
    z-index: 10;
    .multiple-center{
      position: relative;
      width: 80%;
      height: 80%;
      background-color: #fff;
      border-radius: 8px;
      padding: 16px;
      -webkit-overflow-scrolling: touch;
      overflow: scroll;
      .multi-item{
        position: relative;
        width: 100%;
        height: 140px;
        margin-bottom: 20px;
        .image-wrap{
          position: absolute;
          display: flex;
          align-items: center;
          justify-content: center;
          width: 100px;
          height: 100%;
        }
        .stamp-detail{
          position: absolute;
          left: 120px;
          top: 30px;
          text-align: left;
          .name{
            font-size: 15px;
            font-weight: bold;
            margin-bottom: 10px;
          }
          .level, .type{
            font-size: 12px;
          }
        }
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
