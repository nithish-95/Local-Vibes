import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import './css/main.css'; // Import your Tailwind CSS file

import { session } from './session';

const app = createApp(App);
app.use(router);
app.provide('session', session); // Provide the session object globally
app.mount('#app');