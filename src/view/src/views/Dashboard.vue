<script setup>
import { ref, onMounted } from "vue";
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import Meeting from "../components/Meeting.vue";

const userName = $cookies.get("user_name");
const jwtToken = $cookies.get("token");

const meetings = ref([]);
const meeting = ref("hoge");
const role = ref("host");
const hasResponded = ref(true);

let meetingIdList = ref([]);
let meetingInfoList = ref([]);

onMounted(() => {
  axios
    .get(
      `http://localhost:1323/api/restricted/participants/username/${userName}`,
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
        },
      }
    )
    .then((response) => {
      for (let idx = 0; idx < response.data.length; idx++) {
        let meetingId = parseInt(response.data[idx]["meeting_id"]);
        meetingIdList.value.push(meetingId);
      }

      for (let idx = 0; idx < meetingIdList.value.length; idx++) {
        let meetingId = meetingIdList.value[idx];
        let url = `http://localhost:1323/api/restricted/meetings/${meetingId}`;

        axios
          .get(url, {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          })
          .then((response) => {
            console.log(response.data);
            meetingInfoList.value.push(response.data);
          })
          .catch((err) => {
            console.log(err);
          });
      }
    })
    .catch((err) => {
      console.log(err);
    });
});
</script>

_<template>
  <div>
    <header>
      <DashboardHeader></DashboardHeader>
      <section class="hero is-primary is-fullheight">
        <div class="hero-body">
          <div class="container">
            <div class="columns is-centered">
              <div class="column is-half">
                <li
                  v-for="meetingInfo in meetingInfoList"
                  :key="meetingInfo.id"
                >
                  <Meeting
                    :title="meetingInfo.title"
                    :description="meetingInfo.description"
                  ></Meeting>
                </li>
              </div>
            </div>
          </div>
        </div>
      </section>
    </header>
  </div>
</template>