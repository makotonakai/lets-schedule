<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useCookie } from "vue-cookie-next";
import axios from "axios";
import LoginHeader from "../components/header/LoginHeader.vue";

const username = ref("");
const password = ref("");
const router = useRouter();
const { setCookie, getCookie, removeCookie } = useCookie();

function login() {
  axios
    .post("/api/login", {
      username: username.value,
      password: password.value,
    })
    .then((response) => {
      console.log(response.data);
      let token = response.data["token"];
      setCookie("token", token);
      router.push("/dashboard");
    })
    .catch((err) => {
      console.log(err);
    });
}
</script>

<template>
  <div>
    <header>
      <LoginHeader></LoginHeader>
    </header>
    <section class="hero is-primary is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="columns is-centered">
            <div class="column is-5-tablet is-4-desktop is-3-widescreen">
              <form action="" class="box" @submit.prevent="onSubmit">
                <div class="field">
                  <label for="" class="label">Username </label>
                  <div class="control has-icons-left">
                    <input
                      v-model="username"
                      type="username"
                      placeholder="e.g. lets-schedule"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fas fa-envelope"></i>
                    </span>
                  </div>
                </div>
                <div class="field">
                  <label for="" class="label">Password</label>
                  <div class="control has-icons-left">
                    <input
                      v-model="password"
                      type="password"
                      placeholder="*******"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fas fa-lock"></i>
                    </span>
                  </div>
                </div>
                <div class="field">
                  <button @click="login" class="button is-success">
                    Login
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>