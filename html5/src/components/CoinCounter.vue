<template>
  <div class="coin-counter">
    <div class="coin-img"></div>
    <div v-for="(item, index) in counters" :key="index" class="count-window" :style="'background-position: 0 '+ item + 'px'"></div>
  </div>
</template>

<script>
export default {
  name: 'CoinCounter',
  data () {
    return {
      counters: [36, 36, 36, 36, 36, 36]
    }
  },
  methods: {
    setNum (num) {
      let numArrayNew = num.toString().split('').map(i => parseInt(i)).reverse()
      let numArray = Array(6).fill(0)
      for (let i in numArrayNew) {
        if (i >= numArray.length) {
          break
        } else {
          numArray[i] = numArrayNew[i]
        }
      }
      numArray = numArray.reverse()
      this.counters = this.counters.map((i, k) => {
        let num = numArray[k]
        let position = ((num + 1) * 36) % 360
        let curPos = i
        while (curPos % 360 !== position) {
          curPos += 36
        }
        return curPos
      })
    }
  },
  props: {
    number: Number
  },
  watch: {
    number (val, oldVal) {
      this.setNum(val)
    }
  },
  mounted () {
    this.setNum(0)
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
  .coin-counter{
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 36px;
    margin-bottom: 20px;
    .coin-img{
      width: 36px;
      height: 100%;
      background-image: url(/static/ui/coin.w-shadow.png);
      background-size: cover;
      margin-right: 5px;
    }
    .count-window{
      width: 22px;
      height: 100%;
      margin: 0 3px 0 3px;
      background-color: #fff;
      border: 1px solid #ddd;
      border-radius: 3px;
      box-shadow:1px 2px 3px 1px rgba(0, 0, 0, 0.11) inset;
      background-image: url(/static/ui/counter.number.jpg);
      background-size: 20px;
      background-position: 0 36px;
      transition: all 0.3s ease-in-out;
    }
  }
</style>
