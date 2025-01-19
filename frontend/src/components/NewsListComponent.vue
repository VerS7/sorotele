<style scoped>
.news__container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: left;
}
</style>

<template>
  <v-container class="news__container">
    <NewsItemComponent header="Дарим в подарок до 30%">
      Уважаемые Абоненты! Теперь при пополнении счета, вы можете получить в подарок до 30% от суммы
      платежа! При пополнении на <b>6 месяцев - 10%</b>, на <b>12 месяцев - 20%</b> и на
      <b>24 месяца - 30%</b>.
    </NewsItemComponent>
    <NewsItemComponent header="Оплата картами VISA/MASTERCARD">
      Уважаемые Абоненты! В связи с присоединением к эквайрингу процессингового центра PAYONLINE,
      теперь вы можете оплачивать услуги картами VISA/MASTERCARD. Сделать это можно перейдя по
      ссылке "Оплата пластиковой картой" в личном кабинете.
    </NewsItemComponent>
    <NewsItemComponent header="Слияние с К3Телеком">
      Уважаемые Абоненты! Информируем Вас о том, что для улучшения качества работы сети, а так же
      сервиса обслуживания абонентов произошло <b>слияние компании СОРО ТЕЛЕКОМ и К3Телеком.</b>
      Неизменно для Вас: «Личный кабинет», где Вы сможете пополнять и отслеживать свой баланс, в
      любое удобное для вас время и своевременная и грамотная техподдержка.
    </NewsItemComponent>
  </v-container>
</template>

<script setup lang="ts">
import NewsItemComponent from '@/components/NewsItemComponent.vue'
</script>
