<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center">Create New Event</h2>
      <form @submit.prevent="handleCreateEvent">
        <div class="mb-4">
          <label for="title" class="block text-gray-700 text-sm font-bold mb-2">Title</label>
          <input type="text" v-model="event.title" id="title" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="description" class="block text-gray-700 text-sm font-bold mb-2">Description</label>
          <textarea v-model="event.description" id="description" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"></textarea>
        </div>
        <div class="mb-4">
          <label for="date" class="block text-gray-700 text-sm font-bold mb-2">Date</label>
          <input type="date" v-model="event.date" id="date" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="time" class="block text-gray-700 text-sm font-bold mb-2">Time</label>
          <input type="time" v-model="event.time" id="time" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="location" class="block text-gray-700 text-sm font-bold mb-2">Location</label>
          <input type="text" v-model="event.location" id="location" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="rules" class="block text-gray-700 text-sm font-bold mb-2">Rules (comma-separated)</label>
          <input type="text" v-model="rulesInput" id="rules" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
        </div>
        <div class="mb-4">
          <label for="capacity" class="block text-gray-700 text-sm font-bold mb-2">Capacity</label>
          <input type="number" v-model.number="event.capacity" id="capacity" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required min="1">
        </div>
        <div class="flex items-center justify-between">
          <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
            Create Event
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { createEvent } from '../api/events';

export default {
  data() {
    return {
      event: {
        title: '',
        description: '',
        date: '',
        time: '',
        location: '',
        capacity: 1, // Default capacity
      },
      rulesInput: '', // For comma-separated input
    };
  },
  methods: {
    async handleCreateEvent() {
      try {
        const eventData = {
          ...this.event,
          rules: this.rulesInput.split(',').map(rule => rule.trim()).filter(rule => rule.length > 0),
        };
        await createEvent(eventData);
        alert('Event created successfully!');
        this.$router.push('/'); // Redirect to home or event list
      } catch (error) {
        alert(`Error creating event: ${error.message}`);
      }
    },
  },
};
</script>
