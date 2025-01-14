<style scoped>
.item {
  flex: 0 1 calc(33.333% - 16px);
  margin-bottom: 16px;
}

.news__header {
  font-size: 1rem;
}

.news__content {
  text-align: justify;
}
</style>

<template>
  <v-container class="item">
    <h2 class="news__header color-green mb-5">{{ props.header }}</h2>
    <div class="news__content">
      <slot></slot>
    </div>
  </v-container>
</template>

<script setup lang="ts">
const props = defineProps({ header: { type: String, required: true } })
</script>
