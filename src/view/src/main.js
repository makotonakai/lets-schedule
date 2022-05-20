import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueSidebarMenu from 'vue-sidebar-menu'
import 'bulma'

const app = createApp(App)
app.use(router)
app.use(VueSidebarMenu)
app.mount('#app')
