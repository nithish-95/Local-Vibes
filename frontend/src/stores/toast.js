import { defineStore } from 'pinia';

export const useToastStore = defineStore('toast', {
  state: () => ({
    message: '',
    type: 'success', // 'success', 'error', 'info', 'warning'
    isVisible: false,
    timeoutId: null,
  }),
  actions: {
    showToast(message, type = 'success', duration = 3000) {
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
      }
      this.message = message;
      this.type = type;
      this.isVisible = true;

      this.timeoutId = setTimeout(() => {
        this.hideToast();
      }, duration);
    },
    hideToast() {
      this.isVisible = false;
      this.message = '';
      this.type = 'success';
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
        this.timeoutId = null;
      }
    },
  },
});