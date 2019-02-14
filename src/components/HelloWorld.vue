<template>

	<div class='hello'>
			<v-btn @click="signOut" outline color="indigo">Sign out</v-btn>
		<div class="text-xs-center">
			<v-badge color="purple" left overlap >
				<v-icon slot="badge" dark small >done</v-icon>
				<v-icon color="grey lighten-1" large > account_circle </v-icon>
			</v-badge>
		</div>
		<h1>{{ msg }}</h1>
		<h3>{{ name }}</h3>
		<div class="text-xs-center">
			<!-- golangニュースページに飛ぶ -->
			<v-btn @click="hrefURL" outline color="indigo">NEWS</v-btn>
			<!-- golang http から値を取ってくるボタンPublic -->
			<v-btn @click="apiPublic" outline color="indigo">golangから値Public</v-btn>
			<v-btn @click="apiPrivate" outline color="indigo">private商品一覧</v-btn>
		</div>
		<div v-show="state=='notLoggedIn'">
			<button @click="loginWithEmail">ログインtest</button>
		</div>
		<div v-show="state=='onlyEmail'">
			<!-- <p>{{`email: ${auth.email}`}}</p> -->
			<h3>電話番号を入力</h3>
			<input v-model='phoneNumber' />
			<button id='linkPhoneNumberBtn' @click="sendSmsVerification">確認番号を送信</button>
		</div>
		
		<div v-show="state=='waitingVerify'">
			<p>{{`email: ${auth.email}`}}</p>
			<h3>確認番号を入力</h3>
			<input v-model='verificationCode' />
			<button @click="confirmVerification">確認</button>
		</div>

		<div v-show="state=='emailAndPhoneNumber'">
			<p>{{`email: ${auth.email}`}}</p>
			<p>{{`phone number: ${auth.phoneNumber}`}}</p>
		</div>

				<template>
    <div class="signin">
		<v-content>
			<v-container fluid fill-height>
			<v-layout align-center justify-center>
				<v-flex xs12 sm8 md4>
				<v-card class="elevation-12">
					<v-toolbar color="#EEEEEE">
					<v-toolbar-title >マイページ</v-toolbar-title>
					<v-spacer>
					</v-spacer>
					</v-toolbar>
					<v-card-text>
					<v-form>
						<p></p>
						<p>{{`登録ナンバー: ${auth.phoneNumber}`}}</p>
						<p></p>
						<v-btn flat color="black">取り扱い対象外アイテム</v-btn><br />
						<v-btn flat color="black">利用規約</v-btn><br />
						<v-btn flat color="black">プライバシーポリシー</v-btn><br />
						<v-btn flat color="black">古物営業法に基づく表記</v-btn><br />
						<v-btn flat color="black">運営会社</v-btn><br />
						<v-btn @click="inquiryForm" flat color="black">お問い合わせ</v-btn><br />
					</v-form>
					</v-card-text>
					<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="#EEEEEE" @click="signIn">振込申請</v-btn>
					</v-card-actions>
				</v-card>
				</v-flex>
			</v-layout>
			</v-container>
		</v-content>
    </div>
</template>
		<div id="app">
			<h1>{{ msgCalendar }}</h1>
			<v-app id="inspire">
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
			</v-app>
		</div>


  	</div>
</template>

<script>

import axios from 'axios'
import firebase from 'firebase'
export default {
name: 'HelloWorld',
data () {
	return {
		picker: new Date().toISOString().substr(0, 10),
		msg: 'マイページ',
		msgCalendar: '日付を選択',
		name: firebase.auth().currentUser.uid,
		auth: {
			email: 'test@gmail.com',
			phoneNumber: '',
		},
		phoneNumber: '',
		recaptchaVerifier: null,
		confirmationResult: null,
		waitingVerify: false,
		verificationCode: '',
		user: null
	}
},
computed: {
    state() {
		if (!this.auth.email) {
			return 'notLoggedIn'
		}
		if (this.auth.email  && !this.waitingVerify && !this.auth.phoneNumber) {
			return 'onlyEmail'
		}
		if (this.auth.email && this.waitingVerify && !this.auth.phoneNumber) {
			return 'waitingVerify'
		}
			return 'emailAndPhoneNumber'
		},
	},
	async mounted() {
		const _this = this
		firebase.auth().getRedirectResult().then((result) => {
			if (result.credential) {
				const user = result.user
				_this.user = user
				if (user.email) {
				_this.auth.email = user.email
				}
			}
		})

		this.recaptchaVerifier = new firebase.auth.RecaptchaVerifier('linkPhoneNumberBtn', {
		'size': 'invisible',
		})
	},
	methods: {
		signOut: function () {
			firebase.auth().signOut().then(() => {
				localStorage.removeItem('jwt')
				this.$router.push('/signin')
			})
		},
		hrefURL: function () {
			window.location.href='http://localhost:8080/news'
		},
		inquiryForm: function () {
			open( "https://docs.google.com/forms/d/1riQvax43mOky4luvP8K5AKPZNvYEJ3umOcHY_teXTWg/edit", "_blank" ) ;
			// window.location.href='https://docs.google.com/forms/d/1riQvax43mOky4luvP8K5AKPZNvYEJ3umOcHY_teXTWg/edit'
		},
		apiPublic: async function () {
			let res = await axios.get('http://localhost:8000/public')
			this.msg = res.data
		},
		apiPrivate: async function () {
			let res = await axios.get('http://localhost:8000/private', {
				headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
			})
			this.msg = res.data
		},
		loginWithEmail() {
			const provider = new firebase.auth.GoogleAuthProvider()
			firebase.auth().signInWithRedirect(provider)
		},
		async sendSmsVerification() {
			try {
				const confirmationResult = await firebase.auth().signInWithPhoneNumber(this.phoneNumber, this.recaptchaVerifier)
				this.confirmationResult = confirmationResult
				this.waitingVerify = true
			} catch(error) {
				// console.error(error)
			}
		},
		async confirmVerification() {
			const credential = firebase.auth.PhoneAuthProvider.credential(this.confirmationResult.verificationId, this.verificationCode)
			const userCred = await this.user.linkAndRetrieveDataWithCredential(credential)
			this.auth.phoneNumber = userCred.user.phoneNumber
		},
		//カレンダー選択不可の式
		allowedDates: function(val){
			// 今日～100日後までを選べるようにする
			let today = new Date()
			today = new Date(
				today.getFullYear(),
				today.getMonth(),
				today.getDate() + 1
			)
			let maxAllowedDay = new Date()
			maxAllowedDay.setDate(
				today.getDate() + 100
			)
			maxAllowedDay = new Date(
				maxAllowedDay.getFullYear(),
				maxAllowedDay.getMonth(),
				maxAllowedDay.getDate()
			)
			return today <= new Date(val) && new Date(val) <= maxAllowedDay
		}
    // allowedDates: val => parseInt(val.split('-')[2], 10) % 2 === 0
  	}
}

// import axios from 'axios'
// import firebase from 'firebase'
// export default {
//   name: 'HelloWorld',
//   data () {
//     return {
//       msg: 'Welcome to Your Vue.js App',
//       name: firebase.auth().currentUser.email
//     }
//   },
//   methods: {
//     signOut: function () {
//       firebase.auth().signOut().then(() => {
//         localStorage.removeItem('jwt')
//         this.$router.push('/signin')
//       })
//     },
//     apiPublic: async function () {
//       let res = await axios.get('http://localhost:8000/public')
//       this.msg = res.data
//     },
//     apiPrivate: async function () {
//       let res = await axios.get('http://localhost:8000/private', {
//         headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
//       })
//       this.msg = res.data
//     }
//   }
// }
</script>