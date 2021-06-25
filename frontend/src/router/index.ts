import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import ProductDB from '../page/ProductDB.vue';
import Training from '../page/Training.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Main',
    component: Main,
  },
  {
    path: '/product_db',
    name: 'ProductDB',
    component: ProductDB,
  },
  {
    path: '/training',
    name: 'Training',
    component: Training,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
