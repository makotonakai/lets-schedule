<script>
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import {AddNewElement, DeleteLastElement, CreateCandidateTimeList, CreateDateTimeJSONList} from "../utils/CandidateTime"
import {CreateParticipantJSONList, GetHost, GetParticipantList} from "../utils/Participant"
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import axios from "axios";
import { BadRequestStatus } from "../utils/StatusCode.js";

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
      Type: "",
      Place: "",
      Url: "",
      Hour: "",
      DatetimeList:[""],
      DateTimeJSONList:[],
      Host:"",
      ParticipantList:[""],
      ParticipantJSONList:[],
      MeetingId: parseInt(this.$route.params['id']),
      ErrorMessageList: [],
      IsOnsite: false,
      IsOnline: false,
    }
  },
  mounted() {
      this.Loadinfo();
      this.LoadCandidateTime();
      this.LoadParticipant();
  },

  methods: {
    async Loadinfo() {
      await axios.get(`${process.env.HOST}:${process.env.PORT}/api/restricted/meetings/${this.MeetingId}`, {
          headers: {
            Authorization: `Bearer ${this.Token}`,
          },
        })
        .then((response) => {
          let data = response.data;
          this.Title = data["title"];
          this.Description = data["description"];
          this.IsOnsite = data["is_onsite"];
          this.IsOnline = data["is_online"];
          this.Type = this.ChooseType(this.IsOnsite, this.IsOnline);
          this.Place = data["place"];
          this.Url = data["url"];
          this.Host = data["host"];
          this.Hour = data["hour"]
          this.DatetimeList = CreateCandidateTimeList(data);
        })
        .catch((err) => {
          console.log(err);
      });
    },
    async LoadCandidateTime() {
      await axios.get(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/user/${this.UserId}/meeting/${this.MeetingId}`, {
          headers: { 
            Authorization: `Bearer ${this.Token}`
          }
        })
        .then((response) => {
          this.DatetimeList = CreateCandidateTimeList(response.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async LoadParticipant() {
      await axios.get(`${process.env.HOST}:${process.env.PORT}/api/restricted/participants/${this.MeetingId}`, {
          headers: { 
            Authorization: `Bearer ${this.Token}`
          }
        })
        .then((response) => {
          let allParticipantList = response.data;
          this.Host = GetHost(allParticipantList);
          this.ParticipantList = GetParticipantList(allParticipantList);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async Edit() {
      await this.EditBasicInfo();
      await this.EditCandidateTime();
      await this.EditParticipants();
    },
    async EditBasicInfo(){
      this.ErrorMessageList = [];
      this.UpdateType(this.Type);
      await axios.put(`${process.env.HOST}:${process.env.PORT}/api/restricted/meetings/${this.MeetingId}`, {  
        title: this.Title,
        description: this.Description,
        is_onsite: this.IsOnsite,
        is_online: this.IsOnline,
        place: this.Place,
        url: this.Url,
        is_confirmed: false,
        hour: parseInt(this.Hour)
      },
      {
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data);
      })
      .catch((err) => {
        console.log(err);
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessageList = [...err.response.data];
        };
      });
    },
    async EditCandidateTime(){
      this.DateTimeJSONList = CreateDateTimeJSONList(this.DatetimeList, this.UserId, this.MeetingId)
      console.log(this.DateTimeJSONList);

      await axios.put(`${process.env.HOST}:${process.env.PORT}/api/restricted/candidate_times/user/${this.UserId}/meeting/${this.MeetingId}`, this.DateTimeJSONList,{
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data);
        })
      .catch((err) => {
        console.log(err);
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessageList = [...err.response.data];
        };
      });
    },

    async EditParticipants() {
      this.ParticipantJSONList = CreateParticipantJSONList(this.Host, this.ParticipantList, this.MeetingId)

      await axios.put(`${process.env.HOST}:${process.env.PORT}/api/restricted/participants/meeting/${this.MeetingId}`, this.ParticipantJSONList,{
        headers: { 
          Authorization: `Bearer ${this.Token}`
        }
      })
      .then((response) => {
        console.log(response.data);
        })
      .catch((err) => {
        console.log(err);
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessageList.push(err.response.data);
        };
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

    ChooseType(isOnsite, isOnline) {
      if (isOnsite && isOnline) {
        return "ハイブリッド開催";
      } else if (isOnsite) {
        return "現地開催"
      } else if (isOnline) {
        return "オンライン開催"
      } else {
        return "形式不明"
      }
    },

    UpdateType(type) {
      if (type == "ハイブリッド開催") {
        this.IsOnsite = true;
        this.IsOnline = true;
      } else if (type == "現地開催") {
        this.IsOnsite = true;
        this.IsOnline = false;
      } else if (type == "オンライン開催") {
        this.IsOnsite = false;
        this.IsOnline = true;
      } else {
        this.IsOnsite = false;
        this.IsOnline = false;
      }
    }

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
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="number"
                  placeholder="E.g. Alice"
                  v-model="Hour"
                />
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

            <div class="message is-danger">
              <div v-for="(ErrorMessage, index) in ErrorMessageList"
              :key="index"
              > 
                {{ ErrorMessage }}
              </div>
            </div>
            
            <div class="field is-grouped">
              <p class="control">
                <button type="button" @click="Edit" class="button is-success">
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

<style scoped>
.message.is-danger {
  background-color: transparent;
}

.message.is-danger div {
  font-weight: bold;
  font-size: 16px;
  color: white;
}
</style>
