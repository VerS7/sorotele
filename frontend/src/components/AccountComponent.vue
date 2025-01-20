<style scoped>
#user__icon {
  padding: 1.7rem;
  background-color: #e4e4e4;
}

.fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>

<template>
  <v-container class="base-width">
    <v-card class="pa-3 rounded-lg d-flex flex-column justify-end" min-height="580">
      <v-card class="rounded-lg d-flex flex-row pa-2" elevation="2">
        <v-icon size="40" id="user__icon" class="rounded-circle">mdi-account-check-outline</v-icon>
        <div>
          <v-row class="ml-2 pt-3 pb-3 font-weight-bold">
            <p class="mr-1">{{ userData?.name ? userData?.name : 'Имя' }}</p>
            <p>{{ userData?.surname ? userData?.surname : 'Фамилия' }}</p>
          </v-row>
          <p class="ml-2">{{ userData?.role == 'Admin' ? 'Администратор' : 'Пользователь' }}</p>
        </div>
        <v-spacer></v-spacer>
        <v-container class="pa-0" width="300">
          <p class="font-weight-bold text-end">Последнее пополнение</p>
          <p class="text-end fade-in" v-if="isHistoryLoaded && dynamicData?.history != null">
            {{
              dynamicData?.history[0]
                ? formatDateTime(dynamicData?.history[0].datetime)
                : 'Неизвестно'
            }}
          </p>
        </v-container>
        <v-container class="pa-0 mr-3" width="150">
          <p class="font-weight-bold text-end">Баланс</p>
          <template v-if="dynamicData?.balance != null">
            <p class="text-end fade-in" v-if="isHistoryLoaded && dynamicData?.balance != null">
              {{ dynamicData?.balance ? parseFloat(dynamicData?.balance).toFixed(1) : 0 }} ₽
            </p>
          </template>
        </v-container>
      </v-card>
      <v-col class="mt-1 mb-1">
        <v-divider></v-divider>
        <h4 class="ma-2 text-center">История</h4>
        <v-divider></v-divider>
      </v-col>
      <template v-if="dynamicData?.history != null">
        <template
          v-for="(item, index) in dynamicData?.history.sort(
            (a, b) => getUnixTime(b.datetime) - getUnixTime(a.datetime),
          )"
          v-bind:key="index"
        >
          <v-card
            class="mr-1 ml-1 rounded-lg pa-3 mb-3 fade-in d-flex flex-row"
            elevation="2"
            v-if="isHistoryLoaded"
          >
            <v-container
              class="pa-0"
              max-width="120"
              :class="item.amount > 0 ? 'color-green' : 'color-red'"
              ><b>{{ item.amount > 0 ? 'Пополнение' : 'Списание' }}</b></v-container
            >
            <v-divider vertical class="mr-5"></v-divider>
            <p>
              <b>{{ item.amount }} ₽</b>
            </p>
            <v-spacer></v-spacer>
            <v-divider vertical class="mr-5"></v-divider>
            <p>
              <b>{{ formatDateTime(item.datetime) }}</b>
            </p>
          </v-card>
        </template>
      </template>
      <v-row class="ma-1 justify-end align-end">
        <v-btn size="40" icon="mdi-refresh" :disabled="!isRefreshActive" @click="refresh()"></v-btn>
        <v-spacer></v-spacer>
        <v-btn color="red" @click="exit()">Выход</v-btn>
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
  removeSessionUserData,
  getUserData,
  getUserDynamicData,
  getUserRatesData,
} from '@/account'
import type { UserData, UserDynamicData, RateData } from '@/account'
import { formatDateTime, getUnixTime } from '@/time'

const router = useRouter()

const userData = ref<UserData | null>(null)
const rateData = ref<RateData[] | null>(null)
const dynamicData = ref<UserDynamicData | null>(null)
const token = ref<string | null>(null)

const isHistoryLoaded = ref<boolean>(true)
const isRefreshActive = ref<boolean>(true)

function exit() {
  removeSessionUserData()
  removeAuthToken()
  router.push('/login')
}

onMounted(async () => {
  ensureToken()
  refresh()

  try {
    userData.value = loadSessionUserData()
    if (token.value != null) refresh()
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

setInterval(
  async () => {
    dynamicData.value = null
    isHistoryLoaded.value = false
    if (token.value != null) {
      dynamicData.value = await getUserDynamicData(token.value)
    }
    isHistoryLoaded.value = true
  },
  60 * 5 * 1000,
)

function ensureToken() {
  try {
    token.value = loadAuthToken()
  } catch {
    removeAuthToken()
    router.push('/login')
  }
}

async function refresh() {
  dynamicData.value = null
  isRefreshActive.value = false
  isHistoryLoaded.value = false

  if (token.value != null) {
    dynamicData.value = await getUserDynamicData(token.value)
    console.log(dynamicData.value)
  }
  isHistoryLoaded.value = true
  setTimeout(() => {
    isRefreshActive.value = true
  }, 3 * 1000)
}
</script>
