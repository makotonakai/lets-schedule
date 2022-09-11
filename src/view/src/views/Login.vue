<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import LoginHeader from "../components/header/LoginHeader.vue";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    LoginHeader
  },
  data() {
    return {
      username: "",
      password: "",
    }
  },

  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event listeners in templates.
  methods: {
    async login(){
      await axios.post("http://localhost:1323/api/login", {
        user_name: this.username,
        password: this.password,
      })
      .then((response) => {
        console.log(response.data);
        this.setCredential(response);
        },
        this.goToDashboard()
        )
      .catch((err) => {
        console.log(err);
      });
    },
    setCredential(response) {
      $cookies.set("token", response.data["token"]);
      $cookies.set("user_id", response.data["id"]);
      $cookies.set("user_name", response.data["user_name"]);
    },
    goToDashboard() {
      this.$router.push("/dashboard");
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


