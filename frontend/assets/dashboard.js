// Check authentication
if (!localStorage.getItem('token')) {
    window.location.href = '/login';
}

// Page navigation
const navLinks = document.querySelectorAll('.nav-link');
const pages = document.querySelectorAll('.page');

navLinks.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();
        const pageId = link.getAttribute('href').substring(1);
        navigateToPage(pageId);
    });
});

function navigateToPage(pageId) {
    // Update active state
    navLinks.forEach(link => link.classList.remove('active'));
    pages.forEach(page => page.classList.remove('active'));

    document.querySelector(`[href="#${pageId}"]`).classList.add('active');
    document.getElementById(pageId).classList.add('active');

    // Load data if needed
    if (pageId === 'dashboard') {
        loadDashboard();
    }
}

// Logout handler
document.getElementById('logoutBtn').addEventListener('click', () => {
    if (confirm('Apakah Anda yakin ingin logout?')) {
        localStorage.removeItem('token');
        window.location.href = '/login';
    }
});

// Load user profile
async function loadProfile() {
    try {
        const response = await fetch('http://localhost:8080/api/profile', {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (!response.ok) throw new Error('Failed to load profile');

        const data = await response.json();
        const profileEl = document.getElementById('profileInfo');

        profileEl.innerHTML = `
            <div class="profile-info">
                <div class="info-item">
                    <span class="info-label">Nama:</span>
                    <span class="info-value">${data.name}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">Email:</span>
                    <span class="info-value">${data.email}</span>
                </div>
            </div>
        `;
    } catch (error) {
        console.error('Error loading profile:', error);
    }
}

// Load history
async function loadHistory() {
    try {
        const response = await fetch('http://localhost:8080/api/history', {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (!response.ok) throw new Error('Failed to load history');

        const result = await response.json();
        const historyEl = document.getElementById('historyList');

        if (!result.data || result.data.length === 0) {
            historyEl.innerHTML = '<p>Belum ada analisis. <a href="#analysis" class="nav-link" onclick="navigateToPage(\'analysis\')">Buat analisis sekarang</a></p>';
            return;
        }

        let html = '';
        result.data.forEach(analysis => {
            const date = new Date(analysis.created_at * 1000).toLocaleDateString('id-ID');
            html += `
                <div class="history-item" onclick="viewAnalysisDetail(${analysis.id})">
                    <div class="history-item-title">${analysis.skin_type} - ${analysis.sensitivity}</div>
                    <div class="history-item-date">Tanggal: ${date}</div>
                </div>
            `;
        });

        historyEl.innerHTML = html;
    } catch (error) {
        console.error('Error loading history:', error);
    }
}

// View analysis detail
async function viewAnalysisDetail(id) {
    try {
        const response = await fetch(`http://localhost:8080/api/history/${id}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (!response.ok) throw new Error('Failed to load analysis');

        const result = await response.json();
        const analysis = result.data;

        // Display result
        displayRoutineResult(analysis);
        navigateToPage('analysis');
        window.scrollTo(0, 0);
    } catch (error) {
        console.error('Error loading analysis:', error);
        alert('Gagal memuat detail analisis');
    }
}

// Load dashboard
async function loadDashboard() {
    await loadProfile();
    await loadHistory();
}

// Form submission
document.getElementById('analysisForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const skinType = document.getElementById('skinType').value;
    const sensitivity = document.getElementById('sensitivity').value;

    if (!skinType || !sensitivity) {
        showMessage('analysisMessage', 'Harap isi semua field yang diperlukan', 'error');
        return;
    }

    // Get selected issues
    const issueCheckboxes = document.querySelectorAll('input[name="skinIssues"]:checked');
    if (issueCheckboxes.length === 0) {
        showMessage('analysisMessage', 'Harap pilih minimal satu masalah kulit', 'error');
        return;
    }

    const skinIssues = Array.from(issueCheckboxes).map(cb => cb.value);

    try {
        showMessage('analysisMessage', 'Sedang menganalisis kulit Anda...', 'loading');

        const response = await fetch('http://localhost:8080/api/skin-analysis', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify({
                skin_type: skinType,
                skin_issues: skinIssues,
                sensitivity: sensitivity
            })
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Analisis gagal');
        }

        const data = await response.json();
        showMessage('analysisMessage', 'Analisis berhasil!', 'success', 2000);

        // Display result
        displayRoutineResult(data.data);

        // Reset form
        document.getElementById('analysisForm').reset();

        // Reload history
        loadHistory();
    } catch (error) {
        showMessage('analysisMessage', 'Terjadi kesalahan: ' + error.message, 'error');
    }
}
);

// Display routine result
function displayRoutineResult(analysis) {
    const morningEl = document.getElementById('morningRoutine');
    const nightEl = document.getElementById('nightRoutine');
    const resultContainer = document.getElementById('resultContainer');

    // Morning routine
    let morningHtml = '<div class="routine-section">';
    morningHtml += '<h4 style="color: var(--primary-color); margin-bottom: 10px;">Langkah-langkah:</h4>';
    morningHtml += '<ul class="routine-list">';
    analysis.morning_routine.steps.forEach(step => {
        morningHtml += `<li>${step}</li>`;
    });
    morningHtml += '</ul>';
    morningHtml += `<div class="routine-tip">💡 ${analysis.morning_routine.tips}</div>`;
    morningHtml += '</div>';

    // Night routine
    let nightHtml = '<div class="routine-section">';
    nightHtml += '<h4 style="color: var(--primary-color); margin-bottom: 10px;">Langkah-langkah:</h4>';
    nightHtml += '<ul class="routine-list">';
    analysis.night_routine.steps.forEach(step => {
        nightHtml += `<li>${step}</li>`;
    });
    nightHtml += '</ul>';
    nightHtml += `<div class="routine-tip">💡 ${analysis.night_routine.tips}</div>`;
    nightHtml += '</div>';

    morningEl.innerHTML = morningHtml;
    nightEl.innerHTML = nightHtml;
    resultContainer.style.display = 'block';
}

// Helper function to show message
function showMessage(elementId, message, type = 'error', duration = 5000) {
    const messageEl = document.getElementById(elementId);
    if (!messageEl) return;

    messageEl.textContent = message;
    messageEl.className = `message ${type}`;

    if (duration && type !== 'loading') {
        setTimeout(() => {
            messageEl.className = 'message';
        }, duration);
    }
}

// Load dashboard on page load
document.addEventListener('DOMContentLoaded', () => {
    loadDashboard();
});
