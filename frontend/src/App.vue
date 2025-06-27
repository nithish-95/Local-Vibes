<template>
  <div id="app">
    <nav class="bg-white shadow-md">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-xl font-bold text-blue-600">Local-Vibes</router-link>
            </div>
          </div>
          <div class="flex items-center">
            <template v-if="isLoggedIn">
              <router-link to="/create-event" class="px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Create Event</router-link>
              <span class="px-3 py-2 text-sm font-medium text-gray-700">Welcome, {{ userName }}</span>
              <button @click="logout" class="ml-4 px-3 py-2 rounded-md text-sm font-medium text-white bg-red-600 hover:bg-red-700">Logout</button>
            </template>
            <template v-else>
              <router-link to="/login" class="px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50">Login</router-link>
              <router-link to="/register" class="ml-4 px-3 py-2 rounded-md text-sm font-medium text-white bg-blue-600 hover:bg-blue-700">Register</router-link>
            </template>
          </div>
        </div>
      </div>
    </nav>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      userName: null,
      isLoggedIn: false,
    };
  },
  
  watch: {
    $route: 'checkAuth',
  },
  methods: {
    async checkAuth() {
      try {
        const response = await fetch('/api/user');
        if (response.ok) {
          const data = await response.json();
          this.userName = data.userName;
          this.isLoggedIn = true;
        } else {
          this.userName = null;
          this.isLoggedIn = false;
        }
      } catch (error) {
        console.error('Error checking auth status:', error);
        this.userName = null;
        this.isLoggedIn = false;
      }
    },
    async logout() {
      try {
        const response = await fetch('/api/logout', {
          method: 'POST',
        });
        if (response.ok) {
          this.userName = null;
          this.isLoggedIn = false;
          this.$router.push('/login');
        } else {
          console.error('Logout failed');
        }
      } catch (error) {
        console.error('Error during logout:', error);
      }
    },
  },
};
</script>
