<template>
  <div class="home">
    <div class="hero bg-gray-900 text-white text-center py-20">
      <h1 v-if="sessionStore.isAuthenticated && sessionStore.user" class="text-5xl font-bold">Welcome, {{ sessionStore.user.username }}!</h1>
      <h1 v-else class="text-5xl font-bold">Find Your Vibe</h1>
      <p class="text-xl mt-4">Discover and join local events happening near you.</p>
      <div class="mt-8 flex justify-center">
        <div class="w-full max-w-md">
          <input type="text" id="event-search" v-model="searchQuery" placeholder="Search for events..." class="w-full p-4 rounded-full text-gray-900">
        </div>
      </div>
    </div>

    <div class="events-grid container mx-auto mt-8">
      <h2 class="text-3xl font-bold mb-4">Upcoming Events</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <!-- Event Card -->
        <div v-for="event in events" :key="event.id" class="bg-white rounded-lg shadow-lg overflow-hidden">
          <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/400/200`" alt="Event Image" class="w-full h-48 object-cover">
          <div class="p-6">
            <h3 class="text-xl font-bold mb-2">{{ event.title }}</h3>
            <p class="text-gray-700 mb-4">{{ event.description }}</p>
            <div v-if="event.tags && event.tags.length" class="flex flex-wrap gap-2 mb-4">
              <span v-for="tag in event.tags" :key="tag" class="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full">{{ tag }}</span>
            </div>
            <div class="flex justify-between items-center">
              <router-link :to="`/events/${event.id}`" class="bg-blue-500 text-white px-4 py-2 rounded-full hover:bg-blue-600">View Event</router-link>
              <span class="text-gray-600">{{ event.date }}</span>
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
      searchQuery: '',
    };
  },
  computed: {
    ...mapStores(useSessionStore),
  },
  watch: {
    searchQuery: 'fetchEvents',
  },
  async created() {
    this.fetchEvents();
  },
  methods: {
    async fetchEvents() {
      this.events = await getEvents(this.searchQuery);
    },
  },
};
</script>

<style scoped>
.hero {
  background-image: url('https://picsum.photos/seed/hero/1200/400');
  background-size: cover;
  background-position: center;
}
</style>