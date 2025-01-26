<template>
  <v-container class="base-width d-flex flex-row align-center">
    <v-row class="pa-4">
      <a class="link tran-green mr-6 font-weight-medium" href="mailto:support@sorotele.com"
        >support@sorotele.com</a
      >
      <a class="link tran-green mr-3 font-weight-medium" href="tel:84959911108">8 495 991 11 08</a>
      <v-spacer></v-spacer>
      <NavButtonComponent class="font-weight-medium" link-to="/login">
        <v-icon size="20" class="color-green mr-1">mdi-account</v-icon>
        <span class="tran-green">Личный кабинет</span>
      </NavButtonComponent>
    </v-row>
  </v-container>
  <v-divider></v-divider>
  <v-container class="base-width">
    <v-row class="pa-4 font-weight-medium" align="center">
      <NavButtonComponent link-to="/">
        <v-img
          class="link mr-6"
          draggable="false"
          max-width="150"
          width="150"
          height="40"
          src="/logo.png"
        ></v-img>
      </NavButtonComponent>
      <NavButtonComponent class="mr-7 mt-n1" link-to="/service">Услуги</NavButtonComponent>
      <NavButtonComponent class="mr-7 mt-n1" link-to="/tarifs">Тарифы</NavButtonComponent>
      <NavButtonComponent class="mr-7 mt-n1" link-to="/customers">Абонентам</NavButtonComponent>
      <NavButtonComponent class="mr-7 mt-n1" link-to="/order"
        >Заявка на подключение</NavButtonComponent
      >
      <NavButtonComponent class="mr-7 mt-n1" link-to="/contacts">Контакты</NavButtonComponent>
      <v-spacer></v-spacer>
      <a class="link tran-green mt-n1" @click="isPaymentActive = true"> Пополнение счета онлайн </a>
    </v-row>
    <RouterBreadcrumbsComponent></RouterBreadcrumbsComponent>
  </v-container>
  <template>
    <PaymentComponent v-model="isPaymentActive"></PaymentComponent>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import NavButtonComponent from '@/components/NavButtonComponent.vue'
import RouterBreadcrumbsComponent from '@/components/RouterBreadcrumbsComponent.vue'
import PaymentComponent from '@/components/PaymentComponent.vue'

const isPaymentActive = ref(false)
</script>
