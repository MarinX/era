import 'core-js/stable';
import 'regenerator-runtime/runtime';
import '@mdi/font/css/materialdesignicons.css';
import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);

import App from './App.vue';
import router from './router'

Vue.config.productionTip = false;
Vue.config.devtools = true;


import Wails from '@wailsapp/runtime';
Wails.Init(() => {
	new Vue({
		vuetify: new Vuetify({
			theme: {dark:true}
		}),
		render: h => h(App),
		router,
		mounted() {
			this.$router.replace('/contacts')
		},
	}).$mount('#app');
});


