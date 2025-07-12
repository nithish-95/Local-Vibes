
import { defineStore } from 'pinia';
import { getCurrentUser, loginUser, logoutUser as apiLogout } from '../api/auth';

export const useSessionStore = defineStore('session', {
  state: () => ({
    user: null,
    isAuthenticated: false,
  }),
  actions: {
    async initializeSession() {
      try {
        const user = await getCurrentUser();
        if (user && user.id) {
          this.user = user;
          this.isAuthenticated = true;
        } else {
          this.clearSession();
        }
      } catch (error) {
        this.clearSession();
      }
    },
    async login(username, password) {
      const user = await loginUser(username, password);
      await this.initializeSession();
      return user;
    },
    async logout() {
      await apiLogout();
      this.clearSession();
    },
    clearSession() {
      this.user = null;
      this.isAuthenticated = false;
    },
  },
});
