<template>
	<div class='hello'>
		<div class="text-xs-center">
			<!-- golangニュースページに飛ぶ -->
			<v-btn @click="hrefURL" outline color="indigo">NEWS</v-btn>
			<!-- golang http から値を取ってくるボタンPublic -->
			<v-btn @click="reservationForm" outline color="indigo">予約フォーム</v-btn>
			<v-btn @click="apiPrivate" outline color="indigo">private商品一覧</v-btn>
			<v-btn @click="myPage" outline color="indigo">マイページ</v-btn>
		</div>
    <div class="signin">
		<v-content>
			<v-container fluid fill-height>
			<v-layout align-center justify-center>
				<v-flex xs12 sm8 md4>
				<v-card class="elevation-12">
					<!-- <v-toolbar color="#2ECCFA">
					<v-toolbar-title ></v-toolbar-title>
					<v-spacer> 
					</v-spacer>-->
					<v-toolbar color="#2ECCFA">
						<v-toolbar-title ><h1>{{ msg }}</h1></v-toolbar-title>
					</v-toolbar>
					<v-card-text>
						<p>
						<h4>2019年2月14日　　　　　　　newsページ追加</h4>
					
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

import axios from 'axios'
import firebase from 'firebase'
export default {
name: 'HelloWorld',
data () {
	return {
		picker: new Date().toISOString().substr(0, 10),
		msg: 'NEWS'
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
	methods: {
		myPage: function () {
			// window.location.href='http://localhost:8080/news'
			this.$router.push('/')
		},
		hrefURL: function () {
			// window.location.href='http://localhost:8080/news'
			this.$router.push('/news')
		},
		inquiryForm: function () {
			open( "https://docs.google.com/forms/d/1riQvax43mOky4luvP8K5AKPZNvYEJ3umOcHY_teXTWg/edit", "_blank" ) ;
			// window.location.href='https://docs.google.com/forms/d/1riQvax43mOky4luvP8K5AKPZNvYEJ3umOcHY_teXTWg/edit'
		},
		reservationForm: async function () {
			open( "https://docs.google.com/forms/d/e/1FAIpQLSdO02JHSaXkqgsDo8gKJR0oaHDCZs5eWZbmc_1de2kf7g2tNA/viewform", "_blank" ) ;
		},
		apiPrivate: async function () {
			let res = await axios.get('http://localhost:8000/private', {
				headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
			})
			this.msg = res.data
		}
  	}
}
</script>