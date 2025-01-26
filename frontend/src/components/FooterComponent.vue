<style scoped>
footer {
  background-color: #f4f6f8;
  width: 100%;
  padding: 0px;
}

#footer__line {
  width: 100%;
  background-color: #333333;
  padding: 20px;
}

#lic {
  max-width: 33%;
}

.route {
  font-weight: 500;
}
</style>

<template>
  <v-footer>
    <v-container class="base-width mt-10" fluid>
      <v-row>
        <v-col>
          <v-img class="mb-3" src="/logo.png" max-width="200" width="200"></v-img>
          <a class="link tran-green font-weight-medium ml-1" href="tel:84959911108"
            >8 495 991 11 08</a
          >
          <div>
            <p class="font-weight-bold mt-2">
              Общество с ограниченной ответственностью «Соро Телеком»
            </p>
            <p class="mt-2">
              Юридический адрес: 141840, Московская обл., Дмитровский р-н, д. Курово, д. 68, корп. 1
            </p>
            <p>Фактический адрес: 141840, Московская обл., Дмитровский р-н, д. Курово, д. 69</p>
            <p class="mt-3">ОГРН:1135007001381 ИНН: 5007086013</p>
          </div>
        </v-col>
        <v-col id="lic">
          <v-row class="pa-5" justify="space-between" align="end">
            <v-img
              class="mr-10 link"
              max-width="130"
              src="/lic1.jpg"
              @Click="
                () => {
                  isLicence1dialog = true
                  isLicence2dialog = false
                }
              "
            >
            </v-img>
            <v-img
              class="link"
              max-width="130"
              src="/lic2.jpg"
              @Click="
                () => {
                  isLicence2dialog = true
                  isLicence1dialog = false
                }
              "
            >
            </v-img>
            <v-spacer></v-spacer>
          </v-row>
        </v-col>
      </v-row>
      <v-row class="pl-15 pr-15 mt-5 mb-7" justify="space-between">
        <NavButtonComponent link-to="/"><span class="route">Главная</span></NavButtonComponent>
        -
        <NavButtonComponent link-to="/tarifs"><span class="route">Тарифы</span></NavButtonComponent>
        -
        <NavButtonComponent link-to="/customers"
          ><span class="route">Абонентам</span></NavButtonComponent
        >
        -
        <NavButtonComponent link-to="/order"
          ><span class="route">Заявка на подключение</span></NavButtonComponent
        >
        -
        <NavButtonComponent link-to="/routes"
          ><span class="route">Контакты</span></NavButtonComponent
        >
      </v-row>
      <p class="text-center font-weight-bold">
        {{ new Date().getFullYear() }} &copy; ООО «СОРО ТЕЛЕКОМ»
      </p>
    </v-container>
  </v-footer>
  <template>
    <ImageDialogComponent
      src="/lic1.jpg"
      v-model="isLicence1dialog"
      height="700"
    ></ImageDialogComponent>
    <ImageDialogComponent
      src="/lic2.jpg"
      v-model="isLicence2dialog"
      height="700"
    ></ImageDialogComponent>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import NavButtonComponent from '@/components/NavButtonComponent.vue'
import ImageDialogComponent from '@/components/ImageDialogComponent.vue'

const isLicence1dialog = ref(false)
const isLicence2dialog = ref(false)
</script>
