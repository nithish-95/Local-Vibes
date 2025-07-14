<template>
  <div v-if="event" class="max-w-4xl mx-auto p-8 bg-white rounded-lg shadow-xl mt-8">
    <button @click="$router.go(-1)" class="mb-4 text-neutral-600 hover:text-neutral-800 flex items-center font-medium transition-colors duration-300 font-sans">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div>
        <img :src="event.image_url || `https://picsum.photos/seed/${event.id}/600/400`" alt="Event Image" class="w-full h-auto rounded-lg shadow-lg">
      </div>
      <div>
        <h2 class="text-4xl font-serif text-light-text mb-4">{{ event.title }}</h2>
        <p class="text-neutral-700 mb-6 font-sans break-words overflow-hidden w-full max-w-full">{{ event.description }}</p>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <span class="font-serif text-neutral-800 w-24">Date:</span>
            <span class="text-neutral-600 font-sans">{{ event.date }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-neutral-800 w-24">Time:</span>
            <span class="text-neutral-600 font-sans">{{ event.time }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-neutral-800 w-24">Location:</span>
            <span class="text-neutral-600 font-sans">{{ event.location }}</span>
          </div>
          <div class="flex items-center">
            <span class="font-serif text-neutral-800 w-24">Capacity:</span>
            <span class="text-neutral-600 font-sans"> {{ Math.max(0, event.participants - 1)}} / {{ event.capacity }}</span>
          </div>
           <div class="flex items-center">
            <span class="font-serif text-neutral-800 w-24">Host:</span>
            <span class="text-neutral-600 font-sans">{{ event.host_name }}</span>
          </div>
        </div>

        <div v-if="event.rules && event.rules.length > 0" class="mt-6 border border-neutral-200 rounded-lg p-4">
          <h3 class="text-xl font-serif text-light-text mb-3">Custom Rules:</h3>
          <ul class="list-disc list-inside text-neutral-700 font-sans space-y-1">
            <li v-for="(rule, index) in event.rules" :key="index" class="break-words overflow-hidden">{{ rule }}</li>
          </ul>
        </div>

        <div class="mt-8">
          <button v-if="!isJoined" @click="join" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-primary-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none w-full">Join Event</button>
          <button v-else @click="leave" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-red-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none w-full">Leave Event</button>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="text-center mt-8">
    <p class="text-xl text-neutral-600 font-sans">Loading event details...</p>
  </div>
</template>

<script>
import { getEventByID, joinEvent, leaveEvent, checkIfJoined } from '../api/events';
import { useToast } from 'vue-toastification';

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
        const toast = useToast();
        toast.success('Successfully joined the event!');
        this.event = await getEventByID(this.id);
        this.isJoined = true;
      } catch (error) {
        console.error('Error joining event:', error);
        const toast = useToast();
        toast.error('Failed to join the event. You may have already joined or the event is full.');
      }
    },
    async leave() {
      try {
        await leaveEvent(this.id);
        const toast = useToast();
        toast.success('Successfully left the event!');
        this.event = await getEventByID(this.id);
        this.isJoined = false;
      } catch (error) {
        console.error('Error leaving event:', error);
        const toast = useToast();
        toast.error('Failed to leave the event.');
      }
    },
  },
};
</script>