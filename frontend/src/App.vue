<style scoped>
main {
  min-height: 50vh;
}

#app {
  display: flex;
  flex-direction: column;
}

.scroll-btn {
  transition: color 1s ease-in;
  position: fixed;
  bottom: 50px;
  left: 35px;
  z-index: 1000;
}

.fade-enter-active {
  transition: all 0.5s ease;
}

.fade-leave-active {
  transition: all 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

<template>
  <HeaderComponent></HeaderComponent>
  <main>
    <RouterView />
  </main>
  <FooterComponent></FooterComponent>
  <transition name="fade">
    <v-fab
      class="scroll-btn rounded-circle"
      icon="mdi-chevron-up"
      color="green"
      v-if="isScrollBtnVisible"
      @click="scrollToTop"
    >
    </v-fab>
  </transition>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterView } from 'vue-router'
import FooterComponent from '@/components/FooterComponent.vue'
import HeaderComponent from '@/components/HeaderComponent.vue'

const isScrollBtnVisible = ref(false)

onMounted(() => {
  window.addEventListener('scroll', () => (isScrollBtnVisible.value = window.scrollY > 100))
})

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}
</script>
