import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';
import CreateEvent from '../views/CreateEvent.vue';
import AvailableEvents from '../views/AvailableEvents.vue';
import HostedEvents from '../views/HostedEvents.vue';
import JoinedEvents from '../views/JoinedEvents.vue';
import EventDetails from '../views/EventDetails.vue';
import EditEvent from '../views/EditEvent.vue';
import { getCurrentUser } from '../api/auth';
import { session } from '../session';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/create-event',
    name: 'CreateEvent',
    component: CreateEvent,
    meta: { requiresAuth: true },
  },
  {
    path: '/available-events',
    name: 'AvailableEvents',
    component: AvailableEvents,
    meta: { requiresAuth: true },
  },
  {
    path: '/hosted-events',
    name: 'HostedEvents',
    component: HostedEvents,
    meta: { requiresAuth: true },
  },
  {
    path: '/joined-events',
    name: 'JoinedEvents',
    component: JoinedEvents,
    meta: { requiresAuth: true },
  },
  {
    path: '/events/:id',
    name: 'EventDetails',
    component: EventDetails,
    props: true,
    meta: { requiresAuth: true },
  },
  {
    path: '/edit-event/:id',
    name: 'EditEvent',
    component: EditEvent,
    props: true,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  if (!session.isAuthenticated && !session.user) {
    try {
      const user = await getCurrentUser();
      if (user && user.id) {
        session.user = user;
        session.isAuthenticated = true;
      }
    } catch (error) {
      console.error("Error fetching current user in router guard:", error);
      session.user = null;
      session.isAuthenticated = false;
    }
  }

  if (to.meta.requiresAuth && !session.isAuthenticated) {
    next('/login');
  } else {
    next();
  }
});

export default router;