const API_BASE_URL = '/api';

export async function registerUser(username, password) {
  try {
    const response = await fetch(`${API_BASE_URL}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Registration failed');
    }

    return response.json();
  } catch (error) {
    console.error('Error registering user:', error);
    throw error;
  }
}

export async function loginUser(username, password) {
  try {
    const response = await fetch(`${API_BASE_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Login failed');
    }

    return response.json();
  } catch (error) {
    console.error('Error logging in user:', error);
    throw error;
  }
}

export async function logoutUser() {
  try {
    const response = await fetch(`${API_BASE_URL}/logout`, {
      method: 'POST',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Logout failed');
    }

    return {}; // No content expected
  } catch (error) {
    console.error('Error logging out user:', error);
    throw error;
  }
}

export async function getCurrentUser() {
  try {
    const response = await fetch(`${API_BASE_URL}/user`);
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch current user');
    }
    return response.json();
  } catch (error) {
    console.error('Error fetching current user:', error);
    throw error;
  }
}
