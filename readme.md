Go TinyURL Service

A simple URL shortening service built with Go.

Go TinyURL Service is a lightweight, fast, and easy-to-use URL shortening service written in Go. It provides a RESTful API for shortening long URLs into shorter, more manageable links. Users can also set expiration dates for the shortened URLs.


Features

URL Shortening: Shorten long URLs into tiny, manageable links.
Custom Expiry: Set expiration dates for shortened URLs.
RESTful API: Easy-to-use API for URL shortening and redirection.
Configurability: Customizable configurations for port, base URL, and default expiry time (in minutes).


Usage

Shorten a URL
To shorten a URL, send a POST request to the /shorten endpoint with the URL in the request body as JSON.


curl --location 'http://localhost:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.example.com/long-url",
    "expiry": 1
}'


Replace "https://www.example.com/long-url" with the URL you want to shorten. The expiry parameter is optional and denotes the number of days until the shortened URL expires.

Redirect to a Shortened URL
To redirect to a shortened URL, simply make a GET request to the shortened URL endpoint.


curl --location 'http://localhost:8080/shortened-url-code'
Replace http://localhost:8080/shortened-url-code with the shortened URL you want to redirect to.