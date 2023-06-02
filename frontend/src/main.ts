import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './style.css'

import App from './App.vue'
import { router } from './router'

// Font Awesome Icons
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {library} from "@fortawesome/fontawesome-svg-core";
import {fas} from "@fortawesome/free-solid-svg-icons";

import VueTailwindDatepicker from 'vue-tailwind-datepicker'

// Create a new Vue app
const app = createApp(App)

// Setup Pinia
const pinia = createPinia()

// Set default font awesome icon set to Solid
library.add(fas) 

// Mount the router and plugins to the Application
app
  .use(pinia)
  .use(router)
  .use(VueTailwindDatepicker)
  .component('font-awesome-icon', FontAwesomeIcon)
  .mount('#app')

