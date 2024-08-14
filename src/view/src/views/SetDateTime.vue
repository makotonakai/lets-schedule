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
      ErrorMessageList: [],
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
          for(let x = 0; x < error.response.data.length; x++){
              let errorMessage = error.response.data[x];
              this.ErrorMessageList.push(errorMessage);
            };
        };
        console.log(err);
      });
    },
    async RegisterAvailableTime() {
      console.log(this.FinalAvailableTimeList);

      let FormattedActualStartTime = "0001-01-01T00:00:00Z";
      let FormattedActualEndTime = "0001-01-01T00:00:00Z";

      if(Array.isArray(this.FinalAvailableTimeList[0]) && this.FinalAvailableTimeList[0].length == 2) {
        if (this.FinalAvailableTimeList[0][0] != null) {
          const ActualStartTime = this.FinalAvailableTimeList[0][0];
          FormattedActualStartTime = this.convertDate(ActualStartTime);
        }
        
        if (this.FinalAvailableTimeList[0][1] != null) {
          const ActualEndTime = this.FinalAvailableTimeList[0][1];
          FormattedActualEndTime = this.convertDate(ActualEndTime);
        }
      }

      console.log(FormattedActualStartTime);
      console.log(FormattedActualEndTime);
      
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
      .catch((error) => {
        for(let x = 0; x < error.response.data.length; x++){
            let errorMessage = error.response.data;
            this.ErrorMessageList.push(errorMessage);
        };
        console.log(error);
      });
    },
    convertDate(dateStr) {
    // Parse the date string to a Date object
    const parsedDate = new Date(dateStr);

    // Format the date to "YYYY-MM-DDTHH:MM:SSZ"
    const year = parsedDate.getFullYear();
    const month = String(parsedDate.getMonth() + 1).padStart(2, '0');
    const day = String(parsedDate.getDate()).padStart(2, '0');
    const hours = String(parsedDate.getHours()).padStart(2, '0');
    const minutes = String(parsedDate.getMinutes()).padStart(2, '0');
    const seconds = String(parsedDate.getSeconds()).padStart(2, '0');

    // Get the timezone offset in hours and minutes
    const timezoneOffset = -parsedDate.getTimezoneOffset();
    const offsetHours = String(Math.floor(Math.abs(timezoneOffset) / 60)).padStart(2, '0');
    const offsetMinutes = String(Math.abs(timezoneOffset) % 60).padStart(2, '0');
    const offsetSign = timezoneOffset >= 0 ? '+' : '-';

    return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}${offsetSign}${offsetHours}:${offsetMinutes}`;
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
                    <AvailableTime
                      :available_time_list="AvailableTimeList"
                    ></AvailableTime>
                    <br>
                    <b>ミーティング時間</b>
                    <br>

                     <p class="help is-danger">
                      <li
                        v-for="(ErrorMessage, index) in ErrorMessageList"
                        :key="index"
                      > 
                        {{ ErrorMessage }}
                      </li>
                    </p>
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
