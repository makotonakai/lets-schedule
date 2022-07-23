<script setup>
import { useRouter } from 'vue-router'
const props = defineProps({
  id: Number,
  title: String,
  description: String,
  type: String,
  place: String,
  url: String
});
let meeting_id = props.id;

function GoToDetailPage() {
  let router = useRouter()
  router.push(`/meeting/host/confirmed/${meeting_id}`)
}
</script>
<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ title }}</p>
    </header>
    <div class="card-content">
      <div class="meeting-description">
        Description: {{ description }}<br>
        Type: {{ type }}<br>
        Place: {{ place }}<br>
        URL: {{ url }}<br>
      </div>
    </div>
    <footer class="card-footer">
      <button v-on:click="GoToDetailPage" class="card-footer-item">Show Details</button>
    </footer>
  </div>
</template>
