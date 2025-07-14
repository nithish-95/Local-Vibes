<template>
  <div class="home container mx-auto px-4 py-8">
    <div class="hero bg-gradient-to-r from-primary-800 to-neutral-900 text-dark-text text-center py-20 shadow-xl rounded-lg mx-auto mt-8">
      <h1 v-if="sessionStore.isAuthenticated && sessionStore.user" class="text-5xl font-serif">Welcome, {{ sessionStore.user.username }}!</h1>
      <h1 v-else class="text-5xl font-serif">Find Your Vibe</h1>
      <p class="text-xl mt-4 font-sans">Discover and join local events happening near you.</p>
      <div class="mt-8">
        <div class="relative w-full max-w-md mx-auto">
          <input type="text" id="event-search" placeholder=" "
            class="peer w-full p-4 rounded-lg text-light-text focus:outline-none focus:ring-2 focus:ring-accent-500 font-sans border border-neutral-200 placeholder-shown:border-neutral-300 focus:border-accent-500">
          <label for="event-search"
            class="absolute left-4 -top-3.5 text-neutral-600 text-sm transition-all peer-placeholder-shown:text-base peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:top-4 peer-focus:-top-3.5 peer-focus:text-neutral-600 peer-focus:text-sm font-sans">Search for events...</label>
        </div>
      </div>
    </div>

    <div class="events-grid container mx-auto mt-8">
      <h2 class="text-3xl font-serif mb-4 text-light-text">Upcoming Events</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <!-- Event Card -->
        <div v-for="event in events" :key="event.id" class="relative flex flex-col bg-white bg-clip-border rounded-xl shadow-xl transform transition-transform duration-300 hover:scale-105">
          <div class="relative mx-4 mt-4 overflow-hidden text-dark-text shadow-lg bg-clip-border rounded-xl bg-primary-500 shadow-primary-500/40">
            <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/400/200`" alt="Event Image" class="w-full h-48 object-cover" />
          </div>
          <div class="p-6">
            <h3 class="block mb-2 font-serif text-xl antialiased font-semibold leading-snug tracking-normal text-light-text truncate">{{ event.title }}</h3>
            <p class="block font-sans text-base antialiased font-light leading-relaxed text-neutral-700 mb-4 line-clamp-2 overflow-hidden">{{ event.description }}</p>
            <div class="flex justify-between items-center">
              <router-link :to="`/events/${event.id}`" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-accent-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none" type="button">View Event</router-link>
              <span class="text-neutral-600 text-sm font-sans">{{ event.date }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getEvents } from '../api/events';
import { useSessionStore } from '../stores/session';
import { mapStores } from 'pinia';

export default {
  name: 'Home',
  data() {
    return {
      events: [],
    };
  },
  computed: {
    ...mapStores(useSessionStore),
  },
  async created() {
    this.events = await getEvents();
  },
};
</script>

<style scoped>
/* Add any specific styles for the home page here */
</style>