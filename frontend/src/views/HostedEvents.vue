<template>
  <div class="container mx-auto mt-8">
    <button @click="$router.go(-1)" class="mb-4 text-blue-gray-600 hover:text-blue-gray-800 flex items-center font-medium transition-colors duration-300">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-serif mb-4 text-blue-gray-800">My Hosted Events</h2>

    <div v-if="events.length === 0" class="text-center py-10 bg-white rounded-lg shadow-xl">
      <p class="text-xl text-blue-gray-700 mb-4 font-sans">Looks like you haven't hosted any events yet!</p>
      <p class="text-lg text-blue-gray-600 mb-6 font-sans">Ready to bring your community together? Create your first event now!</p>
      <router-link to="/create-event" class="bg-teal-500 hover:bg-teal-600 text-white font-bold py-3 px-6 rounded-lg text-lg transition-colors duration-300 shadow-md font-sans">
        Create Your First Event!
      </router-link>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div v-for="event in events" :key="event.id" class="bg-white rounded-lg shadow-xl overflow-hidden transform transition-transform duration-300 hover:scale-105">
        <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/400/200`" alt="Event Image" class="w-full h-48 object-cover">
        <div class="p-6">
          <h3 class="text-xl font-serif mb-2 text-blue-gray-800">{{ event.title }}</h3>
          <p class="text-blue-gray-700 mb-4 font-sans">{{ event.description }}</p>
          <div class="flex justify-between items-center">
            <router-link :to="`/edit-event/${event.id}`" class="bg-teal-500 text-white px-4 py-2 rounded-lg hover:bg-teal-600 transition-colors duration-300 shadow-md font-sans">Edit Event</router-link>
            <button @click="confirmDelete(event.id)" class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-colors duration-300 shadow-md font-sans">Delete</button>
            <span class="text-blue-gray-600 text-sm font-sans">{{ event.date }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getHostedEvents, deleteEvent } from '../api/events';

export default {
  name: 'HostedEvents',
  data() {
    return {
      events: [],
    };
  },
  async created() {
    this.events = await getHostedEvents();
  },
  methods: {
    async confirmDelete(eventID) {
      if (confirm('Are you sure you want to delete this event?')) {
        await this.deleteEvent(eventID);
      }
    },
    async deleteEvent(eventID) {
      try {
        await deleteEvent(eventID);
        this.events = this.events.filter(event => event.id !== eventID);
        alert('Event deleted successfully!');
      } catch (error) {
        console.error('Error deleting event:', error);
        alert('Failed to delete event.');
      }
    },
  },
};
</script>