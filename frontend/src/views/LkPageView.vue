<template>
  <v-container class="base-width">
    <Suspense>
      <AccountComponent></AccountComponent>
    </Suspense>
  </v-container>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

import AccountComponent from '@/components/AccountComponent.vue'

import { loadAuthToken, removeAuthToken } from '@/account'

const router = useRouter()

try {
  loadAuthToken()
} catch {
  removeAuthToken()
  router.push('/login')
}
</script>
