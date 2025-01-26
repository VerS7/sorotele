<style scoped>
#carousel {
  background-color: #f7f7f7;
}

.carousel__inner {
  padding: 0 5rem 0 5rem;
}

.banner__info {
  font-size: 1.2rem;
}

.banner__additional {
  font-size: 1.5rem;
}

h4 {
  font-weight: 500;
  font-size: 1.7rem;
}
</style>

<template>
  <v-carousel
    class="rounded-lg"
    id="carousel"
    interval="10000"
    show-arrows
    cycle
    hide-delimiter-background
  >
    <template v-slot:prev="{ props }">
      <v-btn size="35" color="green" @click="props.onClick" rounded="circle">
        <v-icon :size="35">mdi-chevron-left</v-icon>
      </v-btn>
    </template>
    <template v-slot:next="{ props }">
      <v-btn size="35" color="green" @click="props.onClick" rounded="circle">
        <v-icon :size="35">mdi-chevron-right</v-icon>
      </v-btn>
    </template>
    <v-carousel-item>
      <v-container class="carousel__inner d-flex align-center justify-center">
        <v-container>
          <v-row class="mt-3" justify="center">
            <v-img class="fill-height" src="/banner1.png" max-width="800" draggable="false" />
          </v-row>
          <v-col class="mt-4">
            <p class="text-center text-h5 font-weight-medium">
              Теперь, при пополнении счета, вам зачисляется бонусная сумма
            </p>
            <v-row class="banner__info mt-5 ml-15 mr-15 text-center" justify="space-between">
              <div class="mb-5">
                <p>На 6 месяцев</p>
                <p class="banner__additional color-green">10%</p>
              </div>
              <div>
                <p>На 12 месяцев</p>
                <p class="banner__additional color-green">20%</p>
              </div>
              <div>
                <p>На 24 месяца</p>
                <p class="banner__additional color-green">30%</p>
              </div>
            </v-row>
          </v-col>
        </v-container>
      </v-container>
    </v-carousel-item>
    <v-carousel-item>
      <v-container
        class="carousel__inner d-flex align-center justify-center flex-column text-center"
      >
        <v-container class="align-center d-flex flex-row pa-10">
          <v-img :width="270" src="/ban2.png" draggable="false" />
          <v-container class="align-center justify-center d-flex flex-column">
            <p class="text-h5 mb-10">БЕСПРОЦЕНТНАЯ РАССРОЧКА ПРИ ПОДКЛЮЧЕНИИ</p>
            <p class="mb-10 color-green">
              Для новых абонентов желающих подключиться к оптоволоконной сети «Соро Телеком»,
              действует специальное предложение
            </p>
            <p class="font-weight-bold">Рассрочка оплаты стоимости подключения на 3 месяца!</p>
          </v-container>
        </v-container>
        <v-container>
          Узнать о технической возможности, а так же о ценах вы всегда можете по телефону
          <span class="font-weight-bold">8 (495) 99-111-08</span> или заполнив форму в разделе
          <span><a class="link__reverse color-green" href="/contacts">контакты</a>. </span>
        </v-container>
      </v-container>
    </v-carousel-item>
  </v-carousel>
</template>
