<script>
import VueCookies from "vue-cookies";
import axios from "axios";
import DashboardHeader from "../components/header/DashboardHeader.vue";
import CandidateTime from "../components/CandidateTime.vue";
import AvailableTime from "../components/AvailableTime.vue";
import { CreateCandidateTimeDict, CreateAvailableTimeList} from "../utils/CandidateTime"
import { BadRequestStatus } from "../utils/StatusCode.js";

export default {
  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  components: {
    DashboardHeader,
    CandidateTime,
    AvailableTime
  },
  mounted() {
    this.MeetingId = this.$route.params['id'];
    this.getMeetings();
    this.getAvailableTime();
      
  },
  data() {
    return {
      Token: $cookies.get("token"),
      UserId: parseInt($cookies.get("user_id")),
      UserName: $cookies.get("user_name"),
      MeetingId: "",
      CandidateTimeDict: {},
      AvailableTimeList: [],
      ErrorMessage: ""
    }
  },
  methods: {
    async getMeetings() {
        await axios
      .get(`http://localhost:1323/api/restricted/candidate_times/meeting/${this.MeetingId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        this.CandidateTimeDict = CreateCandidateTimeDict(response.data);
      })
      .catch((err) => {
        console.log(err);
      });
    },
    async getAvailableTime() {
        await axios
      .get(`http://localhost:1323/api/restricted/candidate_times/available-time/${this.MeetingId}`, {
        headers: {
          Authorization: `Bearer ${this.Token}`,
        },
      })
      .then((response) => {
        console.log(response.data)
        this.AvailableTimeList = CreateAvailableTimeList(response.data)
      })
      .catch((err) => {
        if (err.response.status == BadRequestStatus) {
          this.ErrorMessage = "No available time found";
        };
        console.log(err);
      });
    }
  }
}
</script>

<template>
  <div>
    <header>
      <DashboardHeader></DashboardHeader>
      <section class="hero is-primary is-fullheight">
        <div class="hero-body">
          <div class="container">
            <div class="columns is-centered">
              <div class='column is-7'>
                <div id="app">
                  <div class="box">
                    <b>参加者</b>
                    <li
                    v-for="(value, key) in CandidateTimeDict"
                    :key="key"
                    >
                    <CandidateTime
                      :user_name="key"
                      :time_period_list="value"
                    ></CandidateTime>
                    </li>
                    <br>
                    <b>ミーティング可能時間</b>
                      <p class="help is-danger">
                        {{ ErrorMessage }}
                      </p>
                    <AvailableTime
                      :available_time_list="AvailableTimeList"
                    ></AvailableTime>
                    <br>
                    <b>ミーティング時間</b>
                    <br>
                    <select id="start-time-month">
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                    </select>
                    /
                    <select id="start-time-day">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                      <option value="23">24</option>
                      <option value="23">25</option>
                      <option value="23">26</option>
                      <option value="23">27</option>
                      <option value="23">28</option>
                      <option value="23">29</option>
                      <option value="23">30</option>
                      <option value="23">31</option>
                    </select>
                    &nbsp
                    <select id="start-time-hour">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                    </select>
                    :
                    <select id="start-time-minute">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                      <option value="23">24</option>
                      <option value="23">25</option>
                      <option value="23">26</option>
                      <option value="23">27</option>
                      <option value="23">28</option>
                      <option value="23">29</option>
                      <option value="23">30</option>
                      <option value="23">31</option>
                      <option value="23">32</option>
                      <option value="23">33</option>
                      <option value="23">34</option>
                      <option value="23">35</option>
                      <option value="23">36</option>
                      <option value="23">37</option>
                      <option value="23">38</option>
                      <option value="23">39</option>
                      <option value="23">40</option>
                      <option value="23">41</option>
                      <option value="23">42</option>
                      <option value="23">43</option>
                      <option value="23">44</option>
                      <option value="23">45</option>
                      <option value="23">46</option>
                      <option value="23">47</option>
                      <option value="23">48</option>
                      <option value="23">49</option>
                      <option value="23">50</option>
                      <option value="23">51</option>
                      <option value="23">52</option>
                      <option value="23">53</option>
                      <option value="23">54</option>
                      <option value="23">55</option>
                      <option value="23">56</option>
                      <option value="23">57</option>
                      <option value="23">58</option>
                      <option value="23">59</option>
                    </select>
                    ~
                    <select id="end-time-month">
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                    </select>
                    /
                    <select id="end-time-day">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                      <option value="23">24</option>
                      <option value="23">25</option>
                      <option value="23">26</option>
                      <option value="23">27</option>
                      <option value="23">28</option>
                      <option value="23">29</option>
                      <option value="23">30</option>
                      <option value="23">31</option>
                    </select>
                    &nbsp
                    <select id="end-time-hour">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                    </select>
                    :
                    <select id="end-time-minute">
                      <option value="0">0</option>
                      <option value="1">1</option>
                      <option value="2">2</option>
                      <option value="3">3</option>
                      <option value="4">4</option>
                      <option value="5">5</option>
                      <option value="6">6</option>
                      <option value="7">7</option>
                      <option value="8">8</option>
                      <option value="9">9</option>
                      <option value="10">10</option>
                      <option value="11">11</option>
                      <option value="12">12</option>
                      <option value="13">13</option>
                      <option value="14">14</option>
                      <option value="15">15</option>
                      <option value="16">16</option>
                      <option value="17">17</option>
                      <option value="18">18</option>
                      <option value="19">19</option>
                      <option value="20">20</option>
                      <option value="21">21</option>
                      <option value="22">22</option>
                      <option value="23">23</option>
                      <option value="23">24</option>
                      <option value="23">25</option>
                      <option value="23">26</option>
                      <option value="23">27</option>
                      <option value="23">28</option>
                      <option value="23">29</option>
                      <option value="23">30</option>
                      <option value="23">31</option>
                      <option value="23">32</option>
                      <option value="23">33</option>
                      <option value="23">34</option>
                      <option value="23">35</option>
                      <option value="23">36</option>
                      <option value="23">37</option>
                      <option value="23">38</option>
                      <option value="23">39</option>
                      <option value="23">40</option>
                      <option value="23">41</option>
                      <option value="23">42</option>
                      <option value="23">43</option>
                      <option value="23">44</option>
                      <option value="23">45</option>
                      <option value="23">46</option>
                      <option value="23">47</option>
                      <option value="23">48</option>
                      <option value="23">49</option>
                      <option value="23">50</option>
                      <option value="23">51</option>
                      <option value="23">52</option>
                      <option value="23">53</option>
                      <option value="23">54</option>
                      <option value="23">55</option>
                      <option value="23">56</option>
                      <option value="23">57</option>
                      <option value="23">58</option>
                      <option value="23">59</option>
                    </select>
                    &nbsp
                    <button is-light>時間決定</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </header>
  </div>
</template>

<style>
li{
list-style: none;
}
</style>
