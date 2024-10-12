document.getElementById('url-form').addEventListener('submit', async function (event) {
    event.preventDefault();

    const longUrl = document.getElementById('url-input').value;
    
    // Call the backend API to shorten the URL
    const response = await fetch('/shorten', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ original_url: longUrl })
    });

    if (response.ok) {
        const data = await response.json();
        const shortUrl = data.short_url;

        // Display the shortened URL
        const shortUrlContainer = document.getElementById('short-url-container');
        const shortUrlAnchor = document.getElementById('short-url');

        shortUrlAnchor.href = shortUrl;
        shortUrlAnchor.textContent = shortUrl;
        shortUrlContainer.classList.remove('hidden');
    } else {
        alert('Failed to shorten the URL. Please try again.');
    }
});
