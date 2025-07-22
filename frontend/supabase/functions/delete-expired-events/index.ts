import { createClient } from 'https://esm.sh/@supabase/supabase-js@2'

const supabaseUrl = process.env.SUPABASE_URL
const supabaseKey = process.env.SUPABASE_ANON_KEY

const supabase = createClient(supabaseUrl, supabaseKey)

async function deleteExpiredEvents() {
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

addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  try {
    await deleteExpiredEvents()
    return new Response('Expired events deleted successfully.', { status: 200 })
  } catch (error) {
    return new Response(`Error: ${error.message}`, { status: 500 })
  }
}
