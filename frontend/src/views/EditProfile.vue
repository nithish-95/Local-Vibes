<template>
  <div class="max-w-2xl mx-auto p-8 bg-white rounded-lg shadow-md mt-8">
    <button @click="$router.go(-1)" class="mb-4 text-gray-600 hover:text-gray-800 flex items-center">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Edit Profile</h2>
    <form @submit.prevent="updateProfile" class="space-y-6">
      <div>
        <label for="username" class="text-sm font-medium text-gray-700">Username</label>
        <input type="text" v-model="profile.username" id="username" name="username" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <label for="avatar_url" class="text-sm font-medium text-gray-700">Avatar URL (Optional)</label>
        <input type="url" v-model="profile.avatar_url" id="avatar_url" name="avatar_url" class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
      </div>
      <div>
        <button type="submit" class="w-full py-3 px-4 font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Update Profile</button>
      </div>
    </form>
    <p v-if="message" class="mt-4 text-center text-sm" :class="messageType === 'success' ? 'text-green-600' : 'text-red-600'">{{ message }}</p>
  </div>
</template>

<script>
import { supabase } from '../supabase';
import { useSessionStore } from '../stores/session';

export default {
  name: 'EditProfile',
  data() {
    return {
      profile: {
        username: '',
        avatar_url: '',
      },
      message: '',
      messageType: '',
    };
  },
  async created() {
    await this.fetchProfile();
  },
  methods: {
    async fetchProfile() {
      const sessionStore = useSessionStore();
      const user = sessionStore.user;

      if (!user) {
        this.message = 'You must be logged in to edit your profile.';
        this.messageType = 'error';
        return;
      }

      const { data, error } = await supabase
        .from('profiles')
        .select('username, avatar_url')
        .eq('id', user.id)
        .single();

      if (error) {
        console.error('Error fetching profile:', error);
        this.message = 'Error fetching profile.';
        this.messageType = 'error';
      } else if (data) {
        this.profile.username = data.username;
        this.profile.avatar_url = data.avatar_url;
      }
    },
    async updateProfile() {
      const sessionStore = useSessionStore();
      const user = sessionStore.user;

      if (!user) {
        this.message = 'You must be logged in to update your profile.';
        this.messageType = 'error';
        return;
      }

      const { error } = await supabase
        .from('profiles')
        .update({
          username: this.profile.username,
          avatar_url: this.profile.avatar_url,
        })
        .eq('id', user.id);

      if (error) {
        console.error('Error updating profile:', error);
        this.message = 'Error updating profile: ' + error.message;
        this.messageType = 'error';
      } else {
        // Re-initialize the session to fetch the updated username
        await sessionStore.initializeSession();
        this.message = 'Profile updated successfully!';
        this.messageType = 'success';
      }
    },
  },
};
</script>