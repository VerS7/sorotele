<style scoped>
#captcha__container {
  background-color: #e4e4e4;
  border-radius: 5px;
}

#error__label {
  margin-top: 0.7rem;
  color: red;
}

#payment__link {
  font-weight: 500;
  font-size: 1.2rem;
}

h4 {
  font-size: 1.6rem;
  font-weight: 600;
}
</style>

<template>
  <v-dialog v-model="model" transition="dialog-top-transition">
    <v-container max-width="500">
      <v-card elevation="3" class="pa-3" rounded="xl">
        <v-row class="ma-n2" justify="end">
          <v-btn
            variant="flat"
            size="40"
            icon="mdi-window-close"
            @click="
              () => {
                resetForm()
                requestError = 0
                model = false
              }
            "
          ></v-btn>
        </v-row>
        <h4 class="text-center">Пополнение счета онлайн</h4>
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

          <v-row class="ma-3 mt-3">
            <div id="captcha__container">
              <VueRecaptcha
                ref="recaptcha"
                @verify="isRecaptchaValid = true"
                @expired="resetForm()"
                :sitekey="recaptchaKey"
                :loadRecaptchaScript="true"
              ></VueRecaptcha>
            </div>
          </v-row>
          <v-row class="mt-5" justify="center">
            <v-btn
              height="45"
              width="200"
              class="font-weight-bold"
              variant="flat"
              rounded="xl"
              @click="submit"
              :color="requestError ? 'red' : 'green'"
              :disabled="!isFormValid || !isRecaptchaValid"
              >Оплатить</v-btn
            >
          </v-row>
          <v-row class="pa-5 pt-0" justify="center">
            <div v-if="requestError == 500" id="error__label">
              Не удалось подтвердить оплату. Попробуйте позже.
            </div>
            <div v-if="requestError == 404" id="error__label">
              Пользователь с данным лицевым счетом не найден.
            </div>
          </v-row>
        </v-form>
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
                  model = false
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
  </v-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { VForm } from 'vuetify/components'
import { VueRecaptcha } from 'vue-recaptcha'

import { requestPayment } from '@/payment'
import type { PaymentResponse } from '@/payment'

const model = defineModel<boolean>()

const recaptchaKey = import.meta.env.RECAPTCHA_KEY

const account = ref('')
const email = ref('')
const paymentResponse = ref<PaymentResponse | null>(null)
const paymentSum = ref(0)

const isFormValid = ref(false)
const isRecaptchaValid = ref(false)
const isDialogActive = ref(false)
const requestError = ref(0)

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
      requestError.value = 0
      isDialogActive.value = true
      paymentResponse.value = data
      resetForm()
    })
    .catch((error) => {
      requestError.value = error.code
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
