<style scoped>
#captcha__container {
  background-color: #e4e4e4;
  border-radius: 5px;
}

#error__label {
  margin-top: 0.7rem;
  color: red;
}

h4 {
  font-size: 2rem;
  font-weight: 600;
}
</style>

<template>
  <v-card class="pa-2" elevation="5" rounded="xl">
    <v-form ref="form" v-model="isFormValid" lazy-validation>
      <h4 class="text-center mt-3 mb-5">Вход в личный кабинет</h4>
      <v-form ref="form" v-model="isFormValid" lazy-validation>
        <v-text-field
          class="mx-3"
          label="Номер вашего договора"
          placeholder="Например: sr001"
          density="compact"
          variant="outlined"
          v-model="account"
          required
          :rules="[rules.required]"
        ></v-text-field>
        <v-text-field
          class="mx-3 mt-1"
          label="Пароль"
          placeholder="Введите пароль"
          density="compact"
          variant="outlined"
          v-model="password"
          required
          @click:append-inner="isPasswordVisible = !isPasswordVisible"
          :append-inner-icon="isPasswordVisible ? 'mdi-eye' : 'mdi-eye-closed'"
          :type="isPasswordVisible ? 'text' : 'password'"
          :rules="[rules.required]"
        ></v-text-field>
        <v-row class="ma-3 mt-1 mb-0">
          <div id="captcha__container">
            <VueRecaptcha
              ref="recaptcha"
              @verify="isRecaptchaValid = true"
              @expired="resetForm()"
              :sitekey="recaptchaKey"
              :loadRecaptchaScript="true"
            ></VueRecaptcha>
          </div>
          <v-spacer></v-spacer>
        </v-row>
        <v-row class="px-5 pb-5 mt-5" justify="center">
          <v-btn
            class="font-weight-bold"
            width="250"
            height="45"
            rounded="xl"
            variant="flat"
            @click="submit"
            :color="requestError ? 'red' : 'green'"
            :disabled="!isFormValid || !isRecaptchaValid || !isSubmitEnabled"
          >
            Войти
          </v-btn>
        </v-row>
        <v-row class="mb-2 mt-n5" justify="center">
          <div v-if="requestError" id="error__label">Не удалось войти.</div>
        </v-row>
      </v-form>
    </v-form>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { VForm } from 'vuetify/components'
import { VueRecaptcha } from 'vue-recaptcha'

import { requestAccessToken, setAuthToken } from '@/account'

const recaptchaKey = import.meta.env.RECAPTCHA_KEY

const account = ref('')
const password = ref('')

const isFormValid = ref(false)
const isPasswordVisible = ref(false)
const isRecaptchaValid = ref(false)
const isSubmitEnabled = ref(true)
const requestError = ref(false)

const emit = defineEmits(['success'])

const form = ref<VForm | null>(null)

const recaptcha = ref<VueRecaptcha | null>(null)

const rules = {
  required: (value: string) => !!value || 'Это поле обязательно',
  email: (value: string) => /.+@.+\..+/.test(value) || 'Email должен быть действительным',
  onlyInteger: (value: string) => Number.isInteger(Number(value)) || 'Введите целое число',
}

async function submit(): Promise<void> {
  if (!isFormValid.value || !isRecaptchaValid.value) {
    return
  }
  isSubmitEnabled.value = false
  await requestAccessToken({ login: account.value, password: password.value })
    .then((token) => {
      setAuthToken(token)
      emit('success')
      isSubmitEnabled.value = true
    })
    .catch(() => {
      resetForm()
      requestError.value = true
      isSubmitEnabled.value = true
    })
}

function resetForm() {
  form.value?.reset()
  recaptcha.value?.reset()
  isFormValid.value = false
  isRecaptchaValid.value = false
}
</script>
