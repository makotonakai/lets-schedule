<script>
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import {AddNewElement, DeleteLastElement, CreateDateTimeJSONList} from "../utils/CandidateTime"
import {CreateParticipantJSONList} from "../utils/Participant"
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
      Title: "",
      Description: "",
      Type: "現地開催",
      Place: "なし",
      Url: "なし",
      Hour: "",
      DatetimeList:[""],
      DateTimeJSONList:[],
      Host:"",
      ParticipantList:[""],
      ParticipantJSONList:[],
      MeetingId: null
    }
  },

  methods: {
    async Register() {
      await this.RegisterBasicInfo();
      await this.RegisterCandidateTime();
      await this.RegisterParticipants();
    },

    async RegisterBasicInfo(){

      await axios.post(`${process.env.HOST}/api/restricted/meetings/new`, {  
        title: this.Title,
        description: this.Description,
        type: this.Type,
        place: this.Place,
        url: this.Url,
        is_confirmed: false
      },
      {
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        this.MeetingId = response.data["id"];
        console.log(response.data)
        })
      .catch((err) => {
        console.log(err);
      });
    },
    async RegisterCandidateTime(){

      this.DateTimeJSONList = CreateDateTimeJSONList(this.DatetimeList, this.UserId, this.MeetingId),

      await axios.post(`${process.env.HOST}/api/restricted/candidate_times/new`, this.DateTimeJSONList,{
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data)
        })
      .catch((err) => {
        console.log(err);
      });
    },
    async RegisterParticipants() {

      this.ParticipantJSONList = CreateParticipantJSONList(this.Host, this.ParticipantList, this.MeetingId)

      await axios.post(`${process.env.HOST}/api/restricted/participants/new`, this.ParticipantJSONList,{
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data)
        })
      .catch((err) => {
        console.log(err);
      });
    },

    AddDateTime(){
      AddNewElement(this.DatetimeList);
    },

    DeleteDateTime(){
      DeleteLastElement(this.DatetimeList);
    },

    AddParticipant(){
      AddNewElement(this.ParticipantList);
    },

    DeleteParticipant(){
      DeleteLastElement(this.ParticipantList);
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
              <label class="label">タイトル</label>
              <div class="control">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. 会議"
                  v-model="Title"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">概要</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="textarea"
                  placeholder="E.g. 新プロジェクトの会議"
                  v-model="Description"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">時間</label>
              <div class="control">
                <textarea
                  class="input"
                  type="text"
                  placeholder="E.g. 1"
                  v-model="Hour"
                ></textarea>
              </div>
            </div>


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

            <div class="field">
              <label class="label">主催者</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. Alice"
                  v-model="Host"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">参加者</label>
              <div v-for="(value, key) in ParticipantList" :key="key">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. Alice"
                  v-model="ParticipantList[key]"
                />
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button class="button is-light" @click="AddParticipant">
                  入力欄を追加
                </button>
              </p>
              <p class="control">
                <button class="button is-light" @click="DeleteParticipant">
                  戻る
                </button>
              </p>
            </div>

            <div class="field">
              <label class="label">形式</label>
              <div class="control">
                <div class="select">
                  <select v-model="Type">
                    <option>現地開催</option>
                    <option>オンライン開催</option>
                    <option>ハイブリッド開催</option>
                  </select>
                </div>
              </div>
            </div>

            <div class="field">
              <label class="label">開催場所</label>
              <div class="control">
                <textarea
                  class="input"
                  type="text"
                  placeholder="E.g. 会議室"
                  v-model="Place"
                ></textarea>
              </div>
            </div>

            <div class="field">
              <label class="label">会議URL</label>
              <div class="control">
                <textarea
                  class="input"
                  type="text"
                  placeholder="E.g. https://kaigi-zoom.com"
                  v-model="Url"
                ></textarea>
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <router-link
                  to="/meeting/dashboard"
                  class="button is-light"
                  @click="Register"
                >
                  新規作成
                </router-link>
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
