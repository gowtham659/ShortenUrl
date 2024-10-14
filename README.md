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

<h2>How It Works:</h2>
<ol>
  <li>User submits a long URL via the frontend form.</li>
  <li>The backend API generates a unique short code, stores the mapping in the Oracle database, and returns the shortened URL.</li>
  <li>When a user visits the shortened URL, the backend retrieves the original URL from the database and redirects the user to the destination.</li>
</ol>

<h2>Setup Instructions:</h2>
<ol>
  <li>Clone the repository:</li>
  <pre>git clone https://github.com/your-repo/url-shortener-go-oracle.git</pre>
  <li>Install dependencies:</li>
  <pre>
    go get -u github.com/gorilla/mux
    go get github.com/godror/godror
</pre>
  <li>Set up the Oracle Database with the provided SQL schema.</li>
  <li>Update the Oracle DB connection string in database/db.go</li>
  <li>Run the application:</li>
  <pre>
    go run main.go</pre>
  <li>Open your browser and visit http://localhost:8080/ to use the service.</li>
</ol>
