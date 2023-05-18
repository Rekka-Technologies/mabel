import { createApp } from 'vue'
import './style.css'

import App from './App.vue'
import { createRouter } from './router'

// Font Awesome Icons
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {library} from "@fortawesome/fontawesome-svg-core";
import {fas} from "@fortawesome/free-solid-svg-icons";


// Create a new Vue app
const app = createApp(App)

// Set default font awesome icon set to Solid
library.add(fas) 

// Mount the router and plugins to the Application
app
  .use(createRouter(app))
  .component('font-awesome-icon', FontAwesomeIcon)
  .mount('#app')

