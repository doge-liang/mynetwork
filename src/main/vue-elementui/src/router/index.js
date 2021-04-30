import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Market from '@/views/Market.vue'
import Trade from '@/views/Trade.vue'
// import RSI2Strategy from "@/views/RSI2Strategy";

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/strategy/:strategyId/market',
    component: Market
  },
  {
    path: '/strategy/:strategyId/trade',
    component: Trade
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
