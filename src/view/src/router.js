import Vue from 'vue'
import Router from 'vue-router'
import Home from './component/Home'
import Login from './component/Login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/login',
      component: Login
    }
  ]
})