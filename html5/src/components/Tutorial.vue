<template>
  <div v-if="show" class="tutorial" :class="finish?'hide':null">
    <img class="tutorial-page" :src="currentPageImg" alt="" @click="pageClick">
  </div>
</template>

<script>
export default {
  name: 'Tutorial',
  data () {
    return {
      show: false,
      currentPage: 0,
      tutorPages: [
        '/static/ui/slide.tutorial1.jpg',
        '/static/ui/slide.tutorial2.jpg',
        '/static/ui/slide.tutorial3.jpg',
        '/static/ui/slide.tutorial4.jpg',
        '/static/ui/slide.tutorial5.jpg',
        '/static/ui/slide.tutorial6.jpg'
      ]
    }
  },
  computed: {
    currentPageImg () {
      if (this.currentPage < 6) {
        return this.tutorPages[this.currentPage]
      } else {
        return ''
      }
    },
    finish () {
      return this.currentPage > 5
    }
  },
  methods: {
    pageClick () {
      this.currentPage++
      localStorage.setItem('tutorShowed', 'showed')
    }
  },
  mounted () {
    if (!localStorage.getItem('tutorShowed')) {
      this.show = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.tutorial{
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 9999;
  background-color: #363636;
  transition: all 0.5s ease-in-out;
  opacity: 1;
  &.hide{
    opacity: 0;
    pointer-events: none;
  }
  .tutorial-page{
    position: absolute;
    width: 100%;
    bottom: 0;
    left: 0;
  }
}
</style>
