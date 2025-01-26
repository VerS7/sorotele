<template>
  <v-breadcrumbs class="pa-0 ml-1 mt-3" v-if="isActive" :items="items">
    <template v-slot:item="{ item }">
      <v-breadcrumbs-item
        class="tran-green link"
        :disabled="item.disabled"
        @click="
          () => {
            if (item.href != null) {
              router.push(item.href)
            }
          }
        "
      >
        {{ item.title }}
      </v-breadcrumbs-item>
    </template>
  </v-breadcrumbs>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isActive = ref(false)

const root = router.options.routes.find((route) => route.path === '/')
const items = ref()

watch(router.currentRoute, () => {
  const current = router.currentRoute.value
  const matchedRoutes = current.matched

  items.value = [
    { title: 'Главная', href: '/', disabled: false },
    ...matchedRoutes.map((route) => ({
      title: route.name,
      href: route.path,
      disabled: route.path === current.path,
    })),
  ]

  isActive.value = current.path !== '/'
})
</script>
