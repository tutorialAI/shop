### So how are signed tokens used in authentication? Here’s a simplified outline of the process:

1. A user signs into their account on an authentication server
2. The authentication server returns a signed token with their account information or an ID (or both)
3. The signed token is stored in the browser’s localStorage or sessionStorage or anywhere the website prefers to store it
4. The signed token is retrieved and used anytime a part of the website needs authenticated access

[info]("https://dev.to/behalf/authentication-authorization-in-microservices-architecture-part-i-2cn0")