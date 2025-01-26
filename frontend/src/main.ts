// Vue
import { createApp } from 'vue'

// Router
import { createRouter, createWebHistory } from 'vue-router'

// Vuetify
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import 'vuetify/styles'

// Styles
import '@mdi/font/css/materialdesignicons.css'
import '@/assets/main.css'

// Views
import App from './App.vue'
import MainPageView from '@/views/MainPageView.vue'
import ServicePageView from '@/views/ServicePageView.vue'
import ContactsPageView from '@/views/ContactsPageView.vue'
import TarifsPageView from '@/views/TarifsPageView.vue'
import CustomersPageView from '@/views/CustomersPageView.vue'
import OrderPageView from '@/views/OrderPageView.vue'
import LoginPageView from '@/views/LoginPageView.vue'
import LkPageView from '@/views/LkPageView.vue'
import AdminPanelView from '@/views/AdminPanelView.vue'

const app = createApp(App)
const vuetify = createVuetify({
  components,
  directives,
})
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'Главная', component: MainPageView },
    { path: '/service', name: 'Услуги', component: ServicePageView },
    { path: '/tarifs', name: 'Тарифы', component: TarifsPageView },
    { path: '/customers', name: 'Абонентам', component: CustomersPageView },
    { path: '/contacts', name: 'Контакты', component: ContactsPageView },
    { path: '/order', name: 'Заявка на подключение', component: OrderPageView },
    { path: '/login', name: 'Вход', component: LoginPageView },
    { path: '/lk', name: 'Личный кабинет', component: LkPageView },
    { path: '/admin', name: 'Администрирование', component: AdminPanelView },
  ],
})

app.use(router).use(vuetify).mount('#app')
