import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
import directives from './directives'
const app = createApp(App)
installElementPlus(app)
app.use(store).use(router).use(directives).mount('#app')
