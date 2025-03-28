# URL Shortener
Code is being developed while reading the Let's Go book by Alex Edwards.
For personal interest, I changed the scope from a `Code Snippet Hosting` to a `URL Shortener` service.

## Scope

- Simple HTML UI - little to no CSS
- Accept any URL and output a shortened URL.
- URL redirection mapping stored in DB.

## Project Structure

Using the [Go Doc's Server Project](https://go.dev/doc/modules/layout#server-project) file structure.

- `internal` represents non-application specific and re-usable code. 
`internal` cannot be exported, meaning it's only accessible by this project. 
- `cmd/web` represents the Go code for the web-server.
- `ui` represents the ui assets.