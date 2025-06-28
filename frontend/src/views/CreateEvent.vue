<template>
  <div class="max-w-2xl mx-auto p-8 bg-white rounded-lg shadow-md">
    <button @click="$router.go(-1)" class="mb-4 text-gray-600 hover:text-gray-800 flex items-center">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Create a New Event</h2>
    <form @submit.prevent="create" class="space-y-6">
      <div>
        <label for="title" class="text-sm font-medium text-gray-700">Title</label>
        <input type="text" v-model="event.title" id="title" name="title" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="description" class="text-sm font-medium text-gray-700">Description</label>
        <textarea v-model="event.description" id="description" name="description" autocomplete="off" rows="4" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label for="date" class="text-sm font-medium text-gray-700">Date</label>
          <input type="date" v-model="event.date" id="date" name="date" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label for="time" class="text-sm font-medium text-gray-700">Time</label>
          <input type="time" v-model="event.time" id="time" name="time" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      <div>
        <label for="location" class="text-sm font-medium text-gray-700">Location</label>
        <input type="text" v-model="event.location" id="location" name="location" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="capacity" class="text-sm font-medium text-gray-700">Capacity</label>
        <input type="number" v-model.number="event.capacity" id="capacity" name="capacity" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="rules" class="text-sm font-medium text-gray-700">Custom Rules (one per line)</label>
        <textarea v-model="rulesInput" id="rules" name="rules" autocomplete="off" rows="4" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
      </div>
      <div>
        <button type="submit" class="w-full py-3 px-4 font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Create Event</button>
      </div>
    </form>
  </div>
</template>

<script>
import { createEvent } from '../api/events';

export default {
  name: 'CreateEvent',
  data() {
    return {
      event: {
        title: '',
        description: '',
        date: '',
        time: '',
        location: '',
        capacity: 0,
        rules: [],
      },
      rulesInput: '',
    };
  },
  methods: {
    async create() {
      try {
        this.event.rules = this.rulesInput.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
        await createEvent(this.event);
        this.$router.push('/');
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>