import { createApp } from 'vue';
import App from './App.vue';
import Store from './store';
import Router from './router';
import UI from '../src/gam_sdk_ui/vue/ui';
import Event from '../src/gam_sdk_ui/vue/event';
import '../src/gam_sdk_ui/vue/main.scss';
import './main.scss';

createApp(App).use(Router).use(UI).use(Event).use(Store).mount('#app');
