'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"https://ureteku-mvp.firebaseio.com"',
  YOUR_KEY: '"AIzaSyDK1OEQt9qfgc46JVBaNiZz8YI7t3pvd0Q"',
  YOUR_DOMAIN: '"ureteku-mvp.firebaseapp.com"',
  YOUR_ID: '"ureteku-mvp"',
  YOUR_BUCKET_ID: '"ureteku-mvp.appspot.com"',
  YOUR_SENDER_ID: '"761747549302"',
})
