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
                            <label id="Selected" v-if="type === 'A'">{{ infoMsg }}</label>
                            <label id="selectDate" v-else-if="type === 'B'">
                              {{ picker }}
                              <br>
                              {{ infoMsg }}
                            </label>
                            <label id="selectDate" v-else-if="type === 'C'">
                              {{ picker }}
                              <br>
                              {{ selectedTime }}
                            </label>
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
                            <v-flex style="padding-left:0px; height:60px;" xs9>
                              <span class="grey--text">＊電話番号は数字のみ入力してください</span>
                            </v-flex>

                            <label></label>
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
                            ></v-text-field>
                            <v-flex xs12 sm6 d-flex>
                              <v-select :items="items" label="商品量の目安" solo></v-select>
                            </v-flex>
                            <div v-if="type === 'A'"></div>
                            <div v-else-if="type === 'B'||type === 'C'">
                              <!-- モーダル内容 -->
                              <div class="text-xs-center">
                                <v-dialog v-model="dialog" width="500">
                                  <v-btn slot="activator" color="red lighten-2" dark>時間を表示するアクション</v-btn>

                                  <v-card>
                                    <v-card-title
                                      class="headline grey lighten-3 text-xs-center"
                                      primary-title
                                    >
                                      {{ picker }}
                                      <br>集荷希望時間帯を選択
                                    </v-card-title>
                                    <p></p>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      dark
                                      value="10:00〜11:00"
                                      v-model="selectedTime"
                                    >10:00〜11:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="11:00〜12:00"
                                      v-model="selectedTime"
                                      dark
                                    >11:00〜12:00</v-btn>
                                    <br>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="12:00〜13:00"
                                      v-model="selectedTime"
                                      dark
                                    >12:00〜13:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="13:00〜14:00"
                                      v-model="selectedTime"
                                      dark
                                    >13:00〜14:00</v-btn>
                                    <br>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="14:00〜15:00"
                                      v-model="selectedTime"
                                      dark
                                    >14:00〜15:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="15:00〜16:00"
                                      v-model="selectedTime"
                                      dark
                                    >15:00〜16:00</v-btn>
                                    <br>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="16:00〜17:00"
                                      v-model="selectedTime"
                                      dark
                                    >16:00〜17:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="17:00〜18:00"
                                      v-model="selectedTime"
                                      dark
                                    >17:00〜18:00</v-btn>
                                    <br>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="18:00〜19:00"
                                      v-model="selectedTime"
                                      dark
                                    >18:00〜19:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="19:00〜20:00"
                                      v-model="selectedTime"
                                      dark
                                    >19:00〜20:00</v-btn>
                                    <br>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="20:00〜21:00"
                                      v-model="selectedTime"
                                      dark
                                    >20:00〜21:00</v-btn>
                                    <v-btn
                                      round
                                      color="primary"
                                      @click="timeAction01"
                                      value="21:00〜22:00"
                                      v-model="selectedTime"
                                      dark
                                    >21:00〜22:00</v-btn>
                                    <br>

                                    <v-divider></v-divider>

                                    <v-card-actions>
                                      <v-spacer></v-spacer>
                                      <v-btn color="primary" flat @click="dialog = false">閉じる</v-btn>
                                    </v-card-actions>
                                  </v-card>
                                </v-dialog>
                              </div>
                              <!-- モーダル内容 -->
                            </div>
                            <div v-else>日付の取得ができませんでした</div>
                          </div>
                        </v-flex>
                        <v-flex xs12 md1>
                          <!-- <v-text-field v-model="email" :rules="emailRules" label="E-mail" required></v-text-field>-->
                        </v-flex>
                        <div>
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
                          </div>
                          <br>
                          <v-card-title class="grey lighten-2" primary-title>以下に該当するアイテムは対象外となります。</v-card-title>
                          <v-card-text class="grey lighten-3">
                            <ul class="text-xs-left">
                              <li>盗品</li>
                              <li>コピー品(偽物/海賊版/不明な物)</li>
                              <li>汚れ、破れ、シミ、匂いなどの状態が酷い、酷いと当社が判断した品</li>
                              <li>食料品(飲料/酒を含む)</li>
                              <li>中〜大型家電(エアコン/テレビ/冷蔵庫/洗濯機/その他)</li>
                              <li>中〜大型家具</li>
                              <li>自動車/オートバイ</li>
                              <li>ペットをはじめとする生き物</li>
                              <li>アダルト関連・公序良俗に反する品</li>
                              <li>その他当社が取り扱い不可と判断した品</li>
                            </ul>
                            <br>
                          </v-card-text>
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
                    <div class="text-xs-center">
                      <v-btn round color="primary" dark>予約する</v-btn>
                    </div>
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
      selectedTime: "",
      dialog: "",
      type: "A",
      msg: "集荷予約",
      infoMsg: "カレンダーを選択してください",
      nameRules: [
        v => !!v || "Name is required",
        v => v.length <= 20 || "20文字以内でお願いします。"
      ],
      dialRules: [
        v => !!v || "dial is required",
        v => v.length <= 11 || "電話番号をご確認ください"
      ],
      addressRules: [v => !!v || "address is required"],
      items: ["1個〜50個", "50個〜100個", "150個〜200個", "200個以上"]
    };
  },
  constructor() {
    this.value = "";
  },
  watch: {
    //日にちを取得した時
    picker: function(val, oldVal) {
      (this.type = "B"),
        (this.infoMsg = "時間を指定してください"),
        console.log("new: %s, old: %s", val, oldVal);
    },
    selectedTime: function(val, oldVal) {
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
    },
    timeAction01: function($event) {
      this.type = "C";
      this.selectedTime = $event.target.value;
      console.log($event.target.value);
    }
  }
};
</script>