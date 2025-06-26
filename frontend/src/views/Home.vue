<template>
  <div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <h1 v-if="loading" class="text-4xl font-bold text-gray-500">
      Loading...
    </h1>
    <h1 v-else-if="userName" class="text-4xl font-bold text-blue-600">
      Hey it's Working! 🔥 Welcome {{ userName }}
    </h1>
    <h1 v-else class="text-4xl font-bold text-red-600">
      Please login to continue.
    </h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: null,
      loading: true,
    };
  },
  async mounted() {
    try {
      const response = await fetch('/api/user'); // This will fail for now
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      this.userName = data.userName;
    } catch (error) {
      console.error('Error fetching user:', error);
    } finally {
      this.loading = false;
    }
  }
};
</script>
