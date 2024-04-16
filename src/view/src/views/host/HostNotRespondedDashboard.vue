<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../../components/header/DashboardHeader.vue";
import HostNotRespondedMeeting from "../../components/meetings/host/HostNotRespondedMeeting.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    HostNotRespondedMeeting
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
        // api/restricted/meetings/host/not-responded/
        await axios
      .get(`${process.env.HOST}/YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvaG9zdC9ub3QtcmVzcG9uZGVk/${this.UserId}`, {
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
                  <HostNotRespondedMeeting
                    :id="meeting.id"
                    :title="meeting.title"
                    :description="meeting.description"
                    :type="meeting.type"
                    :place="meeting.place"
                    :url="meeting.url"
                  ></HostNotRespondedMeeting>
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
