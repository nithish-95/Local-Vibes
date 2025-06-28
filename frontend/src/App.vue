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
import { session, initializeSession } from './session';
import { provide } from 'vue';

export default {
  name: 'App',
  components: {
    Navbar,
    Footer,
  },
  setup() {
    // Initialize session on app startup
    initializeSession();
    // Provide the reactive session object to all descendants
    provide('session', session);
  },
};
</script>

<style>
/* Add any global styles here */
</style>