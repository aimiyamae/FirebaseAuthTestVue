// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import App from './App'
import router from './router'
import firebase from 'firebase'
import 'vuetify/dist/vuetify.min.css' // Ensure you are using css-loader

Vue.use(Vuetify)
Vue.config.productionTip = false

const config = {
  apiKey: 'AIzaSyDK1OEQt9qfgc46JVBaNiZz8YI7t3pvd0Q',
  authDomain: 'ureteku-mvp.firebaseapp.com',
  databaseURL: 'https://ureteku-mvp.firebaseio.com',
  projectId: 'ureteku-mvp',
  storageBucket: 'ureteku-mvp.appspot.com',
  messagingSenderId: '761747549302'
}
firebase.initializeApp(config)

let app

firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      el: '#app',
      router,
      components: { App },
      template: '<App/>'
    })
  }
})
