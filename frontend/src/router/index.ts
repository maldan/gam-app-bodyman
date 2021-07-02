import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import ProductDB from '../page/ProductDB.vue';
import Training from '../page/Training.vue';
import Notes from '../page/Notes.vue';

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
  {
    path: '/notes',
    name: 'Notes',
    component: Notes,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
