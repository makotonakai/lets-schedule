<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import Meeting from "../components/Meeting.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    Meeting
  },

  data() {
    return {
      token: $cookies.get("token"),
      user_id: parseInt($cookies.get("user_id")),
      user_name: $cookies.get("user_name"),
      meetings: []
    }
  },

  mounted() {
      this.getMeetings();
  },

  methods: {
      async getMeetings() {
        await axios
      .get(`http://localhost:1323/api/restricted/meetings/${this.user_id}`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      })
      .then((response) => {
        console.log(response.data);
        this.meetings = response.data;
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
                  v-for="meeting in meetings"
                  :key="meeting.id"
                >
                  <Meeting
                    :title="meeting.title"
                    :description="meeting.description"
                    :type="meeting.type"
                    :place="meeting.place"
                    :url="meeting.url"
                  ></Meeting>
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
