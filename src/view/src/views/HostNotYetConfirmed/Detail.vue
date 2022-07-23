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

let title = ref();
let description = ref();
let type = ref();
let place = ref();
let url = ref();

const datetimeNum = ref(0);
let datetimeObjectList = ref([""]);

let host = ref();
let participantObjectList = ref([]);

onMounted(() => {

  axios
    .get(
      `http://localhost:1323/api/restricted/meetings/${meetingId}`,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then((response) => {
      // console.log(response.data);
      let meeting = response.data;
      title.value = meeting["title"];
      description.value = meeting["description"];
      type.value = MakeMeetingType(meeting["type"]);
      place.value = meeting["place"];
      url.value = meeting["url"];
    })
    .catch((err) => {
      console.log(err);
    });

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
        datetimeObjectList.value.push([start_date, end_date])
      }
    })
    .catch((err) => {
      console.log(err);
    });

    axios
    .get(
      `http://localhost:1323/api/restricted/participants/host/${meetingId}`,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then((response) => {
      let data = response.data[0];
      host.value = data["user_name"];
    })
    .catch((err) => {
      console.log(err);
    });

    axios
    .get(
      `http://localhost:1323/api/restricted/participants/guest/${meetingId}`,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then((response) => {
      for (let participant in response.data) {
          participantObjectList.value.push(participant["user_name"]);
        }
    })
    .catch((err) => {
      console.log(err);
    });
});

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

    let meetingId = parseInt($cookies.get("meeting_id"));
    let userId = parseInt($cookies.get("user_id"));

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

function GetParticipantJSONList() {
  let participantListString = participantObjectList.value.toString();
  let participantList = participantListString.split(" ");

  let participantJSONList = [];
  let hostJSON = {
    meeting_id: meetingId,
    user_id: host.value[0]["id"],
    is_host: true,
    has_responded: true,
  };
  participantJSONList.push(hostJSON);

  for (let idx = 0; idx < participantList.length; idx++) {
    let participantJSON = {
      meeting_id: meetingId,
      user_id: 1,
      is_host: false,
      has_responded: false,
    };
    participantJSONList.push(participantJSON);
  }
  return participantJSONList;
}

function GetMeetingType(meetingType) {
  if (meetingType == "現地開催") {
    return "physical";
  } else if (meetingType == "オンライン") {
    return "online";
  } else {
    return "hybrid";
  }
}

function MakeMeetingType(type) {
  if (type == "physical") {
    return "現地開催";
  } else if (type == "online") {
    return "オンライン";
  } else {
    return "ハイブリッド";
  }
}

function Edit() {
  axios
    .put(
      `http://localhost:1323/api/restricted/meetings/${meetingId}`,
      {
        title: title.value,
        description: description.value,
        type: GetMeetingType(type.value),
        place: place.value,
        url: url.value,
        is_confirmed: false
      },
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
      console.log(error);
    });

  // let participantJSONList = GetParticipantJSONList(participantObjectList.value);
  // axios
  //   .put(
  //     `http://localhost:1323/api/restricted/participants/${meetingId}`,
  //     participantJSONList,
  //     {
  //       headers: {
  //         Authorization: `Bearer ${jwtToken}`,
  //       },
  //     }
  //   )
  //   .then(function (response) {
  //     console.log(participantJSONList);
  //   })
  //   .catch(function (error) {
  //     console.log(error);
  //   });
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
              <label class="label">タイトル</label>
              <div class="control">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. 会議"
                  v-model="title"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">概要</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="textarea"
                  placeholder="E.g. 新プロジェクトの会議"
                  v-model="description"
                />
              </div>
            </div>

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

            <div class="field">
              <label class="label">主催者</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. Alice"
                  v-model="host"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">参加者</label>
              <div class="control has-icons-left has-icons-right">
                <input
                  class="input"
                  type="text"
                  placeholder="E.g. Bob Charlie"
                  v-model="participantObjectList"
                />
              </div>
            </div>

            <div class="field">
              <label class="label">形式</label>
              <div class="control">
                <div class="select">
                  <select v-model="type">
                    <option>現地開催</option>
                    <option>オンライン</option>
                    <option>ハイブリッド</option>
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
                  v-model="place"
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
                  v-model="url"
                ></textarea>
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <router-link
                  to="/meeting/host/not-yet-confirmed"
                  class="button is-light"
                  @click="Edit"
                >
                  編集
                </router-link>
              </p>
              <p class="control">
                <router-link to="/meeting/host/not-yet-confirmed" class="button is-light">
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
