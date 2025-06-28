<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
      <button @click="$router.go(-1)" class="mb-4 text-gray-600 hover:text-gray-800 flex items-center">
        <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
        Back
      </button>
      <h2 class="text-3xl font-bold text-center text-gray-900">Login to Your Account</h2>
      <form @submit.prevent="login" class="space-y-6">
        <div>
          <label for="username" class="text-sm font-medium text-gray-700">Username</label>
          <input type="text" v-model="username" id="username" name="username" autocomplete="username" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label for="password" class="text-sm font-medium text-gray-700">Password</label>
          <input type="password" v-model="password" id="password" name="password" autocomplete="current-password" required class="w-full px-4 py-2 mt-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <button type="submit" class="w-full py-2 px-4 font-semibold text-white bg-blue-500 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Login</button>
        </div>
      </form>
      <p class="text-sm text-center text-gray-600">
        Don't have an account? 
        <router-link to="/register" class="font-medium text-blue-500 hover:underline">Register</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import { loginUser } from '../api/auth';
import { initializeSession } from '../session';

export default {
  name: 'Login',
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    async login() {
      try {
        await loginUser(this.username, this.password);
        await initializeSession(); // This will fetch the user and update the reactive session object
        this.$router.push('/');
      } catch (error) {
        console.error('Login failed:', error);
      }
    },
  },
};
</script>