<script setup>
import { ref } from "vue";
import { useRoute } from "vue-router"
import VueCookies from "vue-cookies";
import VueTagsInput from "@johmun/vue-tags-input";
import DashboardHeader from "../../components/header/DashboardHeader.vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import axios from "axios";

const jwtToken = $cookies.get("token");
const route = useRoute()

const datetimeNum = ref(0);
let datetimeObjectList = ref([""]);

function GetMonthByName(monthName) {
  if (monthName == "Jan") {
    return "01";
  } else if (monthName == "Feb") {
    return "02";
  } else if (monthName == "Mar") {
    return "03";
  } else if (monthName == "Apr") {
    return "04";
  } else if (monthName == "May") {
    return "05";
  } else if (monthName == "Jun") {
    return "06";
  } else if (monthName == "Jul") {
    return "07";
  } else if (monthName == "Aug") {
    return "08";
  } else if (monthName == "Sep") {
    return "09";
  } else if (monthName == "Oct") {
    return "10";
  } else if (monthName == "Nov") {
    return "11";
  } else if (monthName == "Dec") {
    return "12";
  }
}

function ConvertStringToDateTime(string) {
  let stringWithSpace = string.split(" ");
  let month = GetMonthByName(stringWithSpace[1]);
  let day = stringWithSpace[2];
  let year = stringWithSpace[3];
  let time = stringWithSpace[4];

  let datetime = year + "-" + month + "-" + day + " " + time;
  return datetime;
}

function GetCandidateTimeJSONList() {

  let datetimeKeyList = Object.keys(datetimeObjectList.value);
  let candidateTimeList = [];

  for (let datetimeKey in datetimeKeyList) {
    let datetimeObjectDict = datetimeObjectList.value[datetimeKey];

    let startTimeObject = datetimeObjectDict[0];
    let endTimeObject = datetimeObjectDict[1];

    let startTimeString = startTimeObject.toString();
    let endTimeString = endTimeObject.toString();

    let startTime = ConvertStringToDateTime(startTimeString);
    let endTime = ConvertStringToDateTime(endTimeString);

    const meetingId = parseInt(route.params.id);
    let userId = parseInt($cookies.get("user_id"));
    let isHost = false;
    let hasResponded = true;

    let candidateTime = {
      meeting_id: meetingId,
      user_id: userId,
      is_host: isHost,
      has_responded: hasResponded,
      start_time: startTime,
      end_time: endTime,
    };
    candidateTimeList.push(candidateTime);
  }

  return candidateTimeList;
}

function Register() {

  let candidateTimeList = GetCandidateTimeJSONList();
  axios
    .post(
      "http://localhost:1323/api/restricted/candidate_times/new",
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
      console.log(error);
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
                  戻る
                </button>
              </p>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <router-link
                  to="/dashboard"
                  class="button is-light"
                  @click="Register"
                >
                  新規作成
                </router-link>
              </p>
              <p class="control">
                <router-link to="/dashboard" class="button is-light">
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