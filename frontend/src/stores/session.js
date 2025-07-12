
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
        if (user) {
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
      // Supabase's signInWithPassword automatically sets the session, so we just need to update our store's state
      this.user = user;
      this.isAuthenticated = true;
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
