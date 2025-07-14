<template>
  <div id="app" class="bg-gray-100 min-h-screen flex flex-col">
    <Navbar />
    <main class="flex-grow container mx-auto p-4">
      <router-view />
    </main>
    <Footer />
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';
import { useSessionStore } from './stores/session';

export default {
  name: 'App',
  components: {
    Navbar,
    Footer,
  },
  setup() {
    const sessionStore = useSessionStore();
    // Initialize session on app startup
    sessionStore.initializeSession();
  },
};
</script>

<style>
/* Add any global styles here */
</style>