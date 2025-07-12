<template>
  <div class="max-w-2xl mx-auto p-8 bg-white rounded-lg shadow-md">
    <button @click="$router.go(-1)" class="mb-4 text-gray-600 hover:text-gray-800 flex items-center">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Edit Event</h2>
    <form @submit.prevent="update" class="space-y-6">
      <div>
        <label for="title" class="text-sm font-medium text-gray-700">Title</label>
        <input type="text" v-model="event.title" id="title" name="title" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="description" class="text-sm font-medium text-gray-700">Description</label>
        <textarea v-model="event.description" id="description" name="description" rows="4" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label for="date" class="text-sm font-medium text-gray-700">Date</label>
          <input type="date" v-model="event.date" id="date" name="date" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label for="time" class="text-sm font-medium text-gray-700">Time</label>
          <input type="time" v-model="event.time" id="time" name="time" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      <div>
        <label for="location" class="text-sm font-medium text-gray-700">Location</label>
        <input type="text" v-model="event.location" id="location" name="location" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="capacity" class="text-sm font-medium text-gray-700">Capacity</label>
        <input type="number" v-model.number="event.capacity" id="capacity" name="capacity" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="image" class="text-sm font-medium text-gray-700">Event Image</label>
        <input type="file" @change="onFileChange" id="image" name="image" accept="image/*" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="rules" class="text-sm font-medium text-gray-700">Custom Rules (one per line)</label>
        <textarea v-model="rulesInput" id="rules" name="rules" rows="4" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
      </div>
      <div>
        <button type="submit" class="w-full py-3 px-4 font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Update Event</button>
      </div>
    </form>
  </div>
</template>

<script>
import { getEventByID, updateEvent } from '../api/events';

export default {
  name: 'EditEvent',
  props: ['id'],
  data() {
    return {
      event: {},
      rulesInput: '',
      selectedFile: null,
    };
  },
  async created() {
    try {
      this.event = await getEventByID(this.id);
      if (this.event.rules) {
        this.rulesInput = this.event.rules.join('\n');
      }
    } catch (error) {
      console.error('Error fetching event details:', error);
    }
  },
  methods: {
    onFileChange(e) {
      this.selectedFile = e.target.files[0];
    },
    async update() {
      try {
        this.event.rules = this.rulesInput.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
        
        const formData = new FormData();
        Object.keys(this.event).forEach(key => {
          if (key === 'rules') {
            formData.append(key, JSON.stringify(this.event[key]));
          } else {
            formData.append(key, this.event[key]);
          }
        });

        if (this.selectedFile) {
          formData.append('image', this.selectedFile);
        }

        await updateEvent(this.id, formData);
        alert('Event updated successfully!');
        this.$router.push('/');
      } catch (error) {
        console.error(error);
        alert('Failed to update event.');
      }
    },
  },
};
</script>