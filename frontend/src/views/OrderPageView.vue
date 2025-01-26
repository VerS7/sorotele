<style scoped>
p {
  font-size: 1.2rem;
}
</style>

<template>
  <v-container class="base-width pl-15 pr-15">
    <div class="ml-15 mr-15 text-justify mt-5">
      <p class="mb-10">
        Наша организация имеет
        <b>обширную волоконно-оптическую сеть по всей Московской области</b>, и слаженную комманду с
        наработанным опытом создания сетевой инфраструктуры. Среди наших объектов есть как и
        удаленные от городов производства, курорты и коттеджные поселки так и многоквартирные дома.
        Список населенных пунктов и стоимость подключения.
      </p>
      <p class="mb-10">
        Даже если вы не проживаете в перечисленных выше населенных пунктах, но хотите подключить
        гарантированно высокоскоростной и безлимитный интернет,
        <b>отправьте нам ваши контакты через заявку.</b>
      </p>
      <v-row class="mt-10" justify="center">
        <v-hover v-slot="{ isHovering, props }"
          ><v-btn
            type="flat"
            color="green"
            rounded="xl"
            height="50"
            width="250"
            v-bind="props"
            :elevation="isHovering ? 10 : 0"
            @click="model = true"
            >Оставить заявку</v-btn
          ></v-hover
        >
      </v-row>
    </div>
  </v-container>
  <template>
    <OrderComponent v-model="model"></OrderComponent>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import OrderComponent from '@/components/OrderComponent.vue'
import { tr } from 'vuetify/locale'

const model = ref(false)
</script>
