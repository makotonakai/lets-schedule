<script>
import SignupHeader from "../components/header/SignupHeader.vue";
import { BadRequestStatus } from "../utils/StatusCode.js";
import axios from "axios";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    SignupHeader
  },
  data() {
    return {
      EmailAddress: "",
      UserName: "",
      Password: "",
      ErrorMessageList: []
    }
  },

  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event listeners in templates.
  methods: {
    async SignUp(){
      await this.CreateUser()
    },
    async CreateUser(){
      // api/signup
      await axios.post(`${process.env.HOST}/YXBpL3NpZ251cA==`, {
        email_address: this.EmailAddress,
        user_name: this.UserName,
        password: this.Password,
        is_admin: false,
        can_login: true
      })
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        if (error.response.status == BadRequestStatus) {
          for(let x = 0; x < error.response.data.length; x++){
              let errorMessage = error.response.data[x];
              this.ErrorMessageList.push(errorMessage);
            };
          };
        }
      )
    },
    GoToLoginPage() {
      this.$router.push("/login");
    },
    GoToSignUpPage() {
      this.$router.push("/signup");
    }
  }
}
</script>

<template>
  <div>
    <header>
      <SignupHeader></SignupHeader>
    </header>
    <section class="hero is-primary is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="columns is-centered">
            <div class="column is-5-tablet is-4-desktop is-3-widescreen">
              <form action="" class="box">
                <div class="field">
                  <label for="" class="label">Email Address</label>
                  <div class="control has-icons-left">
                    <input
                      v-model="EmailAddress"
                      type="username"
                      placeholder="e.g. user@email.com"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fa fa-envelope"></i>
                    </span>
                  </div>
                </div>
                <div class="field">
                  <label for="" class="label">Username</label>
                  <div class="control has-icons-left">
                    <input
                      v-model="UserName"
                      type="username"
                      placeholder="e.g. user"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fa fa-envelope"></i>
                    </span>
                  </div>
                </div>
                <div class="field">
                  <label for="" class="label">Password</label>
                  <div class="control has-icons-left">
                    <input
                      v-model="Password"
                      type="password"
                      placeholder="*******"
                      class="input"
                      required
                    />
                    <span class="icon is-small is-left">
                      <i class="fa fa-lock"></i>
                    </span>
                  </div>
                  <p class="help is-danger">
                    <li
                  v-for="ErrorMessage in ErrorMessageList"
                  :key="ErrorMessage.id">
                    {{ ErrorMessage }}
                    </li>
                  </p>
                </div>
                <div class="field">
                  <button type="button" @click="SignUp()" class="button is-success">Sign up</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
