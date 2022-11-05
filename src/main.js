import Vue from 'vue'
import App from './App.vue' 
import './assets/css/style.css'
import VueResource from 'vue-resource';
import VueRouter from 'vue-router'
import Router from './routes'
import Notifications from "vt-notifications";

Vue.config.productionTip = false
Vue.use(VueResource)
Vue.use(VueRouter)
Vue.use(Notifications);


const router = new VueRouter({
  routes: Router,
  mode: 'history'
});

new Vue({
  render: h => h(App),
  router: router
}).$mount('#app')
