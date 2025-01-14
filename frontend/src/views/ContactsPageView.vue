<style scoped>
#contacts {
  padding-left: 5rem;
  padding-right: 5rem;
}

.td__label {
  text-align: end;
  min-width: 125px;
  vertical-align: top;
  font-weight: bold;
}
</style>

<template>
  <v-container id="contacts" class="base-width">
    <v-row>
      <v-col>
        <h2>ООО «Соро Телеком»</h2>
        <h3 class="mt-3 mb-3">Центральный офис</h3>
        <table>
          <tr>
            <td class="td__label">Адрес:</td>
            <td>141840, Московская обл., Дмитровский р-н, д. Курово, д. 69</td>
          </tr>
          <tr>
            <td class="td__label">Телефон:</td>
            <td>8 (495) 99-111-08</td>
          </tr>
          <tr>
            <td class="td__label">E-Mail:</td>
            <td>
              <a class="link__reverse color-green" href="mailto:i@sorotele.com">i@sorotele.com</a>
            </td>
          </tr>
          <tr>
            <td class="td__label">Время работы:</td>
            <td>ПН-ПТ с 9:00 до 18:00 СБ-ВС Выходной</td>
          </tr>
        </table>
      </v-col>
      <v-col>
        <div ref="map" style="width: 500px; height: 350px"></div>
      </v-col>
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
    'https://api-maps.yandex.ru/services/constructor/1.0/js/?sid=5BLbst97YpbSac8aGVKyWPTjK0kZXvYQ&width=500&height=350'

  map.value?.appendChild(script)
})
</script>
