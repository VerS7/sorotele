<style scoped>
.item {
  flex: 0 1 calc(33.333% - 16px);
  margin-bottom: 16px;
}

.news__header {
  font-weight: 600;
  font-size: 1.1rem;
}

.news__content {
  text-align: justify;
  font-size: 0.9rem;
  color: #757575;
}
</style>

<template>
  <v-card class="item pa-3 ma-2" elevation="0">
    <h2 class="news__header color-green mb-5">{{ props.header }}</h2>
    <div class="news__content">
      <slot></slot>
    </div>
  </v-card>
</template>

<script setup lang="ts">
const props = defineProps({ header: { type: String, required: true } })
</script>
