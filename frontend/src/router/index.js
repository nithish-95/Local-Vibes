
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Register from '../views/Register.vue';
import Login from '../views/Login.vue';
import CreateEvent from '../views/CreateEvent.vue';
import EditEvent from '../views/EditEvent.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/create-event',
    name: 'CreateEvent',
    component: CreateEvent,
  },
  {
    path: '/edit-event/:id',
    name: 'EditEvent',
    component: EditEvent,
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
