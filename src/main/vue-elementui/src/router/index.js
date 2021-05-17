import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Market from '@/views/Market.vue'
import Trade from '@/views/Trade.vue'
import Login from '@/views/Login.vue'
import PrivateMarket from '@/views/PrivateMarket.vue'
import PrivateTrade from '@/views/PrivateTrade.vue'
// import RSI2Strategy from "@/views/RSI2Strategy";

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      requireAuth: true,
    }
  },
  {
    path: '/strategy/:strategyId/market',
    component: Market,
    meta: {
      requireAuth: true,
    }
  },
  {
    path: '/strategy/:strategyId/trade',
    component: Trade,
    requireAuth: true,
  },
  {
    path: '/strategy/:strategyId/private-market',
    component: PrivateMarket,
    requireAuth: true,
  },
  {
    path: '/strategy/:strategyId/private-trade',
    component: PrivateTrade,
    requireAuth: true,
  },
  {
    path: '/login',
    component: Login,
  }]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

router.beforeEach((to, from, next) => {
  let islogin = sessionStorage.getItem("isLogin");
  console.log(islogin);
  console.log(to.path);

  if (to.path == "/login") {
    if (islogin) {
      next('/home');
    } else {
      next();
    }
  } else {
    // requireAuth:可以在路由元信息指定哪些页面需要登录权限
    if (to.meta.requireAuth && islogin) {
      next();
    } else {
      next("/login");
    }
  }
});
