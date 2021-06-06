import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
// import store from './store/state'
import installElementPlus from "./plugins/element";
import directives from "./directives";
import { StateSymbol, CreateState } from "./store/state";
const app = createApp(App);
installElementPlus(app);
app.use(router).use(directives);
app.provide(StateSymbol, CreateState());
// console.log("启动应用");
app.mount("#app");
