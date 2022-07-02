import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import DashBoard from '../views/DashBoard.vue'
import NewMeeting from '../views/NewMeeting.vue'
import HostConfirmed from '../views/HostConfirmed.vue'
import HostNotYetConfirmed from '../views/HostNotYetConfirmed.vue'
import GuestConfirmed from '../views/GuestConfirmed.vue'
import GuestResponded from '../views/GuestResponded.vue'
import GuestNotYetResponded from '../views/GuestNotYetResponded.vue'

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
      path: '/dashboard',
      name: 'dashboard',
      component: DashBoard,
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
    },
    {
      path: '/meeting/host/confirmed',
      name: 'host-confirmed-meeting',
      component: HostConfirmed,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/host/not-yet-confirmed',
      name: 'host-not-yet-confirmed-meeting',
      component: HostNotYetConfirmed,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/confirmed',
      name: 'guest-confirmed-meeting',
      component: GuestConfirmed,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/responded',
      name: 'guest-responded-meeting',
      component: GuestResponded,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/not-yet-responded',
      name: 'guest-not-yet-responded-meeting',
      component: GuestNotYetResponded,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

export default router
