<script>
import { google, outlook, office365, yahoo, ics } from "calendar-link"

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
  data() {
    return {
      event: {
        title: this.title,
        location: this.place,
        description: this.description,
        start: this.start_time,
        end: this.end_time,
        url: this.url
      }
    }
  },
  methods: {
    getICSFile() {
      return ics(this.event)
    },
    goToGoogleCalendar() {
      return google(this.event)
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
     <a :href="this.getICSFile()" class="card-footer-item">iCal</a>
      <a :href="this.goToGoogleCalendar()" target="_blank" class="card-footer-item">Google Calendar</a>
      <a href="#" class="card-footer-item">Outlook Calendar</a>
    </footer>
  </div>
</template>
