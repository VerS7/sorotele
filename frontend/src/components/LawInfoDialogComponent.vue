<style scoped>
ul {
  padding-left: 2.5rem;
}
</style>

<template>
  <v-dialog v-model="isDialogShown" max-width="820" persistent>
    <v-card>
      <v-container class="pl-5 pr-5 pt-5 no-wrap">
        <h3 class="pb-2 text-center">Уважаемый абонент!</h3>
        <p class="pb-2">
          Согласно Федеральному Закону «О связи» от 07.07.2003 №126-ФЗ, Постановлению Правительства
          РФ от 31.12.2021 №2607 «Об утверждении Правил оказания телематических услуг связи» у
          оператора связи должны быть достоверные сведения об абоненте:
        </p>
        <ul class="pb-2">
          <li>Фамилия, Имя, Отчество (при наличии);</li>
          <li>Дата рождения;</li>
          <li>Реквизиты документа, удостоверяющего личность (в т.ч. адрес регистрации)</li>
        </ul>
        <p class="pb-2">
          Вы можете подтвердить свои паспортные данные посетив офис компании с оригиналом документа,
          удостоверяющего личность. Или одним из альтернативных способов, предусмотренных Правилами,
          информацию о которых Вы можете получить, связавшись с нашими сотрудниками по телефону
          <b>8(495)99-111-08</b>, электронной почте info@sorotele.com или при помощи мессенджера
          WhatsApp по номеру <b>8(926)224-17-69</b>.
        </p>
        <b>
          <p class="pb-2">
            Закон требует от операторов связи блокировать договоры абонентов, чьи данные по любым
            причинам неактуальны или не подтверждены.
          </p>
          <p>
            Просим Вас подтвердить паспортные данные по вашему договору в срок до 15 ноября 2024 г.
          </p>
        </b>
        <p class="text-body-2">Нормативные документы:</p>
        <ul class="text-caption">
          <li>
            <a href="http://government.ru/docs/all/138747/" target="_blank"
              >Постановление Правительства Российской Федерации №2607 от 31.12.2021 г. </a
            >(пункт 27, подпункт «а»);
          </li>
          <li>
            <a href="http://government.ru/docs/all/138747/" target="_blank"
              >Федеральный закон "О связи" от 07.07.2003 N 126-ФЗ</a
            >
          </li>
        </ul>
      </v-container>
      <v-row class="ma-4 d-flex justify-end">
        <v-btn class="mr-3" @click="router.push('/')">Закрыть</v-btn>
        <v-btn @click="confirmDialog()">Принять</v-btn>
      </v-row>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const isDialogShown = ref<boolean>(!getDialogShownState())

function getDialogShownState() {
  const value = localStorage.getItem('isDialogShown')
  return value == 'true' ? true : false
}

function confirmDialog() {
  isDialogShown.value = false
  localStorage.setItem('isDialogShown', JSON.stringify(true))
}
</script>
