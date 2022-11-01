<script>
export default {
  props: {
    user_name: String,
    time_period_list: Array
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
        ユーザー名: {{ user_name }} <br>
        <div v-for="(time_period, key) in time_period_list" :key="key">
          {{ this.getStartTime(time_period) }} -  {{ this.getEndTime(time_period) }} 
        </div>
      </div>
    </div>
  </div>
</template>
