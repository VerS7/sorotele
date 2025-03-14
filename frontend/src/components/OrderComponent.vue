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
  font-size: 1.6rem;
  font-weight: 600;
}
</style>

<template>
  <v-dialog v-model="model" transition="dialog-top-transition">
    <v-row justify="center">
      <v-card class="pa-5" width="700" rounded="xl">
        <v-row class="ma-n2" justify="end">
          <v-btn
            variant="flat"
            size="40"
            icon="mdi-window-close"
            @click="
              () => {
                resetForm()
                sendError = false
                model = false
              }
            "
          ></v-btn>
        </v-row>
        <h4 class="text-center">Заявка на подключение</h4>
        <v-form ref="form" v-model="isFormValid" lazy-validation>
          <v-text-field
            variant="underlined"
            v-model="fullName"
            :rules="[rules.required]"
            label="Ваше имя"
            required
          ></v-text-field>
          <v-text-field
            variant="underlined"
            v-model="contacts"
            :rules="[rules.required]"
            label="Контакты для обратной связи"
            required
          ></v-text-field>
          <v-textarea
            variant="underlined"
            v-model="message"
            :rules="[rules.required]"
            label="Адрес подключения и дополнительные комментарии"
            required
          ></v-textarea>
          <v-row class="pl-3 pr-3 mt-1">
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
          <v-row class="mt-10" justify="center">
            <v-btn
              height="45"
              width="250"
              class="font-weight-bold"
              variant="flat"
              rounded="xl"
              @click="submit"
              :color="sendError ? 'red' : 'green'"
              :disabled="!isFormValid || !isRecaptchaValid"
              >Отправить заявку</v-btn
            >
          </v-row>
          <v-row class="pt-1 pb-2" justify="center">
            <div v-if="sendError" id="error__label">
              Не удалось отправить заявку. Попробуйте позже.
            </div>
          </v-row>
        </v-form>
      </v-card>
    </v-row>
  </v-dialog>

  <template>
    <v-dialog class="mb-15" max-width="600" v-model="isDialogActive">
      <v-card>
        <v-card-title class="text-h4 font-weight-bold text-center color-green">
          Спасибо!</v-card-title
        >
        <v-card-text class="text-h6 text-justify">
          <p class="text-center mb-3">
            Уважаемый <b>{{ confirmedFullName }}</b>
          </p>
          <p>
            Спасибо за вашу заявку! Наши сотрудники свяжутся с вами по указанным контактам в
            ближайшее время.
          </p>
        </v-card-text>
        <v-card-actions>
          <v-btn text="Закрыть" @click="isDialogActive = false"></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </template>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { VForm } from 'vuetify/components'
import { VueRecaptcha } from 'vue-recaptcha'

import { sendOrder } from '@/order'

const model = defineModel<boolean>()

const recaptchaKey = import.meta.env.RECAPTCHA_KEY

const fullName = ref('')
const confirmedFullName = ref('')
const contacts = ref('')
const message = ref('')
const isFormValid = ref(false)
const isRecaptchaValid = ref(false)
const isDialogActive = ref(false)
const sendError = ref(false)

const form = ref<VForm | null>(null)
const recaptcha = ref<VueRecaptcha | null>(null)
const rules = {
  required: (value: string) => !!value || 'Это поле обязательно',
}

async function submit(): Promise<void> {
  if (!isFormValid.value || !isRecaptchaValid.value) {
    return
  }

  await sendOrder({
    fullName: fullName.value,
    contacts: contacts.value,
    message: message.value,
  })
    .then(() => {
      confirmedFullName.value = fullName.value
      sendError.value = false
      isDialogActive.value = true
      resetForm()
    })
    .catch(() => {
      sendError.value = true
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
