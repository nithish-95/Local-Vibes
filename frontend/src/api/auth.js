import { supabase } from '../supabase';

export async function registerUser(username, password) {
  const { data, error } = await supabase.auth.signUp({
    email: username, // Supabase uses email for username
    password: password,
  });

  if (error) {
    throw new Error(error.message);
  }
  return data.user;
}

export async function loginUser(username, password) {
  const { data, error } = await supabase.auth.signInWithPassword({
    email: username, // Supabase uses email for username
    password: password,
  });

  if (error) {
    throw new Error(error.message);
  }
  return data.user;
}

export async function logoutUser() {
  const { error } = await supabase.auth.signOut();
  if (error) {
    throw new Error(error.message);
  }
}

export async function getCurrentUser() {
  const { data: { user }, error } = await supabase.auth.getUser();
  if (error) {
    throw new Error(error.message);
  }
  return user;
}