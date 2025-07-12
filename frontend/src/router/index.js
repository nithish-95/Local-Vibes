import { createRouter, createWebHistory } from 'vue-router';
import { useSessionStore } from '../stores/session';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue'),
  },
  {
    path: '/create-event',
    name: 'CreateEvent',
    component: () => import('../views/CreateEvent.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/available-events',
    name: 'AvailableEvents',
    component: () => import('../views/AvailableEvents.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/hosted-events',
    name: 'HostedEvents',
    component: () => import('../views/HostedEvents.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/joined-events',
    name: 'JoinedEvents',
    component: () => import('../views/JoinedEvents.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/events/:id',
    name: 'EventDetails',
    component: () => import('../views/EventDetails.vue'),
    props: true,
    meta: { requiresAuth: true },
  },
  {
    path: '/edit-event/:id',
    name: 'EditEvent',
    component: () => import('../views/EditEvent.vue'),
    props: true,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const sessionStore = useSessionStore();
  if (!sessionStore.isAuthenticated && !sessionStore.user) {
    await sessionStore.initializeSession();
  }

  if (to.meta.requiresAuth && !sessionStore.isAuthenticated) {
    next('/login');
  } else {
    next();
  }
});

export default router;