<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import { BadRequestStatus, UnauthorizedStatus } from "../utils/StatusCode.js";
import LoginHeader from "../components/header/LoginHeader.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    LoginHeader
  },
  data() {
    return {
      UserName: "",
      Password: "",
      ErrorMessageList: []
    }
  },

  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event listeners in templates.
  methods: {
    async Login(){
      // /api/login
      // エラーのロギング
      // URLを共通の公開鍵 (ホスト鍵の公開鍵?) で暗号化
      // APIサーバーはHTTPS化
      await axios.post(`${process.env.HOST}:${process.env.PORT}/api/login`, {
        user_name: this.UserName,
        password: this.Password,
      })
      .then((response) => {
        this.setCredential(response);
        this.goToDashboard();
        }
      )
      .catch((error) => {
        console.log(error);
        if (error.response.status == BadRequestStatus) {
          for(let x = 0; x < error.response.data.length; x++){
              let errorMessage = error.response.data[x];
              this.ErrorMessageList.push(errorMessage);
            };
          } else if (error.response.status == UnauthorizedStatus) {
            let errorMessage = error.response.data;
            this.ErrorMessageList.push(errorMessage);
        };
      });
    },
    setCredential(response) {
      $cookies.set("token", response.data["token"], 0);
      $cookies.set("user_id", response.data["id"]);
      $cookies.set("user_name", response.data["user_name"]);
    },
    goToDashboard() {
      this.$router.push("/meeting/dashboard");
    }
  }
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
              <form action="" class="box">
                <div class="field">
                  <label for="" class="label">Username </label>
                  <div class="control has-icons-left">
                    <input
                      v-model="UserName"
                      type="UserName"
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
                      v-model="Password"
                      type="Password"
                      placeholder="*******"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fas fa-lock"></i>
                    </span>
                  </div>
                </div>

                 <div class="message is-danger">
                    <div v-for="(ErrorMessage, index) in ErrorMessageList"
                    :key="index"
                    > 
                      {{ ErrorMessage }}
                    </div>
                  </div>

                <div class="field">
                  <button type="button" @click="Login" class="button is-success">
                    Login
                  </button>
                </div>
                <router-link to="/send-email" class="button is-success">
                    Forget your password?
                </router-link>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>


