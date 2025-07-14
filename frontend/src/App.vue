<template>
  <div id="app" class="bg-light-bg min-h-screen flex flex-col font-sans">
    <Navbar />
    <main class="flex-grow container mx-auto pt-20 px-4 pb-4">
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