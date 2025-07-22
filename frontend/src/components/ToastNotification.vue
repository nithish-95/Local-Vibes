<template>
  <transition name="slide-fade">
    <div
      v-if="toastStore.isVisible"
      :class="toastClass"
      class="fixed top-4 right-4 p-4 rounded-lg shadow-lg text-white z-50 flex items-center justify-between"
    >
      <span>{{ toastStore.message }}</span>
      <button @click="toastStore.hideToast()" class="ml-4 text-white hover:text-gray-200 focus:outline-none">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>
    </div>
  </transition>
</template>

<script>
import { useToastStore } from '../stores/toast';

export default {
  name: 'ToastNotification',
  setup() {
    const toastStore = useToastStore();
    return { toastStore };
  },
  computed: {
    toastClass() {
      switch (this.toastStore.type) {
        case 'success':
          return 'bg-green-500';
        case 'error':
          return 'bg-red-500';
        case 'info':
          return 'bg-blue-500';
        case 'warning':
          return 'bg-yellow-500';
        default:
          return 'bg-gray-700';
      }
    },
  },
};
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>