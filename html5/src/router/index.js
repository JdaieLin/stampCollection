import Vue from 'vue'
import Router from 'vue-router'
import Slide from '@/pages/Slide'
import Album from '@/pages/Album'
import Trade from '@/pages/Trade'
import Dig from '@/pages/Dig'
import User from '@/pages/User'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'slide',
      component: Slide
    },
    {
      path: '/dig',
      name: 'dig',
      component: Dig
    },
    {
      path: '/album',
      name: 'album',
      component: Album
    },
    {
      path: '/trade',
      name: 'trade',
      component: Trade
    },
    {
      path: '/user',
      name: 'user',
      component: User
    }
  ]
})
