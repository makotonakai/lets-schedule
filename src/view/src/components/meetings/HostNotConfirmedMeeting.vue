<script>
export default {
  props: {
    id: Number,
    title: String,
    description: String,
    type: String,
    place: String,
    url: String
  },
  methods: {
    getURL() {
      return "/meeting/host/" + this.id + "/set-date-time"
    }
  }
};
</script>
<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">{{ title }}</p>
    </header>
    <div class="card-content">
      <div class="meeting-description">
        概要: {{ description }} <br>
        形式: {{ type }} <br>
        開催場所: {{ place }} <br>
        URL: {{ url }} <br>
      </div>
    </div>
   <div class="card-footer-item">
      <a :href="getURL()">詳細</a>
    </div>
  </div>
</template>
