<script setup>
import { ref } from "vue";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";

const title = ref();
const description = ref();
const datetime = ref();
const type = ref();
const place = ref();
const url = ref();

const datetimeNum = ref(0);
const datetimeModelList = ref(["datetime0"]);

function Register() {
  if (title.value == undefined) {
    alert(true);
  } else {
    alert(false);
  }
  // alert(description.value);
  // alert(datetime.value);
  // alert(type.value);
  // alert(place.value);
  // alert(url.value);
}

function AddDateTimeColumn() {
  datetimeNum.value++;
  let newDateTimeModel = "datetime" + datetimeNum.value;
  datetimeModelList.value.push(newDateTimeModel);
}

function DeleteDateTimeColumn() {
  datetimeNum.value--;
  datetimeModelList.value.pop();
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
              <div v-for="(value, key) in datetimeModelList" :key="key">
                <Datepicker
                  v-model="datetimeModelList[key]"
                  range
                  multiCalendars
                />
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button class="button is-primary" @click="AddDateTimeColumn">
                  新規作成
                </button>
              </p>
              <p class="control">
                <button class="button is-light" @click="DeleteDateTimeColumn">
                  戻る
                </button>
              </p>
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
                  to="/dashboard"
                  class="button is-primary"
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