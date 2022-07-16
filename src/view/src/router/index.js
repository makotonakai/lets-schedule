import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import DashBoard from '../views/DashBoard.vue'
import NewMeeting from '../views/NewMeeting.vue'
import HostConfirmedDashBoard from '../views/HostConfirmed/DashBoard.vue'
import HostNotYetConfirmedDashBoard from '../views/HostNotYetConfirmed/DashBoard.vue'
import GuestConfirmedDashBoard from '../views/GuestConfirmed/DashBoard.vue'
import HostConfirmedDetailEdit from '../views/HostConfirmed/Detail.vue'
import GuestRespondedDashBoard from '../views/GuestResponded/DashBoard.vue'
import GuestNotYetRespondedDashBoard from '../views/GuestNotYetResponded/DashBoard.vue'
import GuestDetailEdit from '../views/GuestResponded/Detail.vue'
import GuestDetailNew from '../views/GuestNotYetResponded/Detail.vue'

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
      name: 'host-confirmed-dashboard',
      component: HostConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/host/not-yet-confirmed',
      name: 'host-not-yet-confirmed-dashboard',
      component: HostNotYetConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/confirmed',
      name: 'guest-confirmed-dashboard',
      component: GuestConfirmedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/responded',
      name: 'guest-responded-dashboard',
      component: GuestRespondedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/guest/not-yet-responded',
      name: 'guest-not-yet-responded-dashboard',
      component: GuestNotYetRespondedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/meeting/host/confirmed/detail/edit/:id",
      name: "host-edit-detail",
      component: HostConfirmedDetailEdit,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/meeting/guest/has-not-responded/new/:id",
      name: "guest-enter-candidate-time",
      component: GuestDetailNew,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/meeting/guest/responded/edit/:id",
      name: "guest-edit-candidate-time",
      component: GuestDetailEdit,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

export default router
