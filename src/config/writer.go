func validateAndWriteConfig(url string) error {
    if !isValidURL(url) {
        return errors.New("invalid URL format provided for config file")
    }
    // actual write logic here...
    return nil
}
