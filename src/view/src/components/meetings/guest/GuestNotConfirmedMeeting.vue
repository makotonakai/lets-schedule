<script>
export default {
  props: {
    id: Number,
    title: String,
    description: String,
    isOnsite: Boolean,
    isOnline: Boolean,
    place: String,
    url: String
  },
  methods: {
    getURL() {
      return "/meeting/" + this.id + "/guest/edit"
    },
    showType(isOnsite, isOnline) {
      if (isOnsite && isOnline) {
        return "ハイブリッド開催";
      } else if (isOnsite) {
        return "オンサイト開催";
      } else if (isOnline) {
        return "オンライン開催";
      } else {
        return "形式不明";
      }
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
        形式: {{ showType(isOnsite, isOnline) }} <br>
        開催場所: {{ place }} <br>
        URL: {{ url }} <br>
      </div>
    </div>
    <footer class="card-footer">
      <a :href="getURL()" class="card-footer-item">候補時間編集</a>
    </footer>
  </div>
</template>
