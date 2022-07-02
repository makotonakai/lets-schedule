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
let router = useRouter()

function EnterAvailableTime() {
  router.push(`/meeting/guest/candidate_time/${meeting_id}`)
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
      <button v-on:click="EnterAvailableTime()" class="card-footer-item">Enter Available Time</button>
    </footer>
  </div>
</template>