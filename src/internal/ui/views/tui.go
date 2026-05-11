// Placeholder for fix: Check if config exists but baseURL is missing, and force onboarding flow.
// This function needs to be added or modified in internal/ui/view/tui.go
func checkAndInit(config *Config) {
    if config.ConfigExists && config.BaseURL == "" {
        return StateOnboarding
    }
    // ... rest of the logic
}