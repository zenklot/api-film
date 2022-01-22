import "bootstrap/dist/css/bootstrap.css";

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import { film } from './store/film.js'
import { createStore } from 'vuex'
const store = createStore({
	modules: {
        film
	}
});


import axios from 'axios'
import VueAxios from 'vue-axios'
axios.defaults.baseURL = 'http://localhost:3000';
axios.defaults.headers.post['Content-Type'] = 'application/json';
axios.defaults.headers.put['Content-Type'] = 'application/json';

import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

import { BootstrapIconsPlugin } from 'bootstrap-icons-vue';

createApp(App).use(store).use(router).use(VueAxios, axios).use(VueSweetalert2).use(BootstrapIconsPlugin).mount('#app')
import "bootstrap/dist/js/bootstrap.js";