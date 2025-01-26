<template>
  <a class="link tran-green" @click="router.push(props.linkTo)">
    <slot></slot>
  </a>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const props = defineProps({ linkTo: { type: String, required: true } })
</script>
