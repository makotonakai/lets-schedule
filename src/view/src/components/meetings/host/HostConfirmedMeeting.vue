<script>
import { google, outlook, ics } from "calendar-link"

export default {
  props: {
    title: String,
    description: String,
    isOnsite: Boolean,
    isOnline: Boolean,
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
    },
    goToOutlookCalendar() {
      return outlook(this.event)
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
      <a :href="this.getICSFile()" class="card-footer-item">iCal</a>
      <a :href="this.goToGoogleCalendar()" target="_blank" class="card-footer-item">Google Calendar</a>
      <a :href="this.goToOutlookCalendar()" target="_blank" class="card-footer-item">Outlook Calendar</a>
    </footer>
  </div>
</template>
