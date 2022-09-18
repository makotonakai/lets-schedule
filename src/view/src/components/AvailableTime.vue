<script>
export default {
  props: {
    start_time: String,
    end_time: String
  }
};
</script>
<template>
  <div class="card">
    <div class="card-content">
      <div class="meeting-description">
        {{ start_time }} - {{ end_time }} <br>
      </div>
    </div>
  </div>
</template>
