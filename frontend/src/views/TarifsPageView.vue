<style scoped>
table {
  max-width: 700px;
  text-align: center;
}

td {
  padding: 1rem;
}

.bgr-green {
  background-color: #6ba111;
  color: #ffffff;
}

.bgr-gray {
  background-color: #e8e8e8;
}

.tarifs {
  width: 55%;
}
</style>

<template>
  <v-container class="tarifs pt-0 pr-0 pl-0">
    <v-container class="d-flex flex-column">
      <h2 class="text-h4 mb-5">Тарифы</h2>
      <p class="mb-5">Зона 1. Подключение по радиоканалу</p>
      <v-container class="d-flex justify-center align-center">
        <table>
          <tbody>
            <tr class="bgr-green">
              <td>Тариф</td>
              <td>Скорость подключения*</td>
              <td>Абонентская плата (в месяц)</td>
            </tr>
            <tr class="bgr-gray">
              <td>Единый безлимитный</td>
              <td>до 40 Мбит/с</td>
              <td>1100 руб.</td>
            </tr>
            <tr class="bgr-gray">
              <td colspan="3">Статический IP адрес: 150 руб.</td>
            </tr>
            <tr class="bgr-gray">
              <td colspan="3">Стоимость подключения: от 3 900 руб.</td>
            </tr>
          </tbody>
        </table>
      </v-container>
      <p>
        * - максимально возможная скорость подключения и цена подключения определяется техническим
        обследованием и измерением уровня сигнала на месте подключения.
      </p>
    </v-container>
  </v-container>
</template>
