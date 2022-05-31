<script setup>
import { ref, onMounted } from "vue";
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import Meeting from "../components/Meeting.vue";

const meetings = ref([]);
const meeting = ref("hoge");
const role = ref("host");
const hasResponded = ref(true);

onMounted(() => {
  const id = parseInt($cookies.get("id"));
  const jwtToken = $cookies.get("token");
  const url = `http://localhost:1323/api/restricted/participants/user/${id}`;

  axios
    .get(url, {
      headers: {
        Authorization: `Bearer ${jwtToken}`,
      },
    })
    .then((response) => {
      console.log(response.data);
      meetings.value = response.data;
    })
    .catch((err) => {
      console.log(err);
    });
});
</script>
<template>
  <div>
    <header>
      <DashboardHeader></DashboardHeader>
      <section class="hero is-primary is-fullheight">
        <div class="hero-body">
          <div class="container">
            <div class="columns is-centered">
              <div class="column is-half">
                <Meeting
                  :title="meeting"
                  :role="role"
                  :hasResponded="hasResponded"
                ></Meeting>
              </div>
            </div>
          </div>
        </div>
      </section>
    </header>
  </div>
</template>