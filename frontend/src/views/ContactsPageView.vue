<style scoped>
#contacts {
  padding-left: 5rem;
  padding-right: 5rem;
}

#name {
  font-size: 1.5rem;
  font-weight: 700;
}

#subname {
  font-weight: 500;
}

.label {
  font-weight: 700;
}

.text {
  font-weight: 300;
  font-size: 1rem;
}
</style>

<template>
  <v-container id="contacts" class="base-width mb-15">
    <v-row>
      <v-card class="mr-5 pa-7" rounded="lg" elevation="5" max-width="400">
        <h2 class="mt-n2" id="name">ООО «Соро Телеком»</h2>
        <h3 class="mt-n1" id="subname">Центральный офис</h3>
        <v-row class="mt-1">
          <v-icon class="mt-1 ml-2 mr-5 color-green">mdi-city</v-icon>
          <v-col class="pa-0">
            <p class="label mt-1 mb-2">Адрес:</p>
            <p class="text">141840, Московская обл., Дмитровский р-н, д. Курово, д. 69</p>
          </v-col>
        </v-row>
        <v-row class="mt-5">
          <v-icon class="mt-1 ml-2 mr-5 color-green">mdi-phone</v-icon>
          <v-col class="pa-0">
            <p class="label mt-1 mb-2">Телефон:</p>
            <p class="text">
              <a class="link tran-green" href="tel:84959911108">8 (495) 99-111-08</a>
            </p>
          </v-col>
        </v-row>
        <v-row class="mt-5">
          <v-icon class="mt-1 ml-2 mr-5 color-green">mdi-mail</v-icon>
          <v-col class="pa-0">
            <p class="label mt-1 mb-2">E-Mail:</p>
            <p class="text">
              <a class="link tran-green" href="mailto:i@sorotele.com">i@sorotele.com</a>
            </p>
          </v-col>
        </v-row>
        <v-row class="mt-5">
          <v-icon class="mt-1 ml-2 mr-5 color-green">mdi-clock</v-icon>
          <v-col class="pa-0">
            <p class="label mt-1 mb-2">Время работы:</p>
            <p class="text">ПН-ПТ с 9:00 до 18:00 СБ-ВС Выходной</p>
          </v-col>
        </v-row>
      </v-card>
      <v-card rounded="lg">
        <div id="map" ref="map" style="width: 600px; height: 500px"></div>
      </v-card>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'

const map = ref<HTMLElement | null>(null)

onMounted(() => {
  const script = document.createElement('script')
  script.type = 'text/javascript'
  script.src =
    'https://api-maps.yandex.ru/services/constructor/1.0/js/?sid=5BLbst97YpbSac8aGVKyWPTjK0kZXvYQ&width=600&height=500'

  map.value?.appendChild(script)
})
</script>
