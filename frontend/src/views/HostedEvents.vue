<template>
  <div class="container mx-auto px-4 py-8">
    <button @click="$router.go(-1)" class="mb-4 text-neutral-600 hover:text-neutral-800 flex items-center font-medium transition-colors duration-300 font-sans">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-serif mb-4 text-light-text">My Hosted Events</h2>

    <div v-if="events.length === 0" class="text-center py-10 bg-white rounded-xl shadow-xl">
      <p class="text-xl text-neutral-700 mb-4 font-sans">Looks like you haven't hosted any events yet!</p>
      <p class="text-lg text-neutral-600 mb-6 font-sans">Ready to bring your community together? Create your first event now!</p>
      <router-link to="/create-event" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-primary-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">Create Your First Event!</router-link>
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
            <router-link :to="`/edit-event/${event.id}`" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-accent-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">Edit</router-link>
            <button @click="showConfirmDeleteModal(event.id)" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-red-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">Delete</button>
            <span class="text-neutral-600 text-sm font-sans">{{ event.date }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
  <ConfirmationModal
    :is-visible="isModalVisible"
    title="Delete Event"
    message="Are you sure you want to delete this event? This action cannot be undone."
    @confirm="handleConfirmDelete"
    @cancel="handleCancelDelete"
  />
</template>

<script>
import { getHostedEvents, deleteEvent } from '../api/events';
import { useToast } from 'vue-toastification';
import ConfirmationModal from '../components/ConfirmationModal.vue';

export default {
  name: 'HostedEvents',
  components: {
    ConfirmationModal,
  },
  data() {
    return {
      events: [],
      isModalVisible: false,
      eventToDelete: null,
    };
  },
  async created() {
    console.log('HostedEvents: created hook - fetching events');
    this.events = await getHostedEvents();
    console.log('Hosted events fetched:', this.events);
  },
  methods: {
    showConfirmDeleteModal(eventID) {
      this.eventToDelete = eventID;
      this.isModalVisible = true;
    },
    async handleConfirmDelete() {
      if (this.eventToDelete) {
        await this.deleteEvent(this.eventToDelete);
        this.eventToDelete = null;
      }
      this.isModalVisible = false;
    },
    handleCancelDelete() {
      this.eventToDelete = null;
      this.isModalVisible = false;
    },
    async deleteEvent(eventID) {
      try {
        await deleteEvent(eventID);
        this.events = this.events.filter(event => event.id !== eventID);
        const toast = useToast();
        toast.success('Event deleted successfully!');
      } catch (error) {
        console.error('Error deleting event:', error);
        const toast = useToast();
        toast.error('Failed to delete event.');
      }
    },
  },
};
</script>
