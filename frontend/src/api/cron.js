import { supabase } from './supabase';

export async function deleteExpiredEvents() {
  const twentyFourHoursAgo = new Date(Date.now() - 24 * 60 * 60 * 1000).toISOString();

  const { error } = await supabase
    .from('events')
    .delete()
    .lt('start_date', twentyFourHoursAgo);

  if (error) {
    console.error('Error deleting expired events:', error);
    throw new Error(error.message);
  }

  console.log('Expired events deleted successfully.');
}
