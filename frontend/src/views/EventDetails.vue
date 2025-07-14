<template>
  <div v-if="event" class="max-w-4xl mx-auto p-8 bg-white rounded-lg shadow-xl mt-8">
    <button @click="$router.go(-1)" class="mb-4 text-blue-gray-600 hover:text-blue-gray-800 flex items-center font-medium transition-colors duration-300">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div>
        <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/600/400`" alt="Event Image" class="w-full h-auto rounded-lg shadow-lg">
      </div>
      <div>
        <h2 class="text-4xl font-serif text-blue-gray-900 mb-4">{{ event.title }}</h2>
        <p class="text-blue-gray-700 mb-6 font-sans">{{ event.description }}</p>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <span class="font-serif text-blue-gray-800 w-24">Date:</span>
            <span class="text-blue-gray-600 font-sans">{{ event.date }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-blue-gray-800 w-24">Time:</span>
            <span class="text-blue-gray-600 font-sans">{{ event.time }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-blue-gray-800 w-24">Location:</span>
            <span class="text-blue-gray-600 font-sans">{{ event.location }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-blue-gray-800 w-24">Capacity:</span>
            <span class="text-blue-gray-600 font-sans"> {{ event.participants}} / {{ event.capacity }}</span>
          </div>
           <div class="flex items-center">
            <span class="font-serif text-blue-gray-800 w-24">Host:</span>
            <span class="text-blue-gray-600 font-sans">{{ event.host_name }}</span>
          </div>
        </div>

        <div v-if="event.rules && event.rules.length > 0" class="mt-6">
          <h3 class="text-xl font-serif text-blue-gray-900 mb-2">Custom Rules:</h3>
          <ul class="list-disc list-inside text-blue-gray-700 font-sans">
            <li v-for="(rule, index) in event.rules" :key="index">{{ rule }}</li>
          </ul>
        </div>

        <div class="mt-8">
          <button v-if="!isJoined" @click="join" class="w-full py-3 px-4 font-semibold text-white bg-teal-500 rounded-lg hover:bg-teal-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-teal-500 shadow-md transition-colors duration-300 font-sans">
            Join Event
          </button>
          <button v-else @click="leave" class="w-full py-3 px-4 font-semibold text-white bg-red-500 rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 shadow-md transition-colors duration-300 font-sans">
            Leave Event
          </button>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="text-center mt-8">
    <p class="text-xl text-blue-gray-600 font-sans">Loading event details...</p>
  </div>
</template>

<script>
import { getEventByID, joinEvent, leaveEvent, checkIfJoined } from '../api/events';

export default {
  name: 'EventDetails',
  props: ['id'],
  data() {
    return {
      event: null,
      isJoined: false,
    };
  },
  async created() {
    try {
      this.event = await getEventByID(this.id);
      this.isJoined = await checkIfJoined(this.id);
    } catch (error) {
      console.error('Error fetching event details:', error);
    }
  },
  methods: {
    async join() {
      try {
        await joinEvent(this.id);
        alert('Successfully joined the event!');
        this.event = await getEventByID(this.id);
        this.isJoined = true;
      } catch (error) {
        console.error('Error joining event:', error);
        alert('Failed to join the event. You may have already joined or the event is full.');
      }
    },
    async leave() {
      try {
        await leaveEvent(this.id);
        alert('Successfully left the event!');
        this.event = await getEventByID(this.id);
        this.isJoined = false;
      } catch (error) {
        console.error('Error leaving event:', error);
        alert('Failed to leave the event.');
      }
    },
  },
};
</script>
