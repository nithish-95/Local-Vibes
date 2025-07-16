<template>
  <div class="home container mx-auto px-4 py-8">
    <div class="hero bg-gradient-to-r from-primary-800 to-neutral-900 text-dark-text text-center py-20 shadow-xl rounded-lg mx-auto mt-8">
      <h1 v-if="sessionStore.isAuthenticated && sessionStore.user" class="text-5xl font-serif">Welcome, {{ sessionStore.user.username }}!</h1>
      <h1 v-else class="text-5xl font-serif">Find Your Vibe</h1>
      <p class="text-xl mt-4 font-sans">Discover and join local events happening near you.</p>
      <div class="mt-8">
        <div class="relative w-full max-w-md mx-auto">
          <input type="text" id="event-search" placeholder=" " v-model="searchQuery" @keyup.enter="fetchEvents"
            class="peer w-full p-4 rounded-lg text-light-text focus:outline-none focus:ring-2 focus:ring-accent-500 font-sans border border-neutral-200 placeholder-shown:border-neutral-300 focus:border-accent-500">
          <label for="event-search"
            class="absolute left-4 -top-3.5 text-neutral-600 text-sm transition-all peer-placeholder-shown:text-base peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:top-4 peer-focus:-top-3.5 peer-focus:text-neutral-600 peer-focus:text-sm font-sans">Search for events...</label>
        </div>
      </div>
    </div>

    <div class="filters-section bg-white p-6 rounded-lg shadow-xl mb-8 mt-8">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-xl font-serif text-light-text">Filter Events</h3>
        <button @click="showFilters = !showFilters" class="text-neutral-600 hover:text-neutral-800 font-medium flex items-center">
          {{ showFilters ? 'Hide Filters' : 'Show Filters' }}
          <svg :class="{'rotate-180': showFilters}" class="w-5 h-5 ml-1 transform transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
        </button>
      </div>
      <transition name="fade">
        <div v-if="showFilters" class="space-y-4">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="startDate" class="block text-sm font-medium text-neutral-700">Start Date</label>
              <input type="date" id="startDate" v-model="startDate" class="mt-1 block w-full rounded-md border-neutral-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm">
            </div>
            <div>
              <label for="endDate" class="block text-sm font-medium text-neutral-700">End Date</label>
              <input type="date" id="endDate" v-model="endDate" class="mt-1 block w-full rounded-md border-neutral-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm">
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label for="capacityMin" class="block text-sm font-medium text-neutral-700">Min Capacity</label>
                <input type="number" id="capacityMin" v-model.number="capacityMin" class="mt-1 block w-full rounded-md border-neutral-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm">
              </div>
              <div>
                <label for="capacityMax" class="block text-sm font-medium text-neutral-700">Max Capacity</label>
                <input type="number" id="capacityMax" v-model.number="capacityMax" class="mt-1 block w-full rounded-md border-neutral-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm">
              </div>
            </div>
          </div>
          <div class="flex justify-end space-x-2 mt-4">
            <button @click="resetFilters" class="py-2 px-4 border border-neutral-300 rounded-md shadow-sm text-sm font-medium text-neutral-700 bg-white hover:bg-neutral-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500">
              Reset Filters
            </button>
            <button @click="fetchEvents" class="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-dark-text bg-accent-500 hover:bg-accent-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent-500">
              Apply Filters
            </button>
          </div>
        </div>
      </transition>
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
        <div v-if="events.length === 0" class="col-span-full text-center text-neutral-600 text-lg">No events found matching your criteria.</div>
      </div>
    </div>
  </div>
</template>

<script>
import { getEvents, searchEvents } from '../api/events';
import { useSessionStore } from '../stores/session';
import { mapStores } from 'pinia';

export default {
  name: 'Home',
  data() {
    return {
      events: [],
      searchQuery: '',
      startDate: '',
      endDate: '',
      capacityMin: null,
      capacityMax: null,
      showFilters: false, // New data property
    };
  },
  computed: {
    ...mapStores(useSessionStore),
  },
  methods: {
    async fetchEvents() {
      try {
        const filters = {
          q: this.searchQuery,
          startDate: this.startDate,
          endDate: this.endDate,
          capacityMin: this.capacityMin,
          capacityMax: this.capacityMax,
        };
        this.events = await searchEvents(filters);
      } catch (error) {
        console.error('Error fetching events:', error);
        this.events = [];
      }
    },
    resetFilters() {
      this.searchQuery = '';
      this.startDate = '';
      this.endDate = '';
      this.capacityMin = null;
      this.capacityMax = null;
      this.fetchEvents(); // Re-fetch events after resetting filters
    },
  },
  async created() {
    this.fetchEvents(); // Call fetchEvents on component creation
  },
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease-in-out;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>