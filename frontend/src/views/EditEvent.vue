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
        <label for="image_url" class="text-sm font-medium text-gray-700">Event Image URL (Optional)</label>
        <input type="url" v-model="event.image_url" id="image_url" name="image_url" autocomplete="off" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label class="text-sm font-medium text-gray-700 mb-2 block">Tags</label>
        <div class="flex flex-wrap gap-2">
          <div v-for="tag in predefinedTags" :key="tag" class="flex items-center">
            <input type="checkbox" :id="`tag-${tag}`" :value="tag" v-model="selectedTags" :disabled="selectedTags.length >= 5 && !selectedTags.includes(tag)" class="hidden">
            <label :for="`tag-${tag}`" class="px-3 py-1 rounded-full border cursor-pointer text-sm transition-colors duration-200"
              :class="{
                'bg-blue-500 text-white border-blue-500': selectedTags.includes(tag),
                'bg-gray-200 text-gray-700 border-gray-300 hover:bg-gray-300': !selectedTags.includes(tag),
                'opacity-50 cursor-not-allowed': selectedTags.length >= 5 && !selectedTags.includes(tag)
              }"
            >{{ tag }}</label>
          </div>
        </div>
        <span v-if="formErrors.tags" class="text-red-500 text-sm">{{ formErrors.tags }}</span>
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
      predefinedTags: [
        'Music', 'Sports', 'Food', 'Art', 'Tech', 'Outdoor', 'Community', 'Education', 'Health', 'Family',
      ],
      selectedTags: [],
      formErrors: {
        tags: '',
      },
    };
  },
  async created() {
    try {
      this.event = await getEventByID(this.id);
      if (this.event.rules) {
        this.rulesInput = this.event.rules.join('\n');
      }
      if (this.event.tags) {
        this.selectedTags = this.event.tags;
      }
    } catch (error) {
      console.error('Error fetching event details:', error);
    }
  },
  methods: {
    async update() {
      // Add validation for tags
      if (this.selectedTags.length > 5) {
        this.formErrors.tags = 'You can select a maximum of 5 tags.';
        return;
      } else {
        this.formErrors.tags = '';
      }

      try {
        this.event.rules = this.rulesInput.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
        this.event.tags = this.selectedTags;
        await updateEvent(this.id, this.event);
        alert('Event updated successfully!');
        this.$router.push('/');
      } catch (error) {
        console.error(error);
        alert('Failed to update event.');
      }
    },
  },
  watch: {
    'selectedTags'(newValue) {
      if (newValue.length > 5) {
        this.formErrors.tags = 'You can select a maximum of 5 tags.';
      } else {
        this.formErrors.tags = '';
      }
    },
  },
};
</script>