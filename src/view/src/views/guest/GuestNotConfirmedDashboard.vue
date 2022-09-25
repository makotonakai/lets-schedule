<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../../components/header/DashboardHeader.vue";
import GuestNotConfirmedMeeting from "../../components/meetings/guest/GuestNotConfirmedMeeting.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    GuestNotConfirmedMeeting,
  },

  data() {
    return {
      Token: $cookies.get("token"),
      UserId: parseInt($cookies.get("user_id")),
      UserName: $cookies.get("user_name"),
      Meetings: []
    }
  },

  mounted() {
      this.getMeetings();
  },

  methods: {
      async getMeetings() {
        await axios
      .get(`http://localhost:1323/api/restricted/meetings/guest/not-confirmed/${this.UserId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        this.Meetings = response.data;
      })
      .catch((err) => {
        console.log(err);
      });
    }
  }
}
</script>

<template>
  <div>
    <header>
      <DashboardHeader></DashboardHeader>
      <section class="hero is-primary is-fullheight">
        <div class="hero-body">
          <div class="container">
            <div class="columns is-centered">
              <div class="column is-half is-4">
                <li
                  v-for="meeting in Meetings"
                  :key="meeting.id"
                >
                  <GuestNotConfirmedMeeting
                    :id="meeting.id"
                    :title="meeting.title"
                    :description="meeting.description"
                    :type="meeting.type"
                    :place="meeting.place"
                    :url="meeting.url"
                  ></GuestNotConfirmedMeeting>
                  <br>
                </li>
              </div>
            </div>
          </div>
        </div>
      </section>
    </header>
  </div>
</template>

<style scoped>
li {
  list-style: none;
}
</style>
