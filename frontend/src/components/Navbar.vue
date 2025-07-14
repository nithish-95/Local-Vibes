<template>
  <nav class="block w-full px-4 py-2 bg-dark-bg text-dark-text shadow-md fixed top-0 left-0 z-50 lg:px-8 lg:py-3">
    <div class="container flex flex-wrap items-center justify-between mx-auto">
      <router-link to="/" class="mr-4 block cursor-pointer py-1.5 font-serif text-2xl text-dark-text" @click="isMobileMenuOpen = false">
        LocalVibes
      </router-link>

      <!-- Desktop Menu -->
      <div class="hidden lg:block">
        <ul class="flex flex-col gap-2 mt-2 mb-4 lg:mb-0 lg:mt-0 lg:flex-row lg:items-center lg:gap-6">
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/" class="flex items-center hover:text-neutral-200 transition-colors duration-300">Home</router-link>
          </li>
          
          <!-- Events Dropdown -->
          <li class="relative flex items-center p-1 text-sm font-sans text-dark-text" ref="eventsDropdown">
            <button @click="toggleEventsDropdown" class="flex items-center hover:text-neutral-200 transition-colors duration-300 focus:outline-none">Events <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
            </button>
            <div v-if="isEventsDropdownOpen" class="absolute right-0 mt-2 w-max bg-neutral-800 text-dark-text rounded-md shadow-lg py-2 z-20 top-full">
              <router-link to="/available-events" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans" @click="isEventsDropdownOpen = false">Available Events</router-link>
              <template v-if="sessionStore.isAuthenticated">
                <router-link to="/hosted-events" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans" @click="isEventsDropdownOpen = false">Hosted Events</router-link>
                <router-link to="/create-event" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans" @click="isEventsDropdownOpen = false">Create Event</router-link>
                <router-link to="/joined-events" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans" @click="isEventsDropdownOpen = false">Joined Events</router-link>
              </template>
            </div>
          </li>

          <template v-if="sessionStore.isAuthenticated">
            <!-- Profile Dropdown -->
            <li class="relative flex items-center p-1 text-sm font-sans text-dark-text" ref="profileDropdown">
              <button @click="toggleProfileDropdown" class="flex items-center space-x-1 hover:text-neutral-200 transition-colors duration-300 focus:outline-none">
                <span v-if="sessionStore.user">Welcome, {{ sessionStore.user.username }}!</span>
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
              </button>
              <div v-if="isProfileDropdownOpen" class="absolute right-0 mt-2 w-max bg-neutral-800 text-dark-text rounded-md shadow-lg py-2 z-20 top-full">
                <router-link to="/profile/edit" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans">Edit Profile</router-link>
                <router-link to="/profile/settings" class="block px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans">Settings</router-link>
                <div class="border-t border-neutral-700 my-1"></div>
                <button @click="logoutUser" class="block w-full text-left px-4 py-2 hover:bg-neutral-700 transition-colors duration-300 font-sans">Logout</button>
              </div>
            </li>
          </template>
          <template v-else>
            <li class="flex items-center p-1 text-sm font-sans text-dark-text">
              <router-link to="/login" class="flex items-center hover:text-neutral-200 transition-colors duration-300">Login</router-link>
            </li>
            <li class="flex items-center p-1 text-sm font-sans text-dark-text">
              <router-link to="/register" class="flex items-center hover:text-neutral-200 transition-colors duration-300">Register</router-link>
            </li>
          </template>
        </ul>
      </div>

      <!-- Mobile Menu Button (Hamburger) -->
      <button
        class="relative h-6 max-h-[40px] w-6 max-w-[40px] select-none rounded-lg text-center align-middle font-sans text-xs font-medium uppercase text-dark-text transition-all hover:bg-neutral-500/10 active:bg-neutral-500/30 disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none lg:hidden"
        type="button"
        @click="toggleMobileMenu"
      >
        <span class="absolute transform -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
        </span>
      </button>
    </div>

    <!-- Mobile Menu Content -->
    <div v-if="isMobileMenuOpen" class="block h-auto w-full basis-full overflow-hidden text-dark-text transition-all duration-300 ease-in-out lg:hidden">
      <ul class="flex flex-col gap-2 mt-2 mb-4 lg:mb-0 lg:mt-0">
        <li class="flex items-center p-1 text-sm font-sans text-dark-text">
          <router-link to="/" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Home</router-link>
        </li>
        <li class="flex items-center p-1 text-sm font-sans text-dark-text">
          <router-link to="/available-events" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Available Events</router-link>
        </li>
        <template v-if="sessionStore.isAuthenticated">
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/hosted-events" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Hosted Events</router-link>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/create-event" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Create Event</router-link>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/joined-events" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Joined Events</router-link>
          </li>
          <div class="border-t border-neutral-700 my-2"></div>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <span v-if="sessionStore.user" class="block px-4 py-2">Welcome, {{ sessionStore.user.username }}!</span>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/profile/edit" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Edit Profile</router-link>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/profile/settings" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Settings</router-link>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <button @click="logoutUser" class="block w-full text-left hover:bg-neutral-700 transition-colors duration-300">Logout</button>
          </li>
        </template>
        <template v-else>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/login" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Login</router-link>
          </li>
          <li class="flex items-center p-1 text-sm font-sans text-dark-text">
            <router-link to="/register" class="flex items-center hover:bg-neutral-700 transition-colors duration-300" @click="isMobileMenuOpen = false">Register</router-link>
          </li>
        </template>
      </ul>
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
      isEventsDropdownOpen: false,
    };
  },
  computed: {
    ...mapStores(useSessionStore),
  },
  methods: {
    toggleMobileMenu() {
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
      this.isProfileDropdownOpen = false; // Close profile dropdown if mobile menu opens
      this.isEventsDropdownOpen = false; // Close events dropdown if mobile menu opens
    },
    toggleProfileDropdown() {
      this.isProfileDropdownOpen = !this.isProfileDropdownOpen;
      this.isEventsDropdownOpen = false; // Close events dropdown if profile dropdown opens
    },
    toggleEventsDropdown() {
      this.isEventsDropdownOpen = !this.isEventsDropdownOpen;
      this.isProfileDropdownOpen = false; // Close profile dropdown if events dropdown opens
    },
    closeProfileDropdown(event) {
      if (this.$refs.profileDropdown && !this.$refs.profileDropdown.contains(event.target)) {
        this.isProfileDropdownOpen = false;
      }
    },
    closeEventsDropdown(event) {
      if (this.$refs.eventsDropdown && !this.$refs.eventsDropdown.contains(event.target)) {
        this.isEventsDropdownOpen = false;
      }
    },
    async logoutUser() {
      try {
        await this.sessionStore.logout();
        this.$router.push('/login');
        this.isMobileMenuOpen = false; // Close mobile menu on logout
        this.isProfileDropdownOpen = false; // Close profile dropdown on logout
        this.isEventsDropdownOpen = false; // Close events dropdown on logout
      } catch (error) {
        console.error(error);
      }
    },
  },
  mounted() {
    document.addEventListener('click', this.closeProfileDropdown);
    document.addEventListener('click', this.closeEventsDropdown);
  },
  beforeUnmount() {
    document.removeEventListener('click', this.closeProfileDropdown);
    document.removeEventListener('click', this.closeEventsDropdown);
  },
};
</script>

<style scoped>
/* Add any specific styles for the navbar here */
</style>
