// Authentication page handler
document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');

    if (loginForm) {
        loginForm.addEventListener('submit', handleLogin);
    }

    if (registerForm) {
        registerForm.addEventListener('submit', handleRegister);
    }
});

async function handleLogin(e) {
    e.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    if (!email || !password) {
        showMessage('message', 'Harap isi semua field', 'error');
        return;
    }

    try {
        showMessage('message', 'Sedang login...', 'loading');

        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            })
        });

        const data = await response.json();

        if (!response.ok) {
            showMessage('message', data.error || 'Login gagal', 'error');
            return;
        }

        localStorage.setItem('token', data.token);
        showMessage('message', 'Login berhasil! Mengalihkan...', 'success', 2000);
        
        setTimeout(() => {
            window.location.href = '/dashboard';
        }, 2000);

    } catch (error) {
        showMessage('message', 'Terjadi kesalahan: ' + error.message, 'error');
    }
}

async function handleRegister(e) {
    e.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;

    if (!name || !email || !password || !confirmPassword) {
        showMessage('message', 'Harap isi semua field', 'error');
        return;
    }

    if (password !== confirmPassword) {
        showMessage('message', 'Password tidak cocok', 'error');
        return;
    }

    if (password.length < 6) {
        showMessage('message', 'Password minimal 6 karakter', 'error');
        return;
    }

    try {
        showMessage('message', 'Sedang mendaftar...', 'loading');

        const response = await fetch('/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name,
                email,
                password
            })
        });

        const data = await response.json();

        if (!response.ok) {
            showMessage('message', data.error || 'Pendaftaran gagal', 'error');
            return;
        }

        localStorage.setItem('token', data.token);
        showMessage('message', 'Pendaftaran berhasil! Mengalihkan...', 'success', 2000);
        
        setTimeout(() => {
            window.location.href = '/dashboard';
        }, 2000);

    } catch (error) {
        showMessage('message', 'Terjadi kesalahan: ' + error.message, 'error');
    }
}
