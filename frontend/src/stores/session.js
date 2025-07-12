
import { defineStore } from 'pinia';
import { getCurrentUser, loginUser, logoutUser as apiLogout } from '../api/auth';
import { supabase } from '../supabase';

export const useSessionStore = defineStore('session', {
  state: () => ({
    user: null,
    isAuthenticated: false,
  }),
  actions: {
    async initializeSession() {
      try {
        const { data: { user } } = await supabase.auth.getUser();
        if (user) {
          // Fetch profile to get the username
          const { data: profile, error: profileError } = await supabase
            .from('profiles')
            .select('username')
            .eq('id', user.id)
            .single();

          if (profileError) {
            console.error('Error fetching user profile:', profileError);
            this.clearSession();
            return;
          }

          this.user = { ...user, username: profile.username }; // Attach username to user object
          this.isAuthenticated = true;
        } else {
          this.clearSession();
        }
      } catch (error) {
        console.error('Error initializing session:', error);
        this.clearSession();
      }
    },
    async login(email, password) {
      const user = await loginUser(email, password);
      // After successful login, fetch the profile to get the username
      const { data: profile, error: profileError } = await supabase
        .from('profiles')
        .select('username')
        .eq('id', user.id)
        .single();

      if (profileError) {
        console.error('Error fetching user profile after login:', profileError);
        this.clearSession();
        throw profileError; // Re-throw to indicate login failure due to profile issue
      }

      this.user = { ...user, username: profile.username }; // Attach username to user object
      this.isAuthenticated = true;
      return this.user;
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
