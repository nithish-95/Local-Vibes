<template>
  <div class="container mx-auto mt-8">
    <button @click="$router.go(-1)" class="mb-4 text-gray-600 hover:text-gray-800 flex items-center">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-bold mb-4">Available Events</h2>
    <div class="mb-4">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search events..."
        class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      />
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
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
</template>

<script>
import { getAvailableEvents } from '../api/events';

export default {
  name: 'AvailableEvents',
  data() {
    return {
      events: [],
      searchQuery: '',
    };
  },
  watch: {
    searchQuery: 'fetchEvents',
  },
  async created() {
    this.fetchEvents();
  },
  methods: {
    async fetchEvents() {
      this.events = await getAvailableEvents(this.searchQuery);
    },
  },
};
</script>