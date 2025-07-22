import { supabase } from '../supabase';

export async function createEvent(eventData) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const combinedStartDateTime = `${eventData.date}T${eventData.time}`;
  const eventStartDate = new Date(combinedStartDateTime);
  const now = new Date();
  const eightHoursFromNow = new Date(now.getTime() + 8 * 60 * 60 * 1000);

  console.log('eventData.date:', eventData.date);
  console.log('eventData.time:', eventData.time);
  console.log('combinedStartDateTime:', combinedStartDateTime);
  console.log('eventStartDate:', eventStartDate);
  console.log('now:', now);
  console.log('eightHoursFromNow:', eightHoursFromNow);

  if (isNaN(eventStartDate.getTime())) {
    throw new Error('Invalid event date or time provided.');
  }

  if (eventStartDate < now) {
    throw new Error('Cannot create an event in the past.');
  }

  if (eventStartDate < eightHoursFromNow) {
    throw new Error('Event must be scheduled at least 8 hours from now.');
  }

  const eventToInsert = {
    title: eventData.title,
    description: eventData.description,
    location: eventData.location,
    capacity: eventData.capacity,
    image_url: eventData.image_url,
    date: eventData.date,
    time: eventData.time,
    rules: eventData.rules || [],
    creator_id: user.id,
  };

  console.log('User ID for event creation:', user.id);
  console.log('Event data to insert:', eventToInsert);

  const { data, error } = await supabase
    .from('events')
    .insert([eventToInsert])
    .select();

  if (error) {
    console.error('Supabase createEvent error:', error);
    throw new Error(error.message);
  }

  // Automatically add the creator as a participant
  const { error: participantError } = await supabase
    .from('event_participants')
    .insert([{ event_id: data[0].id, user_id: user.id }]);

  if (participantError) {
    console.warn('Supabase auto-add participant warning:', participantError);
    // Log the error but don't fail the event creation if adding participant fails
  }

  return data[0];
}

export async function getEvents() {
  const { data, error } = await supabase
    .from('event_with_participant_count')
    .select(`
      *,
      creator_username
    `);

  if (error) {
    console.error('Supabase getEvents error:', error);
    throw new Error(error.message);
  }

  // Map data to match existing frontend structure (e.g., host_name)
  return data.map(event => ({
    ...event,
    host_name: event.creator_username || 'Unknown',
    participants: event.participants_count, // Use the count from the view
  }));
}

export async function updateEvent(eventID, eventData) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const { rules, ...rest } = eventData;
  const eventToUpdate = {
    ...rest,
    rules: rules || [],
  };

  const { error } = await supabase
    .from('events')
    .update(eventToUpdate)
    .eq('id', eventID); // Only filter by ID, RLS handles creator_id check

  if (error) {
    console.error('Supabase updateEvent error:', error);
    throw new Error(error.message);
  }
  return {}; // Return an empty object on success
}

export async function deleteEvent(eventID) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const { error } = await supabase
    .from('events')
    .delete()
    .eq('id', eventID)
    .eq('creator_id', user.id); // Ensure only creator can delete

  if (error) {
    console.error('Supabase deleteEvent error:', error);
    throw new Error(error.message);
  }
  return {};
}

export async function getHostedEvents() {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const { data, error } = await supabase
    .from('event_with_participant_count')
    .select(`
      *,
      creator_username
    `)
    .eq('creator_id', user.id);

  if (error) {
    console.error('Supabase getHostedEvents error:', error);
    throw new Error(error.message);
  }

  return data.map(event => ({
    ...event,
    host_name: event.creator_username || 'Unknown',
    participants: event.participants_count, // Use the count from the view
  }));
}

export async function getAvailableEvents(searchQuery = '') {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  let query = supabase
    .from('event_with_participant_count')
    .select(`
      *,
      creator_username
    `)
    .neq('creator_id', user.id); // Not equal to current user's ID

  if (searchQuery) {
    query = query.or(`title.ilike.%${searchQuery}%,description.ilike.%${searchQuery}%`);
  }

  const { data, error } = await query;

  if (error) {
    console.error('Supabase getAvailableEvents error:', error);
    throw new Error(error.message);
  }

  return data.map(event => ({
    ...event,
    host_name: event.creator_username || 'Unknown',
    participants: event.participants_count, // Use the count from the view
  }));
}

export async function getEventByID(eventID) {
  const { data, error } = await supabase
    .from('event_with_participant_count')
    .select(`
      *,
      creator_username
    `)
    .eq('id', eventID)
    .single(); // Expecting a single record

  if (error) {
    console.error('Supabase getEventByID error:', error);
    throw new Error(error.message);
  }

  return {
    ...data,
    host_name: data.creator_username || 'Unknown',
    participants: data.participants_count, // Use the count from the view
  };
}

export async function joinEvent(eventID) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  // Check if already joined
  const { data: existingParticipant, error: checkError } = await supabase
    .from('event_participants')
    .select('*')
    .eq('event_id', eventID)
    .eq('user_id', user.id);

  if (checkError) {
    console.error('Supabase joinEvent check error:', checkError);
    throw new Error(checkError.message);
  }

  if (existingParticipant && existingParticipant.length > 0) {
    throw new Error('User is already a participant in this event');
  }

  // Check capacity (requires RLS or a function on Supabase)
  // For now, this is a basic check. A more robust solution would involve a Supabase Function or RLS.
  const { data: eventData, error: eventError } = await supabase
    .from('events')
    .select('capacity')
    .eq('id', eventID)
    .single();

  if (eventError) {
    console.error('Supabase joinEvent event data error:', eventError);
    throw new Error(eventError.message);
  }

  const { count, error: countError } = await supabase
    .from('event_participants')
    .select('*', { count: 'exact', head: true })
    .eq('event_id', eventID);

  if (countError) {
    console.error('Supabase joinEvent count error:', countError);
    throw new Error(countError.message);
  }

  if (eventData.capacity > 0 && count >= eventData.capacity) {
    throw new Error('Event is full');
  }

  const { error } = await supabase
    .from('event_participants')
    .insert([{ event_id: eventID, user_id: user.id }]);

  if (error) {
    console.error('Supabase joinEvent insert error:', error);
    throw new Error(error.message);
  }
  return {};
}

export async function leaveEvent(eventID) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const { error, count } = await supabase
    .from('event_participants')
    .delete()
    .eq('event_id', eventID)
    .eq('user_id', user.id)
    .select(); // Use select to get count of affected rows

  if (error) {
    console.error('Supabase leaveEvent error:', error);
    throw new Error(error.message);
  }

  if (count === 0) {
    throw new Error('User was not a participant in this event');
  }
  return {};
}

export async function checkIfJoined(eventID) {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    // If not authenticated, they can't be joined
    return false;
  }

  const { data, error } = await supabase
    .from('event_participants')
    .select('*')
    .eq('event_id', eventID)
    .eq('user_id', user.id);

  if (error) {
    console.error('Supabase checkIfJoined error:', error);
    throw new Error(error.message);
  }
  return data.length > 0;
}

export async function getJoinedEvents() {
  const { data: { user } } = await supabase.auth.getUser();
  if (!user) {
    throw new Error('User not authenticated.');
  }

  const { data, error } = await supabase
    .from('event_participants')
    .select(`
      event:event_with_participant_count(*,
        creator_username
      )
    `)
    .eq('user_id', user.id);

  if (error) {
    console.error('Supabase getJoinedEvents error:', error);
    throw new Error(error.message);
  }

  // Extract the event objects from the nested structure
  return data.map(item => ({
    ...item.event,
    host_name: item.event.creator_username || 'Unknown',
    participants: item.event.participants_count, // Use the count from the view
  }));
}