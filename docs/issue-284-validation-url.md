# Issue 284: Validation of --url
# Description: The command line tool needs to validate the provided URL before it gets written into config.toml. Currently, we are accepting any string for this field which can lead to bad configurations and break runtime logic later on.

# Plan: Add a pre-flight validation check that ensures the input is a valid URL format (e.g., regex match or proper DNS resolution test). This should be implemented in the `config` command handler function, right before saving the configuration data.
```
// (Placeholder for code changes to validate URL)
const isValidURL = (url) => {
    try {
        new URL(url); // Basic check
        return true;
    } catch (_) {
        return false;
    }
};

// Implementation needs to be injected here.
```