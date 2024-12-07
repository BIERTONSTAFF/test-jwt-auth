# Implementation of JWT token authorization in Go
## Requirements
- Docker
- PostgreSQL
## Routes
### POST /me/:id/token
Request:
- `id` Parameter

201
```json
{
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM2NTI1OTAsImlwIjoiMTg4LjEzNC44OS4xNjMiLCJzdWIiOiJUQ1FVWm4yU3VHcXdRYzRlS1k1b2JGczZ0aG9neDFlc29YSGRmZHJTVkMwPSIsInVzZXJJZCI6IjFlMjlkNzMzLWMzMzQtNDM5OC05NTRjLTAyMmQzNzY5MzFkNiJ9.b2HwVufqSqkHnbx4erPYaOyAxzVWvkh_nl7-l-mlTH9Io5vfY24jito7N09eYfzWtaWrIkyr__yACDfkbne2Mg"
}
```
### POST /me/token/refresh
Request:
- `Authorization: Bearer TOKEN` Header
- `refreshToken` Cookie

200
```json
{
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM2NTI1OTcsImlwIjoiMTg4LjEzNC44OS4xNjMiLCJzdWIiOiJjTXF2L0VmcEFJY1IwYmc0WlR1NHo0aWJSVGd6VHFLWHJ2ZENxYUV1TU4wPSIsInVzZXJJZCI6IjFlMjlkNzMzLWMzMzQtNDM5OC05NTRjLTAyMmQzNzY5MzFkNiJ9.XZmcI8yh4mXoiZ_uccTIvbphilXAnSo3-sK3CZv0gEXxaw92CEBIldHe--qjlRrtKH4qr8WY423ELNCdU9XE9w"
}
```
## Tests
Run: `go test -v ./...`
## Docker
Image: https://hub.docker.com/r/biertonstaff/testjwtauth
### Run
```console
$ docker run --name testjwtauth \
-e DSN="host=127.0.0.1 user=postgres password=12345678 dbname=testjwtauth" \
-e SMTP_EMAIL="SMTPEmail" \
-e SMTP_PASSWORD="SMTPPassword" \
-e SMTP_HOST="SMTPHost" \
-e MOCK_SMTP_RECIPIENT="mockRecipient" \
-e JWT_SECRET="secret" \
-p 8080:8080 \
testjwtauth
```
## Example of IP changed warning
<img width="704" alt="Screenshot 2024-12-07 at 1 26 09 PM" src="https://github.com/user-attachments/assets/e0792f75-5a81-4b38-a3a7-29ab763d8c76">
