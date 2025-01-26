<template>
  <v-container class="base-width">
    <MainCarouselComponent> </MainCarouselComponent>
  </v-container>
  <v-container class="base-width">
    <div class="mt-5 mx-2 d-flex flex-row align-center">
      <h3 class="text-left">Новости</h3>
      <v-divider class="mt-1 ml-3"></v-divider>
    </div>
    <NewsListComponent></NewsListComponent>
  </v-container>
</template>

<script setup lang="ts">
import NewsListComponent from '@/components/NewsListComponent.vue'
import MainCarouselComponent from '@/components/MainCarouselComponent.vue'
</script>
