<style scoped>
.btn {
  margin-left: 1px;
  margin-top: 1px;
}
</style>

<template>
  <v-btn
    variant="flat"
    height="50"
    rounded="0"
    class="btn"
    :ripple="false"
    :style="{ backgroundColor: buttonColor() }"
    @click="
      () => {
        router.push(props.linkTo)
      }
    "
  >
    <template v-if="props.icon" v-slot:prepend>
      <v-icon size="30" color="white">{{ props.icon }}</v-icon>
    </template>
    <span class="link color-white">{{ props.title }}</span>
  </v-btn>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const props = defineProps<ButtonProps>()

interface ButtonProps {
  readonly title: string
  readonly linkTo: string
  readonly linkRefered?: string
  readonly icon?: string
}

function buttonColor() {
  return [props.linkTo, props.linkRefered].includes(router.currentRoute.value.path)
    ? '#56831a'
    : '#6fa712'
}
</script>
