<template>
  <nav class="bg-gray-800 text-white p-4 fixed top-0 left-0 w-full z-50">
    <div class="container mx-auto flex justify-between items-center">
      <router-link to="/" class="text-2xl font-bold" @click="isMobileMenuOpen = false">LocalVibes</router-link>

      <!-- Desktop Menu -->
      <div class="hidden md:flex space-x-4 items-center">
        <router-link to="/" class="hover:text-gray-300">Home</router-link>
        
        <!-- Events Dropdown -->
        <div class="relative group">
          <button class="hover:text-gray-300 focus:outline-none">Events ▼</button>
          <div class="absolute hidden group-hover:block bg-gray-700 text-white rounded-md shadow-lg py-2 z-10">
            <router-link to="/available-events" class="block px-4 py-2 hover:bg-gray-600">Available Events</router-link>
            <template v-if="sessionStore.isAuthenticated">
              <router-link to="/hosted-events" class="block px-4 py-2 hover:bg-gray-600">Hosted Events</router-link>
              <router-link to="/create-event" class="block px-4 py-2 hover:bg-gray-600">Create Event</router-link>
              <router-link to="/joined-events" class="block px-4 py-2 hover:bg-gray-600">Joined Events</router-link>
            </template>
          </div>
        </div>

        <template v-if="sessionStore.isAuthenticated">
          <!-- Profile Dropdown -->
          <div class="relative" ref="profileDropdown">
            <button @click="toggleProfileDropdown" class="flex items-center space-x-1 hover:text-gray-300 focus:outline-none">
              <span v-if="sessionStore.user">Welcome, {{ sessionStore.user.username }}!</span>
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
            </button>
            <div v-if="isProfileDropdownOpen" class="absolute right-0 mt-2 w-48 bg-gray-700 text-white rounded-md shadow-lg py-2 z-20">
              <router-link to="/profile/edit" class="block px-4 py-2 hover:bg-gray-600">Edit Profile</router-link>
              <router-link to="/profile/settings" class="block px-4 py-2 hover:bg-gray-600">Settings</router-link>
              <div class="border-t border-gray-600 my-1"></div>
              <button @click="logoutUser" class="block w-full text-left px-4 py-2 hover:bg-gray-600">Logout</button>
            </div>
          </div>
        </template>
        <template v-else>
          <router-link to="/login" class="hover:text-gray-300">Login</router-link>
          <router-link to="/register" class="hover:text-gray-300">Register</router-link>
        </template>
      </div>

      <!-- Mobile Menu Button (Hamburger) -->
      <div class="md:hidden flex items-center">
        <button @click="toggleMobileMenu" class="text-white focus:outline-none">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
        </button>
      </div>
    </div>

    <!-- Mobile Menu Content -->
    <div v-if="isMobileMenuOpen" class="md:hidden bg-gray-800 py-2">
      <router-link to="/" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Home</router-link>
      <router-link to="/available-events" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Available Events</router-link>
      <template v-if="sessionStore.isAuthenticated">
        <router-link to="/hosted-events" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Hosted Events</router-link>
        <router-link to="/create-event" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Create Event</router-link>
        <router-link to="/joined-events" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Joined Events</router-link>
        <div class="border-t border-gray-700 my-2"></div>
        <span v-if="sessionStore.user" class="block px-4 py-2 text-white">Welcome, {{ sessionStore.user.username }}!</span>
        <router-link to="/profile/edit" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Edit Profile</router-link>
        <router-link to="/profile/settings" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Settings</router-link>
        <button @click="logoutUser" class="block w-full text-left px-4 py-2 hover:bg-gray-700">Logout</button>
      </template>
      <template v-else>
        <router-link to="/login" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Login</router-link>
        <router-link to="/register" class="block px-4 py-2 hover:bg-gray-700" @click="isMobileMenuOpen = false">Register</router-link>
      </template>
    </div>
  </nav>
</template>

<script>
import { useSessionStore } from '../stores/session';
import { mapStores } from 'pinia';

export default {
  name: 'Navbar',
  data() {
    return {
      isMobileMenuOpen: false,
      isProfileDropdownOpen: false,
    };
  },
  computed: {
    ...mapStores(useSessionStore),
  },
  methods: {
    toggleMobileMenu() {
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
      this.isProfileDropdownOpen = false; // Close profile dropdown if mobile menu opens
    },
    toggleProfileDropdown() {
      this.isProfileDropdownOpen = !this.isProfileDropdownOpen;
    },
    closeProfileDropdown(event) {
      if (this.$refs.profileDropdown && !this.$refs.profileDropdown.contains(event.target)) {
        this.isProfileDropdownOpen = false;
      }
    },
    async logoutUser() {
      try {
        await this.sessionStore.logout();
        this.$router.push('/login');
        this.isMobileMenuOpen = false; // Close mobile menu on logout
        this.isProfileDropdownOpen = false; // Close profile dropdown on logout
      } catch (error) {
        console.error(error);
      }
    },
  },
  mounted() {
    document.addEventListener('click', this.closeProfileDropdown);
  },
  beforeUnmount() {
    document.removeEventListener('click', this.closeProfileDropdown);
  },
};
</script>

<style scoped>
/* Add any specific styles for the navbar here */
</style>