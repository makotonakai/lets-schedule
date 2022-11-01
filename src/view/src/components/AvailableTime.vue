<script>
export default {
  props: {
    available_time_list: Array
  },
  methods: {
    getStartTime(time_period) {
      return time_period["start_time"]
    },
    getEndTime(time_period) {
      return time_period["end_time"]
    },
  }
};
</script>
<template>
  <div class="card">
    <div class="card-content">
      <div class="meeting-description">
        <div v-for="(available_time, key) in available_time_list" :key="key">
          {{ this.getStartTime(available_time) }} -  {{ this.getEndTime(available_time) }} 
        </div>
      </div>
    </div>
  </div>
</template>
