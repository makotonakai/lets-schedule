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
      EmailAddress: "",
    }
  },
  methods: {
    async SendEmail(){
      await axios.post(`${process.env.HOST}/api/send-email`, {
        email_address: this.EmailAddress
      })
      .then((response) => {
        console.log(response.data);
        },
        )
      .catch((err) => {
        console.log(err);
      });
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
                  <label for="" class="label">Email address</label>
                  <div class="control has-icons-left">
                    <input
                      v-model="EmailAddress"
                      type="EmailAddress"
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
                  <button @click="SendEmail" class="button is-success">
                    Send email
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

<style>
li{
list-style: none;
}
</style>
