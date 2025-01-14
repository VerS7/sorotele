<style scoped>
#app {
  display: flex;
  flex-direction: column;
}

main {
  min-height: 100vh;
}
</style>

<template>
  <HeaderComponent></HeaderComponent>

  <main class="mb-10">
    <RouterView />
    <v-container class="base-width">
      <v-divider class="mt-10" thickness="3"></v-divider>
      <h4 class="text-h4 text-center mt-5">Новости</h4>
      <NewsListComponent></NewsListComponent>
    </v-container>
  </main>

  <FooterComponent></FooterComponent>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import FooterComponent from './components/FooterComponent.vue'
import HeaderComponent from './components/HeaderComponent.vue'
import NewsListComponent from './components/NewsListComponent.vue'
</script>
