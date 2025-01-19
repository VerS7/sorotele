<style scoped>
ul {
  padding-left: 2.5rem;
}

.no-wrap {
  word-wrap: normal;
  hyphens: none;
}
</style>

<template>
  <v-container class="base-width">
    <div class="d-flex justify-center mt-5">
      <LoginComponent
        v-if="token == null"
        width="530"
        @success="router.push('/lk')"
      ></LoginComponent>
    </div>
    <LawInfoDialogComponent></LawInfoDialogComponent>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import LawInfoDialogComponent from '@/components/LawInfoDialogComponent.vue'
import LoginComponent from '@/components/LoginComponent.vue'
import { loadAuthToken } from '@/account'

const router = useRouter()
const token = ref<string | null>(null)

try {
  token.value = loadAuthToken()
  router.push('/lk')
} catch {
  router.push('/login')
}
</script>
