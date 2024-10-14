<h1>URL Shortener Service with Go (Backend) and Oracle (Database)</h1>
<p>This project is a URL Shortener Service built with Golang for the backend and Oracle as the database. It provides a simple web interface for users to shorten long URLs and retrieve the original URL using a generated short code. The service is lightweight and scalable, making use of modern web development practices and efficient database operations.</p>
<h2>Features:</h2>
<ul>
  <li><b>Frontend</b>: A clean, responsive UI built with HTML, CSS, and JavaScript for user interaction.
    <ul>
      <li>Enter a long URL and receive a shortened URL.</li>
      <li>Click the shortened URL to redirect to the original URL.</li>
    </ul>
  </li>
  <li>Backend: A RESTful API built with Golang and Gorilla Mux for routing.
    <ul>
      <li>Generate random short codes for URLs.</li>
      <li>Store URL mappings in Oracle Database.</li>
      <li>Handle HTTP requests to shorten URLs and redirect short URLs to their original counterparts</li>
    </ul>
  </li>
  <li>Database: Oracle Database is used to persist the URL and short code mappings.
    <ul>
      <li>Efficient database queries using Oracle's SQL..</li>
      <li>Supports horizontal scaling with the help of Oracle and Go's concurrency features.</li>
    </ul>
  </li>
</ul>

<h2>Technology Stack:</h2>
<ul>
  <li>Frontend: HTML, CSS, JavaScript</li>
  <li>Backend: Golang, Gorilla Mux, Oracle Go Driver (goora)</li>
  <li>Database: Oracle Database</li>
</ul>
