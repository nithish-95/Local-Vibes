const API_BASE_URL = '/api';

export async function createEvent(eventData) {
  try {
    const response = await fetch(`${API_BASE_URL}/events`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(eventData),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to create event');
    }

    return response.json();
  } catch (error) {
    console.error('Error creating event:', error);
    throw error;
  }
}

export async function getEvents() {
  try {
    const response = await fetch(`${API_BASE_URL}/events`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch events');
    }
    const data = await response.json();
    return data || [];
  } catch (error) {
    console.error('Error fetching events:', error);
    throw error;
  }
}

export async function updateEvent(eventID, eventData) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(eventData),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to update event');
    }

    return response.json();
  } catch (error) {
    console.error('Error updating event:', error);
    throw error;
  }
}

export async function deleteEvent(eventID) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      // If the response is 204 No Content, it's still a success
      if (response.status === 204) {
        return {}; // Return an empty object or null to indicate success
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to delete event');
    }

    // Only try to parse JSON if there's content
    if (response.headers.get('content-length') > 0) {
      return response.json();
    } else {
      return {}; // No content to parse
    }
  } catch (error) {
    console.error('Error deleting event:', error);
    throw error;
  }
}

export async function getHostedEvents() {
  try {
    const response = await fetch(`${API_BASE_URL}/events/hosted`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch hosted events');
    }
    const data = await response.json();
    return data || [];
  } catch (error) {
    console.error('Error fetching hosted events:', error);
    throw error;
  }
}

export async function getAvailableEvents() {
  try {
    const response = await fetch(`${API_BASE_URL}/events/available`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch available events');
    }
    const data = await response.json();
    return data || [];
  } catch (error) {
    console.error('Error fetching available events:', error);
    throw error;
  }
}

export async function getEventByID(eventID) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch event by ID');
    }
    return response.json();
  } catch (error) {
    console.error('Error fetching event by ID:', error);
    throw error;
  }
}

export async function joinEvent(eventID) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}/join`, {
      method: 'POST',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to join event');
    }

    return {}; // No content expected
  } catch (error) {
    console.error('Error joining event:', error);
    throw error;
  }
}

export async function leaveEvent(eventID) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}/leave`, {
      method: 'POST',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to leave event');
    }

    return {}; // No content expected
  } catch (error) {
    console.error('Error leaving event:', error);
    throw error;
  }
}

export async function checkIfJoined(eventID) {
  try {
    const response = await fetch(`${API_BASE_URL}/events/${eventID}/is-joined`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to check join status');
    }
    const data = await response.json();
    return data.is_joined;
  } catch (error) {
    console.error('Error checking join status:', error);
    throw error;
  }
}
