<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router"
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../../components/header/DashboardHeader.vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import axios from "axios";

const jwtToken = $cookies.get("token");
const route = useRoute()

const userId = parseInt($cookies.get("user_id"));
const meetingId = parseInt(route.params.id);

const datetimeNum = ref(0);
let datetimeObjectList = ref([""]);

onMounted(() => {
  axios
    .get(
      `http://localhost:1323/api/restricted/candidate_times/${userId}/${meetingId}`,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then((response) => {
      for (let key in response.data) {
        let start_date = Date(response.data[key]["start_time"]);
        let end_date = Date(response.data[key]["end_time"]);
        let datetimeObject = [start_date, end_date]
        datetimeObjectList.value.push(datetimeObject)
      }
    })
    .catch((err) => {
      console.log(err);
    });
 
});

function GetCandidateTimeJSONList() {

  let datetimeKeyList = Object.keys(datetimeObjectList.value);
  let candidateTimeList = [];

  for (let datetimeKey in datetimeKeyList) {
    let datetimeObjectDict = datetimeObjectList.value[datetimeKey];

    let startTimeObject = datetimeObjectDict[0];
    let endTimeObject = datetimeObjectDict[1];

    let startTime = startTimeObject.toISOString().split('Z')[0] + '+09:00';
    let endTime = endTimeObject.toISOString().split('Z')[0] + '+09:00';

    let candidateTime = {
      meeting_id: meetingId,
      user_id: userId,
      start_time: startTime,
      end_time: endTime,
    };
    candidateTimeList.push(candidateTime);
  }

  return candidateTimeList;
}

function Edit() {

  let candidateTimeList = GetCandidateTimeJSONList();
  axios
    .put(
      `http://localhost:1323/api/restricted/candidate_times/${userId}/${meetingId}`,
      candidateTimeList,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then(function (response) {
      console.log(response);
    })
    .catch(function (error) {
      console.log(error.response.data);
    });
}

function AddDateTimeColumn() {
  datetimeNum.value++;
  let newDateTime = "";
  datetimeObjectList.value.push(newDateTime);
}

function DeleteDateTimeColumn() {
  datetimeNum.value--;
  datetimeObjectList.value.pop();
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
              <div v-for="(value, key) in datetimeObjectList" :key="key">
                <Datepicker
                  v-model="datetimeObjectList[key]"
                  range
                  multiCalendars
                />
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button class="button is-light" @click="AddDateTimeColumn">
                  入力欄を追加
                </button>
              </p>
              <p class="control">
                <button class="button is-light" @click="DeleteDateTimeColumn">
                  入力欄を削除
                </button>
              </p>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <router-link
                  to="/dashboard"
                  class="button is-light"
                  @click="Edit"
                >
                  編集
                </router-link>
              </p>
              <p class="control">
                <router-link to="/meeting/guest/responded" class="button is-light">
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
