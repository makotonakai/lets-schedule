<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import CandidateTime from "../components/CandidateTime.vue";
import AvailableTime from "../components/AvailableTime.vue";
import Datepicker from "@vuepic/vue-datepicker";
import { CreateCandidateTimeDict, CreateAvailableTimeList} from "../utils/CandidateTime"
import { BadRequestStatus } from "../utils/StatusCode.js";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    CandidateTime,
    AvailableTime,
    Datepicker
  },
  mounted() {
    this.MeetingId = this.$route.params['id'];
    this.getMeetings();
    this.getAvailableTime();
      
  },
  data() {
    return {
      Token: $cookies.get("token"),
      UserId: parseInt($cookies.get("user_id")),
      UserName: $cookies.get("user_name"),
      MeetingId: "",
      CandidateTimeDict: {},
      AvailableTimeList: [],
      FinalAvailableTimeList: [""],
      ErrorMessage: ""
    }
  },
  methods: {
    async getMeetings() {
      // api/restricted/candidate_times/meeting
        await axios
      .get(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/meeting/${this.MeetingId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        this.CandidateTimeDict = CreateCandidateTimeDict(response.data);
      })
      .catch((err) => {
        console.log(err);
      });
    },
    async getAvailableTime() {
        await axios
      .get(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/available-time/${this.MeetingId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        console.log(response.data)
        this.AvailableTimeList = CreateAvailableTimeList(response.data)
      })
      .catch((err) => {
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessage = err.response.data;
        };
        console.log(err);
      });
    },
    async RegisterAvailableTime() {
      const ActualStartTime = FinalAvailableTimeList[0][0];
      const ActualEndTime = FinalAvailableTimeList[0][1];
      const FormattedActualStartTime = convertDate(ActualStartTime);
      const FormattedActualEndTime = convertDate(ActualEndTime);
        await axios
      .put(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/available-time/${this.MeetingId}`, {
        "actual_start_time": FormattedActualStartTime,
        "actual_end_time": FormattedActualEndTime
      },
      {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        console.log(response.data)
      })
      .catch((err) => {
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessage = err.response.data;
        };
        console.log(err);
      });
    },
    convertDate(dateStr) {
      // Parse the date string to a Date object
      const parsedDate = new Date(dateStr);

      // Format the date to "YYYY-MM-DD HH:MM:SS"
      const year = parsedDate.getFullYear();
      const month = String(parsedDate.getMonth() + 1).padStart(2, '0');
      const day = String(parsedDate.getDate()).padStart(2, '0');
      const hours = String(parsedDate.getHours()).padStart(2, '0');
      const minutes = String(parsedDate.getMinutes()).padStart(2, '0');
      const seconds = String(parsedDate.getSeconds()).padStart(2, '0');

      return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
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
              <div class='column is-7'>
                <div id="app">
                  <div class="box">
                    <b>参加者</b>
                    <li
                    v-for="(value, key) in CandidateTimeDict"
                    :key="key"
                    >
                    <CandidateTime
                      :user_name="key"
                      :time_period_list="value"
                    ></CandidateTime>
                    </li>
                    <br>
                    <b>ミーティング可能時間</b>
                      <p class="help is-danger">
                        {{ ErrorMessage }}
                      </p>
                    <AvailableTime
                      :available_time_list="AvailableTimeList"
                    ></AvailableTime>
                    <br>
                    <b>ミーティング時間</b>
                    <br>

                     <!-- <Datepicker
                        v-model="AvailableTime"
                        range
                        multiCalendars
                      /> -->
                      <div v-for="(value, key) in FinalAvailableTimeList" :key="key">
                        <Datepicker
                          v-model="FinalAvailableTimeList[key]"
                          range
                          multiCalendars
                        />
                      </div>
                      <div class="field is-grouped">
                        <p class="control">
                          <button type="button" @click="RegisterAvailableTime" class="button is-success">
                            編集
                          </button>
                        </p>
                        <p class="control">
                          <router-link to="/meeting/dashboard" class="button is-light">
                            戻る
                          </router-link>
                        </p>
                    </div>
                  </div>
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
