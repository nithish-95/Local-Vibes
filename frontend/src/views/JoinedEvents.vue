<template>
  <div class="container mx-auto px-4 py-8">
    <button @click="$router.go(-1)" class="mb-4 text-neutral-600 hover:text-neutral-800 flex items-center font-medium transition-colors duration-300 font-sans">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-serif mb-4 text-light-text">My Joined Events</h2>

    <div v-if="events.length === 0" class="text-center py-10 bg-white rounded-xl shadow-xl">
      <p class="text-xl text-neutral-700 mb-4 font-sans">You haven't joined any events yet!</p>
      <p class="text-lg text-neutral-600 mb-6 font-sans">Explore available events and find your next adventure.</p>
      <router-link to="/available-events" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-primary-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">
        Find Events to Join!
      </router-link>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div v-for="event in events" :key="event.id" class="relative flex flex-col bg-white bg-clip-border rounded-xl shadow-xl transform transition-transform duration-300 hover:scale-105">
        <div class="relative mx-4 mt-4 overflow-hidden text-dark-text shadow-lg bg-clip-border rounded-xl bg-primary-500 shadow-primary-500/40">
          <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/400/200`" alt="Event Image" class="w-full h-48 object-cover" />
        </div>
        <div class="p-6">
          <h3 class="block mb-2 font-serif text-xl antialiased font-semibold leading-snug tracking-normal text-light-text truncate">{{ event.title }}</h3>
          <p class="block font-sans text-base antialiased font-light leading-relaxed text-neutral-700 mb-4 line-clamp-2 overflow-hidden max-w-full">{{ event.description }}</p>
          <div class="flex justify-between items-center">
            <router-link :to="`/events/${event.id}`" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-accent-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">View Event</router-link>
            <span class="text-neutral-600 text-sm font-sans">{{ event.date }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getJoinedEvents } from '../api/events';

export default {
  name: 'JoinedEvents',
  data() {
    return {
      events: [],
    };
  },
  async created() {
    try {
      this.events = await getJoinedEvents();
    } catch (error) {
      console.error('Error fetching joined events:', error);
    }
  },
};
</script>