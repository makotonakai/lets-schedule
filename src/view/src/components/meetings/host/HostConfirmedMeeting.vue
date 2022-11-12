<script>
export default {
  props: {
    title: String,
    description: String,
    type: String,
    place: String,
    url: String,
    start_time: String,
    end_time: String
  },
  data () {
    return {
      startTimeForGoogleScholar: this.changeTimeForGoogleCalendar(this.start_time),
      endTimeForGoogleScholar: this.changeTimeForGoogleCalendar(this.end_time)
    }
  },
  methods: {
    changeTimeForGoogleCalendar(fullTime){
      let [year, month, dayTimeDifference] = fullTime.split("-")
      let [day, timeDifference] = dayTimeDifference.split("T")
      let [time, difference] = timeDifference.split("+")
      let [hour, minute, second] = time.split(":")
      let timeForGoogleCalendar = year + month + day + hour + minute + second
      return timeForGoogleCalendar
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
        開始時間 {{ start_time }}<br>
        終了時間 {{ end_time }}<br>
      </div>
    </div>
    <footer class="card-footer">
     <a href="#" class="card-footer-item">iCal</a>
      <a href="http://www.google.com/calendar/event?action=TEMPLATE&" onclick="location.href=this.href+text='this.title'&detail='this.description'&location='this.place'&dates='startTimeForGoogleScholar'/'endTimeForGoogleScholar';return false;" target="_blank" class="card-footer-item">Google Calendar</a>
      <a href="#" class="card-footer-item">Outlook Calendar</a>
    </footer>
  </div>
</template>
