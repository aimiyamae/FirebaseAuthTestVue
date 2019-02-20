<template>
  <div class="hello">
    <div class="text-xs-center">
      <!-- golangニュースページに飛ぶ -->
      <v-btn @click="hrefNEWS" outline color="indigo">NEWS</v-btn>
      <v-btn @click="reservationForm" outline color="indigo">予約フォーム</v-btn>
      <v-btn @click="apiPrivate" outline color="indigo">private商品一覧</v-btn>
      <v-btn @click="myPage" outline color="indigo">マイページ</v-btn>
    </div>
    <div class="signin">
      <v-content>
        <v-container fluid fill-height>
          <v-layout align-center justify-center>
            <v-flex xs12 sm8 md10>
              <v-card class="elevation-12">
                <!-- <v-toolbar color="#2ECCFA">
									<v-toolbar-title ></v-toolbar-title>
									<v-spacer> 
                </v-spacer>-->
                <!-- <v-toolbar color="#2ECCFA">-->
                <v-toolbar-title>
                  <h1>{{ msg }}</h1>
                </v-toolbar-title>
                <!-- </v-toolbar>-->
                <v-card-text>
                  <p></p>

                  <!-- フォーム-->
                  <v-form v-model="valid">
                    <v-container>
                      <v-layout row>
                        <v-flex xs12 md5>
                          <label>対象エリア・大阪府堺市</label>
                          <p></p>
                          <div class="form-group mb-5">
                            集荷日時
                            <br>
                            <br>
                            <label id="Selected">{{ strDate }}</label>
                            <br>
                            <label id="selectDate">{{ picker }}</label>
                            <!-- <v-text-field
                              v-model="carender"
                              :rules="dateRules"
                              label="カレンダーを選択してください(TODO.読みより専用にする)"
                              required
                            ></v-text-field>-->
                            <v-text-field
                              v-model="name"
                              :rules="nameRules"
                              :counter="20"
                              label="氏名"
                              required
                            ></v-text-field>

                            <v-text-field
                              prepend-icon="dialpad"
                              v-model="dial"
                              :rules="dialRules"
                              :counter="11"
                              label="電話番号"
                              type="number"
                              required
                            ></v-text-field>

                            <v-text-field
                              v-model="address"
                              :rules="addressRules"
                              label="住所"
                              type="text"
                              required
                            ></v-text-field>

                            <v-text-field
                              v-model="building"
                              :rules="addressRules"
                              label="建物名・部屋番号"
                              type="text"
                              required
                            ></v-text-field>個数はプルダウン
                            <label>電話番号は数字のみ入力してください</label>
                            <div v-if="type === 'A'">A</div>
                            <div v-else-if="type === 'B'">B</div>
                            <div v-else-if="type === 'C'">C</div>
                            <div v-else>Not A/B/C</div>
                          </div>
                        </v-flex>
                        <v-flex xs12 md1>
                          <!-- <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>-->
                        </v-flex>
                        <div id="app">
                          <!-- <h1>{{ msgCalendar }}</h1> -->
                          <div>
                            <v-date-picker
                              locale="jp-ja"
                              v-model="picker"
                              :allowed-dates="allowedDates"
                              class="mt-3"
                              min="2019-01-15"
                              max="2020-03-20"
                              :day-format="date => new Date(date).getDate()"
                            ></v-date-picker>
                            <v-card-title class="grey lighten-2" primary-title>以下に該当するアイテムは対象外となります。</v-card-title>
                            <v-card-title class="grey lighten-2" primary-title>
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>以下に該当するアイテムは対象外となります。
                              <br>
                            </v-card-title>
                          </div>
                        </div>
                        <!-- <div class="form-group mb-5">
                          <label>日付選択</label>
                          <input v-model="birthday" type="date" class="form-control">

                          <v-text-field
                            prepend-icon="dialpad"
                            name="dialpad"
                            label="電話番号入力"
                            type="number"
                            v-model="email"
                          ></v-text-field>
                          <v-text-field v-model="firstname" :rules="nameRules" label="氏名" required></v-text-field>
                        </div>-->
                      </v-layout>
                    </v-container>
                    <p>
                      <input type="submit" value="送信">
                    </p>
                  </v-form>
                </v-card-text>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import firebase from "firebase";
export default {
  data() {
    return {
      valid: false,
      name: "",
      dial: "",
      address: "",
      building: "",
      PickingDate: "",
      picker: "",
      type: "A",
      msg: "集荷予約",
      strDate: "カレンダーを選択してください",
      nameRules: [
        v => !!v || "Name is required",
        v => v.length <= 20 || "20文字以内でお願いします。"
      ],
      dialRules: [
        v => !!v || "dial is required",
        v => v.length <= 11 || "電話番号をご確認ください"
      ],
      addressRules: [v => !!v || "dial is required"]
    };
  },
  watch: {
    //日にちを取得した時
    picker: function(val, oldVal) {
      console.log("new: %s, old: %s", val, oldVal);
    }
  },
  methods: {
    myPage: function() {
      this.$router.push("/");
    },
    hrefNEWS: function() {
      this.$router.push("/news");
    },
    inquiryForm: function() {
      open(
        "https://docs.google.com/forms/d/1riQvax43mOky4luvP8K5AKPZNvYEJ3umOcHY_teXTWg/edit",
        "_blank"
      );
    },
    reservationForm: async function() {
      open(
        "https://docs.google.com/forms/d/e/1FAIpQLSdO02JHSaXkqgsDo8gKJR0oaHDCZs5eWZbmc_1de2kf7g2tNA/viewform",
        "_blank"
      );
    },
    apiPrivate: async function() {
      let res = await axios.get("http://localhost:8000/private", {
        headers: { Authorization: `Bearer ${localStorage.getItem("jwt")}` }
      });
      this.msg = res.data;
    },
    //カレンダー選択不可の式
    allowedDates: function(val) {
      // 今日～100日後までを選べるようにする
      let today = new Date();
      today = new Date(
        today.getFullYear(),
        today.getMonth(),
        today.getDate() + 1
      );
      let maxAllowedDay = new Date();
      maxAllowedDay.setDate(today.getDate() + 100);
      maxAllowedDay = new Date(
        maxAllowedDay.getFullYear(),
        maxAllowedDay.getMonth(),
        maxAllowedDay.getDate()
      );
      return today <= new Date(val) && new Date(val) <= maxAllowedDay;
    }
  }
};
</script>