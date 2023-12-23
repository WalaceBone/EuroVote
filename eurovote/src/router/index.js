import { createRouter, createWebHistory } from 'vue-router';
import HelloWorldVue from '@/components/HelloWorld.vue';
import D3ChartVue from '@/components/D3Chart.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HelloWorldVue,
  },
  {
    path: '/d3',
    name: 'd3',
    component: D3ChartVue,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
