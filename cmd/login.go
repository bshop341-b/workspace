// cmd/login.go: Added validation for --url flag in loginCmd's RunE() or PreRunE().
// This ensures that any provided base URL must be a properly formed http(s) URI with a valid host before saving to config.toml.
// Original code path modified to include:
// 1. url.Parse(baseURL) check for nil error.
// 2. Check if parsed.Scheme is "http" or "https".
// 3. Check if parsed.Host is not empty.
/* ... rest of file */