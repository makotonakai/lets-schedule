<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import DefaultMeeting from "../components/meetings/DefaultMeeting.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    DefaultMeeting
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

        await axios.get(`http://localhost:1323/api/restricted/meetings/${this.UserId}`, {
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
                  <DefaultMeeting
                    :title="meeting.title"
                    :description="meeting.description"
                    :type="meeting.type"
                    :place="meeting.place"
                    :url="meeting.url"
                  ></DefaultMeeting>
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
