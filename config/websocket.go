package config

// WebsocketAllowedOrigins restricts which origins may open a websocket
// connection. Patterns are matched case-insensitively with path.Match (e.g.
// "app.example.com", "*.example.com"); the request host is always allowed.
//
// When empty (the default) all origins are accepted, i.e. origin verification
// is disabled via coder/websocket's InsecureSkipVerify. Set this to a list of
// host patterns to enforce cross-origin checks in production.
var WebsocketAllowedOrigins []string
