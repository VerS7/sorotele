<template>
  <v-dialog v-model="model" transition="dialog-top-transition">
    <v-container class="d-flex" color="red" @click="model = !model">
      <v-img :src="props.src" :height="props.height" draggable="false"></v-img>
    </v-container>
  </v-dialog>
</template>

<script setup lang="ts">
const model = defineModel<boolean>()

const props = defineProps({
  src: { type: String, required: true },
  height: { type: String || Number, requred: false },
})
</script>
