<script>
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import {AddNewElement, DeleteLastElement, CreateCandidateTimeList, CreateDateTimeJSONList} from "../utils/CandidateTime"
import { BadRequestStatus } from "../utils/StatusCode.js";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import axios from "axios";

export default {
  components: {
    DashboardHeader,
    Datepicker
  },
  data() {
    return {
      Token: $cookies.get("token"),
      UserId: parseInt($cookies.get("user_id")),
      MeetingId: parseInt(this.$route.params['id']),
      DatetimeList:[""],
      DateTimeJSONList:[],
      ErrorMessageList: [],
    }
  },
  mounted() {
    this.LoadInfo();
  },
  methods: {

    async LoadInfo() {
      await axios.get(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/user/${this.UserId}/meeting/${this.MeetingId}`, {
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        this.DatetimeList = CreateCandidateTimeList(response.data);
        console.log(this.DatetimeList)
      })
      .catch((err) => {
        console.log(err);
      });
    },

    async Register() {
      await this.RegisterCandidateTime();
    },

    async RegisterCandidateTime(){

      console.log(this.DatetimeList)

      this.DateTimeJSONList = CreateDateTimeJSONList(this.DatetimeList, this.UserId, this.MeetingId)
      console.log(this.DateTimeJSONList)

      await axios.put(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/user/${this.UserId}/meeting/${this.MeetingId}`, this.DateTimeJSONList,{
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data)
        })
      .catch((error) => {
        console.log(error);
         if (error.response && error.response.status === BadRequestStatus) {
          this.ErrorMessageList = error.response.data.map((message) => message);
        }
      });
    },

    AddDateTime(){
      AddNewElement(this.DatetimeList);
    },

    DeleteDateTime(){
      DeleteLastElement(this.DatetimeList);
    },

  }
}
</script>
<template>
  <div>
    <header>
      <DashboardHeader></DashboardHeader>
    </header>
    <section class="hero is-primary is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="column is-6 is-size-1 has-text-left">

            <div class="field">
              <label class="label">日時</label>
              
              <div v-for="(value, key) in DatetimeList" :key="key">
                <Datepicker
                  v-model="DatetimeList[key]"
                  range
                  multiCalendars
                />
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button class="button is-light" @click="AddDateTime">
                  入力欄を追加
                </button>
              </p>
              <p class="control">
                <button class="button is-light" @click="DeleteDateTime">
                  戻る
                </button>
              </p>
            </div>

             <div class="message is-danger">
              <div v-for="(ErrorMessage, index) in ErrorMessageList"
              :key="index"
              > 
                {{ ErrorMessage }}
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button type="button" @click="RegisterCandidateTime" class="button is-success">
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
    </section>
  </div>
</template>
