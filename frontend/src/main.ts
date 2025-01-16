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
import TarifsPageView from './views/TarifsPageView.vue'
import CustomersPageView from './views/CustomersPageView.vue'
import OrderPageView from './views/OrderPageView.vue'

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
    { path: '/tarifs', component: TarifsPageView },
    { path: '/customers', component: CustomersPageView },
    { path: '/contacts', component: ContactsPageView },
    { path: '/order', component: OrderPageView },
    { path: '/lk', component: MainPageView },
  ],
})

app.use(router).use(vuetify).mount('#app')
