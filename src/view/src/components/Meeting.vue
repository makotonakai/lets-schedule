<script setup>
const props = defineProps({
  title: String,
  description: String,
  type: String,
  place: String,
  url: String
});
</script>
<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ title }}</p>
    </header>
    <div class="card-content">
      <div class="meeting-description">
        Description: {{ description }} <br>
        Type: {{ type }} <br>
        Place: {{ place }} <br>
        URL: {{ url }} <br>
      </div>
    </div>
  </div>
</template>
