<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import CandidateTime from "../components/CandidateTime.vue";
import {CreateCandidateTimeDict} from "../utils/CandidateTime"

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    CandidateTime
  },
  mounted() {
    this.MeetingId = this.$route.params['id'];
    this.getMeetings();
      
  },
  data() {
    return {
      Token: $cookies.get("token"),
      UserId: parseInt($cookies.get("user_id")),
      UserName: $cookies.get("user_name"),
      MeetingId: "",
      CandidateTimeJSONList: [],
      CandidateTimeDict: {}
    }
  },
  methods: {
    async getMeetings() {
        await axios
      .get(`http://localhost:1323/api/restricted/candidate_times/meeting/${this.MeetingId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        this.CandidateTimeJSONList = response.data;
        console.log(this.CandidateTimeJSONList);
        this.CandidateTimeDict = CreateCandidateTimeDict(this.CandidateTimeJSONList);
        console.log(this.CandidateTimeDict);
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
              <div id="app">
                <div class="box">
                  <b>参加者</b>
                  <li
                  v-for="(value, key, index) in CandidateTimeDict"
                  :key="index"
                  >
                   <CandidateTime
                    :user_name="key"
                    :start_time="value.start_time"
                    :end_time="value.end_time"
                  ></CandidateTime>
                  </li>
                  <br>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </header>
  </div>
</template>

<style>
li{
list-style: none;
}
</style>
