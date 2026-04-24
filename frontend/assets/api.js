// API Configuration
const API_BASE_URL = 'http://localhost:8080/api';

// Helper function to get auth token
function getToken() {
    return localStorage.getItem('token');
}

// Helper function to make API calls
async function apiCall(endpoint, method = 'GET', data = null) {
    const options = {
        method,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${getToken()}`
        }
    };

    if (data && (method === 'POST' || method === 'PUT')) {
        options.body = JSON.stringify(data);
    }

    try {
        const response = await fetch(`${API_BASE_URL}${endpoint}`, options);
        
        if (!response.ok) {
            if (response.status === 401) {
                logout();
                return null;
            }
            throw new Error(`HTTP ${response.status}`);
        }

        return await response.json();
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}

// Show message function
function showMessage(elementId, message, type = 'error', duration = 5000) {
    const messageEl = document.getElementById(elementId);
    if (!messageEl) return;

    messageEl.textContent = message;
    messageEl.className = `message ${type}`;
    
    if (duration) {
        setTimeout(() => {
            messageEl.className = 'message';
        }, duration);
    }
}

// Logout function
function logout() {
    localStorage.removeItem('token');
    window.location.href = '/login';
}

// Check authentication
function checkAuth() {
    if (!getToken()) {
        window.location.href = '/login';
    }
}
