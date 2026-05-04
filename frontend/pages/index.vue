<script setup lang="ts">
const { currentUser, isAdmin } = useTrainingTrackerDemo()

await navigateTo(currentUser.value ? (isAdmin.value ? '/admin' : '/app') : '/login')
</script>

<template>
  <div />
</template>
