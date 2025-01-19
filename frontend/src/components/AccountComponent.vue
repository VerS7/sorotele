<style scoped>
#user__icon {
  padding: 1.7rem;
  background-color: #e4e4e4;
}
</style>

<template>
  <v-container class="base-width">
    <v-card class="pa-3 rounded-lg">
      <v-card class="rounded-lg d-flex flex-row pa-2" elevation="2">
        <v-icon size="40" id="user__icon" class="rounded-circle">mdi-account-check-outline</v-icon>
        <div>
          <v-row class="ml-2 pt-3 pb-3 font-weight-bold">
            <p class="mr-1">{{ userData?.name ? userData?.name : 'Имя' }}</p>
            <p>{{ userData?.surname ? userData?.name : 'Фамилия' }}</p>
          </v-row>
          <p class="ml-2">{{ userData?.role == 'Admin' ? 'Администратор' : 'Пользователь' }}</p>
        </div>
        <v-spacer></v-spacer>
        <div class="mr-10">
          <p class="font-weight-bold">Последнее пополнение:</p>
          <p class="text-end">{{ dynamicData?.lastPaid ? dynamicData?.lastPaid : 'Неизвестно' }}</p>
        </div>
        <div class="mr-3">
          <p class="font-weight-bold">Баланс:</p>
          <p class="text-end">
            {{ dynamicData?.balance ? parseFloat(dynamicData?.balance).toFixed(1) : 0 }} ₽.
          </p>
        </div>
      </v-card>
      <v-col class="mt-1 mb-1">
        <v-divider></v-divider>
        <h4 class="ma-2 text-center">История</h4>
        <v-divider></v-divider>
      </v-col>
      <v-card
        v-for="item in 5"
        v-bind:key="item"
        class="mr-1 ml-1 rounded-lg pa-3 mb-3"
        elevation="2"
        >hello
      </v-card>
      <v-row class="ma-1 justify-end">
        <v-btn size="40" icon="mdi-refresh" :disabled="!isRefreshActive" @click="refresh()"></v-btn>
      </v-row>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import {
  loadAuthToken,
  removeAuthToken,
  loadSessionUserData,
  setSessionUserData,
  getUserData,
  getUserDynamicData,
  getUserRatesData,
} from '@/account'
import type { UserData, UserDynamicData, RateData } from '@/account'

const router = useRouter()

const userData = ref<UserData | null>(null)
const rateData = ref<RateData[] | null>(null)
const dynamicData = ref<UserDynamicData | null>(null)
const token = ref<string | null>(null)

const isRefreshActive = ref<boolean>(true)

onMounted(async () => {
  ensureToken()

  try {
    userData.value = loadSessionUserData()
  } catch {
    if (token.value != null) {
      try {
        userData.value = await getUserData(token.value)
        setSessionUserData(userData.value)
      } catch {
        removeAuthToken()
        router.push('/login')
      }
    }
  }
})

function ensureToken() {
  try {
    token.value = loadAuthToken()
  } catch {
    removeAuthToken()
    router.push('/login')
  }
}

async function refresh() {
  isRefreshActive.value = false

  if (token.value != null) {
    dynamicData.value = await getUserDynamicData(token.value)
    rateData.value = await getUserRatesData(token.value)
  }

  setTimeout(() => {
    isRefreshActive.value = true
  }, 10 * 1000)
}
</script>
