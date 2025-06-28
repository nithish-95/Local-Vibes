import { reactive } from 'vue';
import { getCurrentUser } from './api/auth';

export const session = reactive({
  user: null,
  isAuthenticated: false,
});

export async function initializeSession() {
  try {
    const user = await getCurrentUser();
    if (user && user.id) {
      session.user = user;
      session.isAuthenticated = true;
    } else {
      session.user = null;
      session.isAuthenticated = false;
    }
  } catch (error) {
    session.user = null;
    session.isAuthenticated = false;
  }
}

export function clearUser() {
  session.user = null;
  session.isAuthenticated = false;
}