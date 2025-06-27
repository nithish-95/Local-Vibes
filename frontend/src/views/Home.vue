<template>
  <div class="min-h-screen bg-gray-100 p-4">
    <h1 v-if="loading" class="text-4xl font-bold text-gray-500 text-center mb-8">
      Loading...
    </h1>
    <div v-else>
      <h1 v-if="userName" class="text-4xl font-bold text-blue-600 text-center mb-8">
        Hey it's Working! 🔥 Welcome {{ userName }}
      </h1>
      <h1 v-else class="text-4xl font-bold text-red-600 text-center mb-8">
        Please login to continue.
      </h1>

      <div v-if="isLoggedIn">
        <h2 class="text-3xl font-bold text-gray-800 mb-6 text-center">My Hosted Events</h2>
        <div v-if="hostedEvents.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
          <div v-for="event in hostedEvents" :key="event.id" class="bg-white p-6 rounded-lg shadow-md">
            <h3 class="text-xl font-semibold text-gray-900 mb-2">{{ event.title }}</h3>
            <p class="text-gray-700 mb-2">{{ event.description }}</p>
            <p class="text-gray-600 text-sm"><strong>Date:</strong> {{ event.date }}</p>
            <p class="text-gray-600 text-sm"><strong>Time:</strong> {{ event.time }}</p>
            <p class="text-gray-600 text-sm"><strong>Location:</strong> {{ event.location }}</p>
            <p class="text-gray-600 text-sm"><strong>Participants:</strong> {{ event.participants }}/{{ event.capacity }}</p>
            <p class="text-gray-600 text-sm"><strong>Hosted by:</strong> {{ event.host_name }}</p>
            <div v-if="event.rules && event.rules.length > 0" class="mt-2">
              <p class="text-gray-600 text-sm font-semibold">Rules:</p>
              <ul class="list-disc list-inside text-gray-600 text-sm">
                <li v-for="rule in event.rules" :key="rule">{{ rule }}</li>
              </ul>
            </div>
            <div class="mt-4 flex justify-end space-x-2">
              <router-link :to="`/edit-event/${event.id}`" class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded">Edit</router-link>
              <button @click="confirmDelete(event.id)" class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Delete</button>
            </div>
          </div>
        </div>
        <p v-else class="text-center text-gray-600 mb-12">You haven't hosted any events yet. Create one!</p>

        <h2 class="text-3xl font-bold text-gray-800 mb-6 text-center">Events Available to Join</h2>
        <div v-if="availableEvents.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="event in availableEvents" :key="event.id" class="bg-white p-6 rounded-lg shadow-md">
            <h3 class="text-xl font-semibold text-gray-900 mb-2">{{ event.title }}</h3>
            <p class="text-gray-700 mb-2">{{ event.description }}</p>
            <p class="text-gray-600 text-sm"><strong>Date:</strong> {{ event.date }}</p>
            <p class="text-gray-600 text-sm"><strong>Time:</strong> {{ event.time }}</p>
            <p class="text-gray-600 text-sm"><strong>Location:</strong> {{ event.location }}</p>
            <p class="text-gray-600 text-sm"><strong>Participants:</strong> {{ event.participants }}/{{ event.capacity }}</p>
            <p class="text-gray-600 text-sm"><strong>Hosted by:</strong> {{ event.host_name }}</p>
            <div v-if="event.rules && event.rules.length > 0" class="mt-2">
              <p class="text-gray-600 text-sm font-semibold">Rules:</p>
              <ul class="list-disc list-inside text-gray-600 text-sm">
                <li v-for="rule in event.rules" :key="rule">{{ rule }}</li>
              </ul>
            </div>
            <div class="mt-4 flex justify-end">
              <button v-if="event.isJoined" @click="leaveEvent(event.id)" class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Leave Event</button>
              <button v-else @click="joinEvent(event.id)" class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded">Join Event</button>
            </div>
          </div>
        </div>
        <p v-else class="text-center text-gray-600">No events available to join at the moment.</p>
      </div>
    </div>
  </div>
</template>

<script>
import { getHostedEvents, getAvailableEvents, deleteEvent, joinEvent, leaveEvent, checkIfJoined } from '../api/events';

export default {
  data() {
    return {
      userName: null,
      loading: true,
      isLoggedIn: false,
      hostedEvents: [],
      availableEvents: [],
    };
  },
  async mounted() {
    await this.fetchData();
  },
  watch: {
    $route: 'fetchData',
  },
  methods: {
    async fetchData() {
      this.loading = true;
      console.log('fetchData: Starting...');
      try {
        console.log('fetchData: Fetching user data...');
        const userResponse = await fetch('/api/user');
        console.log('fetchData: User response status:', userResponse.status);
        if (userResponse.ok) {
          const userData = await userResponse.json();
          this.userName = userData.userName;
          this.isLoggedIn = true;
          console.log('fetchData: User logged in as:', this.userName);

          console.log('fetchData: Fetching hosted events...');
          this.hostedEvents = [...await getHostedEvents()];
          console.log('fetchData: Hosted events after reassignment:', this.hostedEvents);
          console.log('fetchData: Hosted events:', this.hostedEvents);

          console.log('fetchData: Fetching available events...');
          let availableEventsData = await getAvailableEvents();
          for (let i = 0; i < availableEventsData.length; i++) {
            availableEventsData[i].isJoined = await checkIfJoined(availableEventsData[i].id);
          }
          this.availableEvents = [...availableEventsData];
          console.log('fetchData: Available events after reassignment:', this.availableEvents);
          console.log('fetchData: Available events:', this.availableEvents);

        } else {
          console.log('fetchData: User not logged in.');
          this.userName = null;
          this.isLoggedIn = false;
          this.hostedEvents = [];
          this.availableEvents = [];
        }
      } catch (error) {
        console.error('fetchData: Error fetching data:', error);
        this.userName = null;
        this.isLoggedIn = false;
        this.hostedEvents = [];
        this.availableEvents = [];
      } finally {
        this.loading = false;
        console.log('fetchData: Loading set to false.');
      }
    },
    async confirmDelete(eventID) {
      if (confirm('Are you sure you want to delete this event?')) {
        try {
          await deleteEvent(eventID);
          alert('Event deleted successfully!');
          this.fetchData(); // Refresh the event list
        } catch (error) {
          alert(`Error deleting event: ${error.message}`);
        }
      }
    },
    async joinEvent(eventID) {
      try {
        await joinEvent(eventID);
        alert('Successfully joined event!');
        this.fetchData();
      } catch (error) {
        alert(`Error joining event: ${error.message}`);
      }
    },
    async leaveEvent(eventID) {
      try {
        await leaveEvent(eventID);
        alert('Successfully left event!');
        this.fetchData();
      } catch (error) {
        alert(`Error leaving event: ${error.message}`);
      }
    },
  },
};
</script>
