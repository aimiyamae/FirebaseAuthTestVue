<template>
  <div class="signin">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar color="#EEEEEE">
                <v-toolbar-title>Login form</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
                <div class="sign-in">
                  <h1 class="title is-5">Sign in</h1>
                  <div id="firebaseui-auth-container"></div>
                </div>
                <h3>電話番号を入力</h3>
                <input v-model="phoneNumber">
                <div id="recaptch"></div>
                <button id="linkPhoneNumberBtn" @click="sendSmsVerification">確認番号を送信</button>
                <v-form>
                  <v-text-field
                    prepend-icon="dialpad"
                    name="dial"
                    label="電話ラベル"
                    type="number"
                    v-model="phoneNumber"
                  ></v-text-field>
                  <div id="recaptch"></div>
                  <v-text-field
                    prepend-icon="email"
                    name="email"
                    label="Email"
                    type="text"
                    v-model="email"
                  ></v-text-field>
                  <v-text-field
                    id="password"
                    prepend-icon="lock"
                    name="password"
                    label="Password"
                    type="password"
                    v-model="password"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="#EEEEEE" @click="signIn">Sign in</v-btn>
                <v-btn id="linkPhoneNumberBtn" @click="sendSmsVerification">確認番号を送信</v-btn>
                <v-btn @click="confirmVerification">確認</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
    <!-- <v-btn @click="authGoogle" outline midium fab color="black">
      <i class="fab fa-google fa-2x"></i>
    </v-btn>
    <v-btn @click="authFaceBook" outline midium fab color="black">
      <i class="fab fa-facebook fa-2x"></i>
    </v-btn>
    <v-btn @click="authTwitter" outline midium fab color="black">
      <i class="fab fa-twitter-square fa-2x"></i>
    </v-btn>-->
    <!-- <p>You don't have an account?
      <router-link to="/signup">create account now!!</router-link>
    </p>-->
  </div>
</template>

<script>
import firebase from "firebase";
// import firebase from "@/plugins/firebase";
import "firebaseui/dist/firebaseui.css";
export default {
  name: "Signin",
  data: function() {
    return {
      email: "",
      password: "",
      auth: {
        email: "test@gmail.com",
        phoneNumber: ""
      },
      phoneNumber: "",
      recaptchaVerifier: null,
      confirmationResult: null,
      waitingVerify: false,
      verificationCode: "",
      user: null
    };
  },
  computed: {
    state() {
      if (!this.auth.email) {
        return "notLoggedIn";
      }
      if (this.auth.email && !this.waitingVerify && !this.auth.phoneNumber) {
        return "onlyEmail";
      }
      if (this.auth.email && this.waitingVerify && !this.auth.phoneNumber) {
        return "waitingVerify";
      }
      return "emailAndPhoneNumber";
    }
  },
  mounted() {
    var firebaseui = require("firebaseui");
    var ui = new firebaseui.auth.AuthUI(firebase.auth());
    ui.start("#firebaseui-auth-container", {
      credentialHelper: firebaseui.auth.CredentialHelper.NONE,
      callbacks: {
        signInSuccessWithAuthResult: function(authResult, redirectUrl) {
          // User successfully signed in.
          // Return type determines whether we continue the redirect automatically
          // or whether we leave that to developer to handle.
          return true;
        },
        uiShown: function() {
          // The widget is rendered.
          // Hide the loader.
          document.getElementById("loader").style.display = "none";
        }
      },
      // Will use popup for IDP Providers sign-in flow instead of the default, redirect.
      // signInFlow: 'popup',
      signInSuccessUrl: "MyPage",
      signInOptions: [
        // Leave the lines as is for the providers you want to offer your users.
        {
          provider: firebase.auth.EmailAuthProvider.PROVIDER_ID,
          requireDisplayName: false,
          signInMethod:
            firebase.auth.EmailAuthProvider.EMAIL_LINK_SIGN_IN_METHOD
        },
        {
          provider: firebase.auth.PhoneAuthProvider.PROVIDER_ID,
          defaultCountry: "JP"
        }
      ],
      // Terms of service url.
      tosUrl: "<your-tos-url>",
      // Privacy policy url.
      privacyPolicyUrl: "<your-privacy-policy-url>"
    });
    const _this = this;
    firebase
      .auth()
      .getRedirectResult()
      .then(result => {
        if (result.credential) {
          const user = result.user;
          _this.user = user;
          if (user.phoneNumber) {
            _this.auth.phoneNumber = user.phoneNumber;
          }
        }
      });
    this.recaptchaVerifier = new firebase.auth.RecaptchaVerifier("recaptch");
    // this.recaptchaVerifier = new firebase.auth.RecaptchaVerifier(
    //   "linkPhoneNumberBtn",
    //   {
    //     size: "invisible"
    //   }
    // );
  },
  methods: {
    async sendSmsVerification() {
      console.log(this.phoneNumber);
      try {
        const confirmationResult = await firebase
          .auth()
          .signInWithPhoneNumber(this.phoneNumber, this.recaptchaVerifier);
        this.confirmationResult = confirmationResult;
        this.waitingVerify = true;
        console.log(this.this.confirmationResult);
      } catch (error) {
        // console.error(error)
      }
    },
    async confirmVerification() {
      const credential = firebase.auth.PhoneAuthProvider.credential(
        this.confirmationResult.verificationId,
        this.verificationCode
      );
      const userCred = await this.user.linkAndRetrieveDataWithCredential(
        credential
      );
      this.auth.phoneNumber = userCred.user.phoneNumber;
    },
    signIn: function() {
      firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then(
          res => {
            res.user.getIdToken().then(idToken => {
              localStorage.setItem("jwt", idToken.toString());
            });
            this.$router.push("/");
          },
          err => {
            alert(err.message);
          }
        );
    }
    // authGoogle: function() {
    //   const provider = new firebase.auth.GoogleAuthProvider();
    //   firebase
    //     .auth()
    //     .signInWithPopup(provider)
    //     .then(
    //       res => {
    //         res.user.getIdToken().then(idToken => {
    //           localStorage.setItem("jwt", idToken.toString());
    //         });
    //         this.$router.push("/");
    //       },
    //       err => {
    //         alert(err.message);
    //       }
    //     );
    // },
    // authTwitter: function() {
    //   const provider = new firebase.auth.TwitterAuthProvider();
    //   firebase
    //     .auth()
    //     .signInWithPopup(provider)
    //     .then(
    //       res => {
    //         res.user.getIdToken().then(idToken => {
    //           localStorage.setItem("jwt", idToken.toString());
    //         });
    //         this.$router.push("/");
    //       },
    //       err => {
    //         alert(err.message);
    //       }
    //     );
    // },

    // authFaceBook: function() {
    //   const provider = new firebase.auth.FacebookAuthProvider();
    //   firebase
    //     .auth()
    //     .signInWithPopup(provider)
    //     .then(
    //       res => {
    //         res.user.getIdToken().then(idToken => {
    //           localStorage.setItem("jwt", idToken.toString());
    //         });
    //         this.$router.push("/");
    //       },
    //       err => {
    //         alert(err.message);
    //       }
    //     );
    // }
  }
};
</script>
<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>