<style scoped>
table {
  max-width: 700px;
  text-align: center;
}

td {
  padding: 1rem;
}

.bgr-green {
  background-color: #6ba111;
  color: #ffffff;
}

.bgr-gray {
  background-color: #e8e8e8;
}

.tarifs {
  width: 55%;
}

.tarif__price {
  font-size: 1.6rem;
  font-weight: bold;
}

.tarif__title {
  font-size: 2rem;
  font-weight: 700;
}

.tarif__additional {
  font-weight: 300;
}
</style>

<template>
  <v-container class="tarifs pt-0 pr-0 pl-0">
    <v-container class="d-flex flex-column">
      <v-container class="d-flex justify-center align-center">
        <v-hover v-slot="{ isHovering, props }">
          <v-card
            class="rounded-lg pa-5"
            max-width="500"
            v-bind="props"
            :elevation="isHovering ? 24 : 7"
          >
            <v-card-title class="tarif__title">Единый безлимитный</v-card-title>
            <v-card-subtitle>Подключение по радиоканалу</v-card-subtitle>
            <v-card-text>
              <p><b class="tarif__price mr-1">1100</b> руб./мес.</p>
              <v-divider class="mt-4 mb-4"></v-divider>
              <v-row class="mt-4">
                <v-icon class="ml-2 mr-3">mdi-web</v-icon>
                <p>Скорость интернета:</p>
              </v-row>
              <p class="ml-7 mt-4"><b>до 40 Мбит/с</b></p>
              <p class="mt-5">Стоимость подключения: <b>от 3 900 руб.</b></p>
              <p class="mt-2">Статический IP адрес: <b>150 руб.</b></p>
              <p class="mt-5 tarif__additional">
                Максимально возможная скорость подключения и цена подключения определяется
                техническим обследованием и измерением уровня сигнала на месте подключения.
              </p>
            </v-card-text>
            <v-card-actions>
              <v-btn
                class="pa-6"
                variant="flat"
                color="green"
                rounded="xl"
                @click="model = true"
                block
                >Подключить</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-hover>
      </v-container>
    </v-container>
  </v-container>
  <template>
    <OrderComponent v-model="model"></OrderComponent>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import OrderComponent from '@/components/OrderComponent.vue'

const model = ref(false)
</script>
