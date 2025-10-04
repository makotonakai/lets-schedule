import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import DashBoard from '../views/Dashboard.vue'
import NewMeeting from '../views/NewMeeting.vue'
import SendEmail from '../views/SendEmail.vue'
import ResetPassword from '../views/ResetPassword.vue'
import HostConfirmedDashBoard from '../views/host/HostConfirmedDashboard.vue'
import HostNotConfirmedDashBoard from '../views/host/HostNotConfirmedDashboard.vue'
import HostNotRespondedDashBoard from '../views/host/HostNotRespondedDashboard.vue'
import GuestConfirmedDashBoard from '../views/guest/GuestConfirmedDashboard.vue'
import GuestNotConfirmedDashBoard from '../views/guest/GuestNotConfirmedDashboard.vue'
import GuestNotRespondedDashBoard from '../views/guest/GuestNotRespondedDashboard.vue'
import SetDateTime from '../views/SetDateTime.vue'
import NewCandidateTime from '../views/NewCandidateTime.vue'
import EditMeeting from '../views/EditMeeting.vue'
import EditCandidateTime from '../views/EditCandidateTime.vue'


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
      path: '/send-email',
      name: 'send-email',
      component: SendEmail
    },
    {
      path: '/:id/reset-password',
      name: 'reset-password',
      component: ResetPassword
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
      path: '/meeting/new',
      name: 'new-meeting',
      component: NewMeeting,
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
      path: '/meeting/host/not-responded/dashboard',
      name: 'host-not-responded-dashboard',
      component: HostNotRespondedDashBoard,
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
      path: '/meeting/guest/not-responded/dashboard',
      name: 'guest-not-responded-dashboard',
      component: GuestNotRespondedDashBoard,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/host/:id/set-date-time',
      name: 'set-date-time',
      component: SetDateTime,
      props: true,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/:id/new',
      name: 'new-candidate-time',
      component: NewCandidateTime,
      props: true,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/:id/host/edit',
      name: 'edit-meeting',
      component: EditMeeting,
      props: true,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/meeting/:id/guest/edit',
      name: 'edit-candidate-time',
      component: EditCandidateTime,
      props: true,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

export default router
