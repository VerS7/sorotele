<style scoped>
#app {
  display: flex;
  flex-direction: column;
}

main {
  min-height: 70vh;
  background-color: rgb(250, 250, 250);
}
</style>

<template>
  <HeaderComponent></HeaderComponent>
  <v-divider thickness="3"></v-divider>
  <main>
    <RouterView />
  </main>
  <v-divider thickness="3"></v-divider>
  <FooterComponent></FooterComponent>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import FooterComponent from './components/FooterComponent.vue'
import HeaderComponent from './components/HeaderComponent.vue'
</script>
