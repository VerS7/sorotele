// Vue
import { createApp } from 'vue'
import App from './App.vue'

// Router
import { createRouter, createWebHistory } from 'vue-router'

// Vuetify
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import '@mdi/font/css/materialdesignicons.css'
import './assets/main.css'

// Views
import MainPageView from './views/MainPageView.vue'
import ServicePageView from './views/ServicePageView.vue'
import ContactsPageView from './views/ContactsPageView.vue'

const app = createApp(App)
const vuetify = createVuetify({
  components,
  directives,
})

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: MainPageView },
    { path: '/service', component: ServicePageView },
    { path: '/tarifs', component: MainPageView },
    { path: '/customers', component: MainPageView },
    { path: '/contacts', component: ContactsPageView },
    { path: '/order', component: MainPageView },
    { path: '/lk', component: MainPageView },
  ],
})

app.use(router).use(vuetify).use(router).mount('#app')
