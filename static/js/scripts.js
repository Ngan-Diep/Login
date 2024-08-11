document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();

    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    // Here you can perform client-side validation if needed

    // Example: Simulate login request to backend (replace with actual API call)
    if (username === "admin" && password === "password") {
        document.getElementById("message").textContent = "Login successful!";
        // Redirect to dashboard or perform other actions upon successful login
        // window.location.href = "/dashboard";
    } else {
        document.getElementById("message").textContent = "Invalid username or password.";
    }
});
