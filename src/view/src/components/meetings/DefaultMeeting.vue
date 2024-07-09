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
        URL: <a :href="url" target="_blank">{{ url }}</a> <br>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    description: String,
    isOnsite: Boolean,
    isOnline: Boolean,
    place: String,
    url: String
  },

  methods: {
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

<style scoped>
.card {
  margin-bottom: 1rem;
}

.card-header-title {
  font-weight: bold;
}

.meeting-description {
  line-height: 1.5;
}
</style>


