<script setup>
const props = defineProps({
  title: String,
  role: String,
  hasResponded: Boolean,
});
</script>
<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ title }}</p>
    </header>
    <div class="card-content">
      <div class="meeting-role">
        {{ role }}
      </div>
      <div class="meeting-hasResponded">
        {{ hasResponded }}
      </div>
    </div>
  </div>
</template>
<style scoped>
</style>