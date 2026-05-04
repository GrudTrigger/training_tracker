<script setup lang="ts">
defineProps<{
  eyebrow?: string
  title: string
  meta?: string
}>()
</script>

<template>
  <header class="page-header">
    <div>
      <p v-if="eyebrow" class="section-kicker">{{ eyebrow }}</p>
      <h2>{{ title }}</h2>
    </div>
    <div v-if="meta" class="header-meta">
      <span>{{ meta }}</span>
    </div>
  </header>
</template>
