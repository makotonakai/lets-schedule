import { createApp } from 'vue'
import { VueCookieNext } from 'vue-cookie-next'
import App from './App.vue'
import router from './router'
import 'bulma'

const app = createApp(App)

app.use(VueCookieNext)
app.use(router)

app.mount('#app')
