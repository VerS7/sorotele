<template>
  <v-container class="base-width">
    <v-card class="pa-3" min-height="500">
      <h2 class="ml-1">Админ-панель</h2>
      <v-container class="pa-0 mt-3" fluid>
        <v-btn class="mr-5" elevation="1" @click="createRateDialog = true"
          >Создать новый тариф</v-btn
        >
        <v-btn
          elevation="1"
          @click="
            () => {
              createUserDialog = true
              getRates()
            }
          "
          >Создать нового пользователя</v-btn
        >
      </v-container>
    </v-card>
  </v-container>
  <v-dialog v-model="createRateDialog" max-width="500">
    <v-card class="pa-5" rounded="lg">
      <h4>Новый тариф</h4>
      <v-form ref="newRateForm" v-model="isNewRateFormValid" lazy-validation>
        <v-text-field
          variant="underlined"
          v-model="newRateName"
          :rules="[rules.required]"
          label="Название"
          required
        ></v-text-field>
        <v-text-field
          variant="underlined"
          type="number"
          v-model="newRatePrice"
          :rules="[rules.required, rules.onlyInteger]"
          label="Цена"
          required
        ></v-text-field>
      </v-form>
      <v-btn
        class="font-weight-bold"
        variant="flat"
        rounded="xl"
        @click="submitCreateNewRate"
        :disabled="!isNewRateFormValid"
      >
        Создать
      </v-btn>
    </v-card>
  </v-dialog>
  <v-dialog v-model="createUserDialog" max-width="500">
    <v-card class="pa-5" rounded="lg">
      <h4>Новый пользователь</h4>
      <v-form ref="newUserForm" v-model="isNewUserFormValid" lazy-validation>
        <v-text-field
          variant="underlined"
          v-model="newUserName"
          :rules="[rules.required]"
          label="Имя"
          required
        ></v-text-field>
        <v-text-field
          variant="underlined"
          v-model="newUserSurname"
          :rules="[rules.required]"
          label="Фамилия"
          required
        ></v-text-field>
        <v-text-field
          label="Пароль"
          placeholder="Введите пароль"
          variant="underlined"
          v-model="newUserPassword"
          required
          @click:append-inner="isPasswordVisible = !isPasswordVisible"
          :append-inner-icon="isPasswordVisible ? 'mdi-eye' : 'mdi-eye-closed'"
          :type="isPasswordVisible ? 'text' : 'password'"
          :rules="[rules.required]"
        ></v-text-field>
        <v-select
          variant="underlined"
          v-model="newUserRate"
          label="Тариф"
          no-data-text="Отсутствуют тарифы"
          :items="rates?.map((v) => v.name)"
          required
        ></v-select>
      </v-form>
      <v-btn
        class="font-weight-bold"
        variant="flat"
        rounded="xl"
        @click="submitCreateNewUser()"
        :disabled="!isNewUserFormValid"
      >
        Создать
      </v-btn>
    </v-card>
  </v-dialog>
  <v-dialog v-model="createdNewUserDialog" max-width="300">
    <v-card class="pa-5" rounded="lg">
      <h4>Создан пользователь</h4>
      <div class="font-weight-medium mt-1">
        <v-row class="pa-3 pb-2">
          <p class="mr-1">{{ createdNewUserData?.name }}</p>
          <p>{{ createdNewUserData?.surname }}</p>
        </v-row>
        <p>{{ createdNewUserData?.account }}</p>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { type UserData, type NewUserResponse, CreateNewUser } from '@/account'
import type { Rate } from '@/rate'

import { loadAuthToken, removeAuthToken, loadSessionUserData } from '@/account'
import { createNewRate, getAllRates } from '@/rate'
import type { VForm } from 'vuetify/components'

const userData = ref<UserData | null>(null)
const createdNewUserData = ref<NewUserResponse | null>(null)
const token = ref<string | null>(null)
const router = useRouter()

const createUserDialog = ref(false)
const createRateDialog = ref(false)
const createdNewUserDialog = ref(false)

const isNewRateFormValid = ref(false)
const newRateForm = ref<VForm | null>()
const newRateName = ref('')
const newRatePrice = ref('')

const isNewUserFormValid = ref(false)
const newUserForm = ref<VForm | null>()
const newUserName = ref('')
const newUserSurname = ref('')
const newUserPassword = ref('')
const newUserRate = ref()
const isPasswordVisible = ref(false)

const rates = ref<Rate[] | null>(null)

const rules = {
  required: (value: string) => !!value || 'Это поле обязательно',
  onlyInteger: (value: string) => Number.isInteger(Number(value)) || 'Введите целое число',
}

try {
  userData.value = loadSessionUserData()
  token.value = loadAuthToken()
  getRates()
  if (userData.value.role != 'Admin') {
    router.push('/')
  }
} catch {
  removeAuthToken()
  router.push('/')
}

async function getRates() {
  if (token.value == null) {
    return
  }
  rates.value = await getAllRates(token.value)
}

async function submitCreateNewRate() {
  if (token.value == null) {
    return
  }

  await createNewRate(token.value, {
    name: newRateName.value,
    price: parseFloat(newRatePrice.value),
  })
  createRateDialog.value = false
}

async function submitCreateNewUser() {
  if (token.value == null) {
    return
  }

  const rateid = rates.value?.filter((v) => v.name == newUserRate.value)[0].id
  if (rateid == undefined) {
    return
  }

  await CreateNewUser(token.value, {
    name: newUserName.value,
    surname: newUserSurname.value,
    role: 'User',
    rateID: rateid,
  }).then((data) => (createdNewUserData.value = data))
  createUserDialog.value = false
  createdNewUserDialog.value = true
}
</script>
