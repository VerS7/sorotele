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

p,
a {
  font-weight: 500;
}
</style>

<template>
  <v-container class="base-width">
    <v-card class="pa-3 mt-n6" elevation="0">
      <v-row>
        <v-col class="d-flex flex-column mr-0 pr-0" cols="4">
          <v-card class="pa-5 ma-3" rounded="xl" height="100%">
            <v-row>
              <v-icon size="40" id="user__icon" class="rounded-circle"
                >mdi-account-check-outline</v-icon
              >
              <div>
                <v-row class="ml-2 pt-3 pb-3">
                  <p class="mr-1 font-weight-bold">
                    {{ userData?.name ? userData?.name : 'Имя' }}
                  </p>
                  <p class="font-weight-bold">
                    {{ userData?.surname ? userData?.surname : 'Фамилия' }}
                  </p>
                </v-row>
                <template v-if="userData?.role == 'Admin'">
                  <p class="ml-2">
                    <a class="link tran-green" @click="router.push('/admin')">Администратор</a>
                  </p>
                </template>
                <template v-else>
                  <p class="ml-2">Пользователь</p>
                </template>
              </div>
            </v-row>
            <v-col class="mt-6">
              <v-row>
                <p class="font-weight-bold">Лицевой счёт</p>
              </v-row>
              <v-row>
                <v-container class="pa-0" min-height="24">
                  <p>{{ userData?.account }}</p>
                </v-container>
              </v-row>
            </v-col>
            <v-col class="mt-2">
              <v-row>
                <p class="font-weight-bold">Тариф</p>
              </v-row>
              <v-row>
                <p v-if="userData?.rateName != null">
                  {{ userData.rateName }} / {{ userData?.ratePrice }} ₽
                </p>
                <p v-else>Неизвестно</p>
              </v-row>
            </v-col>
            <v-divider class="mt-3 mb-3"></v-divider>
            <v-col>
              <v-row>
                <p class="font-weight-bold">Последнее пополнение</p>
              </v-row>
              <v-row>
                <v-container class="pa-0" min-height="24">
                  <p class="fade-in" v-if="isHistoryLoaded && dynamicData?.history != null">
                    {{
                      dynamicData?.history[0]
                        ? formatDateTime(dynamicData?.history[0].datetime)
                        : 'Неизвестно'
                    }}
                  </p>
                  <p v-if="isHistoryLoaded && dynamicData?.history == null">Неизвестно</p>
                </v-container>
              </v-row>
            </v-col>
            <v-col class="mt-2">
              <v-row>
                <p class="font-weight-bold">Баланс</p>
              </v-row>
              <v-row>
                <v-container class="pa-0" min-height="24">
                  <template v-if="dynamicData?.balance != null">
                    <p class="fade-in" v-if="isHistoryLoaded && dynamicData?.balance != null">
                      {{ dynamicData?.balance ? parseFloat(dynamicData?.balance).toFixed(1) : 0 }}
                      ₽
                    </p>
                  </template>
                </v-container>
                <a class="fade-in link color-green" v-if="isHistoryLoaded" @click="payment()">
                  Пополнить
                </a>
              </v-row>
            </v-col>
          </v-card>
        </v-col>

        <v-col class="d-flex flex-column ml-0 pl-0" cols="8">
          <v-card class="pa-3 ma-3 pt-0" rounded="xl">
            <v-container class="px-0 pb-0">
              <v-col class="pl-0">
                <v-container class="pa-0" min-height="600">
                  <template v-if="dynamicData?.history != null">
                    <template
                      v-for="(item, index) in dynamicData?.history.sort(
                        (a, b) => getUnixTime(b.datetime) - getUnixTime(a.datetime),
                      )"
                      v-bind:key="index"
                    >
                      <v-card
                        class="rounded-lg pa-3 mb-3 fade-in d-flex flex-row"
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
                  <template v-if="dynamicData?.history != null">
                    <v-spacer
                      v-for="i in 10 - dynamicData?.history.length"
                      v-bind:key="i"
                    ></v-spacer>
                  </template>
                </v-container>
              </v-col>
              <v-row class="mr-2 mb-1 ml-n1">
                <v-btn
                  class="ml-2"
                  size="40"
                  icon="mdi-refresh"
                  :disabled="!isRefreshActive"
                  @click="refresh()"
                ></v-btn>
                <v-spacer></v-spacer>
                <v-btn color="red" width="120" rounded="xl" @click="exit()">Выход</v-btn>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
      </v-row>
    </v-card>
  </v-container>

  <template>
    <v-dialog
      class="mb-15"
      max-width="400"
      v-model="isDialogActive"
      transition="dialog-top-transition"
    >
      <v-card class="pa-3" rounded="xl" min-height="150">
        <v-row class="pa-3">
          <v-spacer></v-spacer>
          <v-btn
            variant="flat"
            size="25"
            icon="mdi-window-close"
            @click="
              () => {
                isDialogActive = false
              }
            "
          ></v-btn>
        </v-row>
        <h3 class="text-center">{{ paymentResponse?.service }}</h3>
        <a
          id="payment__link"
          class="link color-green mt-5 mb-3 text-center"
          target="_blank"
          v-bind:href="paymentResponse?.link"
          >Перейти на страницу оплаты</a
        >
      </v-card>
    </v-dialog>
  </template>
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
} from '@/account'
import type { UserData, UserDynamicData } from '@/account'
import type { PaymentResponse } from '@/payment'
import { formatDateTime, getUnixTime } from '@/time'
import { requestPayment } from '@/payment'

const router = useRouter()

const userData = ref<UserData | null>(null)
const dynamicData = ref<UserDynamicData | null>(null)
const token = ref<string | null>(null)

const isHistoryLoaded = ref(false)
const isRefreshActive = ref(true)
const isDialogActive = ref(false)

const paymentResponse = ref<PaymentResponse | null>(null)

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

function exit() {
  removeSessionUserData()
  removeAuthToken()
  router.push('/login')
}

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
  }
  isHistoryLoaded.value = true
  setTimeout(() => {
    isRefreshActive.value = true
  }, 3 * 1000)
}

async function payment() {
  if (userData.value == null) {
    return
  }
  await requestPayment({
    account: userData.value?.account,
    email: 'null',
    sum: userData.value.ratePrice.toString(),
  }).then((data) => {
    isDialogActive.value = true
    paymentResponse.value = data
  })
}
</script>
