import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import DashBoard from '../views/Dashboard.vue'
import HostConfirmedDashBoard from '../views/host/confirmed/HostConfirmedDashBoard.vue'
import HostNotConfirmedDashBoard from '../views/host/not-confirmed/HostNotConfirmedDashBoard.vue'
import GuestConfirmedDashBoard from '../views/guest/confirmed/GuestConfirmedDashBoard.vue'
import GuestNotConfirmedDashBoard from '../views/guest/not-confirmed/GuestNotConfirmedDashBoard.vue'
import NewMeeting from '../views/NewMeeting.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/meeting/dashboard',
      name: 'dashboard',
      component: DashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/host/confirmed/dashboard',
      name: 'host-confirmed-dashboard',
      component: HostConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/host/not-confirmed/dashboard',
      name: 'host-not-confirmed-dashboard',
      component: HostNotConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/confirmed/dashboard',
      name: 'guest-confirmed-dashboard',
      component: GuestConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/not-confirmed/dashboard',
      name: 'guest-not-confirmed-dashboard',
      component: GuestNotConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/new',
      name: 'new-meeting',
      component: NewMeeting,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

export default router
