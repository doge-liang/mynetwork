import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Market from '@/views/Market.vue'
import Trade from '@/views/Trade.vue'
import Login from '@/views/Login.vue'
import PrivateMarket from '@/views/PrivateMarket.vue'

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/strategy/:strategyId/market',
    component: Market,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/strategy/:strategyId/trade',
    component: Trade,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/strategy/:strategyId/private-market',
    component: PrivateMarket,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/',
    redirect: '/home'
  }]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

router.beforeEach((to, from, next) => {
  let islogin = sessionStorage.getItem("isLogin");
  if (to.matched.some(record => record.meta.requiresAuth)) {
    console.log(islogin);
    console.log(to.path);

    console.log(to.meta.requiresAuth)
    if (islogin) {
      next();
    } else {
      next("/login");
    }
  } else {
    next();
  }
});