<style scoped>
footer {
  background-color: #393939;
  width: 100%;
  color: #ffffff;
  padding: 0px;
}

#footer__line {
  width: 100%;
  background-color: #333333;
  padding: 20px;
}
</style>

<template>
  <v-footer class="flex-column">
    <v-container fluid class="pa-5 d-flex justify-center base-width">
      <div>
        <p class="font-weight-bold">Общество с ограниченной ответственностью «Соро Телеком»</p>
        <p>
          Юридический адрес: 141840, Московская обл., Дмитровский р-н, д. Курово, д. 68, корп. 1
        </p>
        <p>Фактический адрес: 141840, Московская обл., Дмитровский р-н, д. Курово, д. 69</p>
        <p>ОГРН:1135007001381 ИНН: 5007086013</p>
        <div class="mt-10">
          <a class="link" href="/">Главная</a>
          -
          <a class="link" href="/tarifs">Тарифы</a>
          -
          <a class="link" href="/customers">Абонентам</a>
          -
          <a class="link" href="/order">Заявка на подключение</a>
          -
          <a class="link" href="/contacts">Контакты</a>
        </div>
      </div>
      <div class="d-flex ml-15">
        <v-img :width="110" class="ml-5" src="/lic1.jpg"> </v-img>
        <v-img :width="110" class="ml-5" src="/lic2.jpg"> </v-img>
      </div>
    </v-container>
    <div id="footer__line" class="text-center">
      <p>{{ new Date().getFullYear() }} &copy; ООО «СОРО ТЕЛЕКОМ»</p>
    </div>
  </v-footer>
</template>
