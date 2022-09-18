<script>
export default {
  props: {
    user_name: String,
    start_time: String,
    end_time: String
  }
};
</script>
<template>
  <div class="card">
    <div class="card-content">
      <div class="meeting-description">
        ユーザー名: {{ user_name }} <br>
        開始時間: {{ start_time }} <br>
        終了時間: {{ end_time }} <br>
      </div>
    </div>
  </div>
</template>
