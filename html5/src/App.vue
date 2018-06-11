<template>
  <div id="app">
    <router-view/>
    <Tabbar/>
    <Tutorial/>
  </div>
</template>

<script>
import Tabbar from './components/Tabbar'
import Tutorial from './components/Tutorial'

export default {
  name: 'App',
  components: {
    Tabbar,
    Tutorial
  }
}
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
  body,html{
    margin: 0;
    width: 100%;
    height: 100%;
    background-color: #eee;
    overflow: hidden;
  }
  *{
    box-sizing: border-box;
    -webkit-tap-highlight-color: transparent;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    user-select:none;
  }
</style>
