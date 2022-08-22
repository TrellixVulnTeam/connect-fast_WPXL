import { defaultConfig, plugin } from "@formkit/vue";
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

const app = createApp({})

createApp(App).use(plugin, defaultConfig({  theme: 'genesis'})).use(router).mount("#app");

