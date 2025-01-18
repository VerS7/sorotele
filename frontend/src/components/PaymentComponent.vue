<style scoped>
#captcha__container {
  background-color: #e4e4e4;
  border-radius: 5px;
}

#error__label {
  margin-top: 0.7rem;
  color: red;
}
</style>

<template>
  <v-container max-width="800">
    <v-card elevation="3" class="pa-3">
      <h4 class="text-center text-h4 mt-5">Пополнение счета онлайн</h4>
      <v-form ref="form" v-model="isFormValid" lazy-validation>
        <v-text-field
          class="ma-3"
          variant="underlined"
          label="Номер вашего договора"
          placeholder="Например: sr001"
          v-model="account"
          required
          :rules="[rules.required]"
        ></v-text-field>

        <v-text-field
          class="ma-3"
          variant="underlined"
          placeholder="Например: example@mail.ru"
          v-model="email"
          label="E-mail"
          required
          :rules="[rules.required, rules.email]"
        ></v-text-field>

        <v-text-field
          class="ma-3"
          type="number"
          variant="underlined"
          v-model="paymentSum"
          min="1"
          max="3000"
          label="Сумма к оплате"
          required
          :rules="[
            rules.required,
            rules.onlyInteger,
            (value) => value >= 1 || 'Минимум 1',
            (value) => value <= 3000 || 'Максимум 3000',
          ]"
        ></v-text-field>

        <v-row class="ma-3 mt-5">
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
          <div class="d-flex align-end flex-column">
            <v-btn
              elevation="2"
              class="font-weight-bold"
              @click="submit"
              :color="requestError ? 'red' : 'gray'"
              :disabled="!isFormValid || !isRecaptchaValid"
              >Оплатить</v-btn
            >
            <div v-if="requestError" id="error__label">
              Не удалось подтвердить оплату. Попробуйте позже.
            </div>
          </div>
        </v-row>
      </v-form>
    </v-card>
  </v-container>
  <template>
    <v-dialog class="mb-15" max-width="400" v-model="isDialogActive">
      <v-card class="pa-3">
        <h3 class="text-center">{{ paymentResponse?.service }}</h3>
        <a
          v-bind:href="paymentResponse?.link"
          target="_blank"
          class="link__reverse color-green mt-5 mb-5 text-center"
          >Перейти на страницу оплаты</a
        >
        <v-row class="pa-3">
          <v-spacer></v-spacer>
          <v-btn text="Закрыть" @click="isDialogActive = false"></v-btn>
        </v-row>
      </v-card>
    </v-dialog>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { VForm } from 'vuetify/components'
import { VueRecaptcha } from 'vue-recaptcha'

import { requestPayment } from '@/payment'
import type { PaymentResponse } from '@/payment'

const recaptchaKey = import.meta.env.RECAPTCHA_KEY

const account = ref('')
const email = ref('')
const paymentResponse = ref<PaymentResponse | null>(null)
const paymentSum = ref(0)

const isFormValid = ref(false)
const isRecaptchaValid = ref(false)
const isDialogActive = ref(false)
const requestError = ref(false)

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

  await requestPayment({
    account: account.value,
    email: email.value,
    sum: paymentSum.value.toString(),
  })
    .then((data) => {
      requestError.value = false
      isDialogActive.value = true
      paymentResponse.value = data
      resetForm()
    })
    .catch(() => {
      requestError.value = true
      resetForm()
    })
}

function resetForm() {
  form.value?.reset()
  recaptcha.value?.reset()
  isFormValid.value = false
  isRecaptchaValid.value = false
}
</script>
