<script>
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import { AddNewElement, DeleteLastElement, CreateDateTimeJSONList } from "../utils/CandidateTime";
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
      Token: VueCookies.get("token"),
      UserId: parseInt(VueCookies.get("user_id")),
      MeetingId: parseInt(this.$route.params.id),
      DatetimeList: [""],
      DateTimeJSONList: [],
      ErrorMessageList: []
    };
  },
  methods: {
    async Register() {
      await this.RegisterCandidateTime();
    },
    async RegisterCandidateTime() {
      this.ErrorMessageList = []; // エラーメッセージリストを初期化

      this.DateTimeJSONList = CreateDateTimeJSONList(this.DatetimeList, this.UserId, this.MeetingId);
      console.log(this.DateTimeJSONList);

      try {
        const response = await axios.post(
          `${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/new`,
          this.DateTimeJSONList,
          {
            headers: {
              Authorization: `Bearer ${this.Token}`
            }
          }
        );
        console.log(response.data);
        this.goToDashboard();
      } catch (error) {
        console.log(error);
        if (error.response && error.response.status === BadRequestStatus) {
          this.ErrorMessageList = error.response.data.map((message) => message);
        }
      }
    },
    goToDashboard() {
      this.$router.push("/meeting/dashboard");
    },
    AddDateTime() {
      AddNewElement(this.DatetimeList);
    },
    DeleteDateTime() {
      DeleteLastElement(this.DatetimeList);
    }
  }
};
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
              <p class="help is-danger">
                <li
                  v-for="(ErrorMessage, index) in ErrorMessageList"
                  :key="index"
                >
                  {{ ErrorMessage }}
                </li>
              </p>
              <div v-for="(value, key) in DatetimeList" :key="key">
                <Datepicker v-model="DatetimeList[key]" range multiCalendars />
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

            <div class="field is-grouped">
              <p class="control">
                <button type="button" @click="RegisterCandidateTime" class="button is-success">
                  新規作成
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
