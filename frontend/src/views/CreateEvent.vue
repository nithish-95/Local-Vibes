

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
        <span v-if="formErrors.title" class="text-red-500 text-sm">{{ formErrors.title }}</span>
      </div>
      <div>
        <label for="description" class="text-sm font-medium text-gray-700">Description</label>
        <textarea v-model="event.description" id="description" name="description" autocomplete="off" rows="4" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
        <span v-if="formErrors.description" class="text-red-500 text-sm">{{ formErrors.description }}</span>
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
      <span v-if="formErrors.dateTime" class="text-red-500 text-sm">{{ formErrors.dateTime }}</span>
      <div>
        <label for="location" class="text-sm font-medium text-gray-700">Location</label>
        <input type="text" v-model="event.location" id="location" name="location" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        <span v-if="formErrors.location" class="text-red-500 text-sm">{{ formErrors.location }}</span>
      </div>
      <div>
        <label for="capacity" class="text-sm font-medium text-gray-700">Capacity</label>
        <input type="number" v-model.number="event.capacity" id="capacity" name="capacity" autocomplete="off" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        <span v-if="formErrors.capacity" class="text-red-500 text-sm">{{ formErrors.capacity }}</span>
      </div>
      <div>
        <label for="image_url" class="text-sm font-medium text-gray-700">Event Image URL (Optional)</label>
        <input type="url" v-model="event.image_url" id="image_url" name="image_url" autocomplete="off" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="rules" class="text-sm font-medium text-gray-700">Custom Rules (one per line)</label>
        <textarea v-model="rulesInput" id="rules" name="rules" autocomplete="off" rows="4" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
        <span v-if="formErrors.rules" class="text-red-500 text-sm">{{ formErrors.rules }}</span>
      </div>
      <div>
        <button type="submit" :disabled="!isFormValid" class="w-full py-3 px-4 font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Create Event</button>
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
      formErrors: {
        title: '',
        description: '',
        location: '',
        capacity: '',
        rules: '',
        dateTime: '',
      },
    };
  },
  computed: {
    combinedDateTime() {
      if (this.event.date && this.event.time) {
        return `${this.event.date}T${this.event.time}`;
      }
      return '';
    },
    dateTimeError() {
      return this.formErrors.dateTime;
    },
    isFormValid() {
      return !this.formErrors.title &&
             !this.formErrors.description &&
             !this.formErrors.location &&
             !this.formErrors.capacity &&
             !this.formErrors.rules &&
             !this.formErrors.dateTime;
    },
  },
  watch: {
    'event.title'(newValue) {
      this.formErrors.title = newValue.trim() === '' ? 'Title cannot be empty.' : '';
    },
    'event.description'(newValue) {
      this.formErrors.description = newValue.trim() === '' ? 'Description cannot be empty.' : '';
    },
    'event.location'(newValue) {
      this.formErrors.location = newValue.trim() === '' ? 'Location cannot be empty.' : '';
    },
    'event.capacity'(newValue) {
      this.formErrors.capacity = (newValue === null || newValue === undefined || newValue < 2) ? 'Capacity must be at least 2.' : '';
    },
    rulesInput(newValue) {
      const trimmedRules = newValue.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
      this.formErrors.rules = trimmedRules.length === 0 ? 'At least one rule is required.' : '';
    },
    combinedDateTime(newValue) {
      if (!newValue) {
        this.formErrors.dateTime = '';
        return;
      }

      const eventStartDate = new Date(newValue);
      const now = new Date();
      const eightHoursFromNow = new Date(now.getTime() + 8 * 60 * 60 * 1000);

      if (isNaN(eventStartDate.getTime())) {
        this.formErrors.dateTime = 'Invalid date or time format.';
      } else if (eventStartDate < now) {
        this.formErrors.dateTime = 'Cannot create an event in the past.';
      } else if (eventStartDate < eightHoursFromNow) {
        this.formErrors.dateTime = 'Event must be scheduled at least 8 hours from now.';
      } else {
        this.formErrors.dateTime = '';
      }
    },
  },
  methods: {
    async create() {
      // Trigger all validations on submit attempt
      this.formErrors.title = this.event.title.trim() === '' ? 'Title cannot be empty.' : '';
      this.formErrors.description = this.event.description.trim() === '' ? 'Description cannot be empty.' : '';
      this.formErrors.location = this.event.location.trim() === '' ? 'Location cannot be empty.' : '';
      this.formErrors.capacity = (this.event.capacity === null || this.event.capacity === undefined || this.event.capacity < 2) ? 'Capacity must be at least 2.' : '';
      const trimmedRules = this.rulesInput.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
      this.formErrors.rules = trimmedRules.length === 0 ? 'At least one rule is required.' : '';
      // Re-run dateTime validation
      this.combinedDateTime; // Accessing computed property triggers its re-evaluation

      if (!this.isFormValid) {
        console.error('Form has validation errors. Please correct them.');
        return;
      }

      try {
        this.event.rules = trimmedRules;
        await createEvent(this.event);
        this.$router.push('/');
      } catch (error) {
        console.error(error);
        // Display backend error if any
        this.formErrors.dateTime = error.message; // Assuming backend errors are primarily date/time related
      }
    },
  },
};
</script>