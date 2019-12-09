import Vue from 'vue'
import Router from 'vue-router'

import Home from '../sections/Home'
import JobLocation from '../sections/JobLocation'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  linkActiveClass: 'is-active',
  routes: [
    {
      path: '/',
      name: 'CalenderView',
      component: Home,
      props: true,
    },
    {
      path: '/location/:jobLocation',
      name: 'JobLocationView',
      component: JobLocation,
      props: true,
    },
  ],
})

router.afterEach((to) => {
  if (window.ga) {
    window.ga('set', 'page', to.fullPath)
    window.ga('send', 'pageview')
  }
})

export default router
