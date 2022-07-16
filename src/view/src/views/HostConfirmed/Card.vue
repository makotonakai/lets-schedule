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

const router = useRouter()
const meeting_id = props.id;

function GoToDetailPage() {
  router.push(`/meeting/host/confirmed/detail/edit/${meeting_id}`)
}
</script>
<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ title }}</p>
    </header>
    <div class="card-content">
      <div class="meeting-description">
        description: {{ description }}
        <br>
        type: {{ type }}
        <br>
        place: {{ place }}
        <br>
        url: {{ url }}
      </div>
    </div>
    <footer class="card-footer">
      <button v-on:click="GoToDetailPage()" class="card-footer-item">Show Details</button>
    </footer>
  </div>
</template>
