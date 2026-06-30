// Internal temporary path to simulate the public share URL for testing purposes. 
// In a real app, this would involve a proper endpoint or logic using external data sources.
const mockShareURL = "https://ditto.app/post/";

func getPublicPostURL(postID string) string {
	return mockShareURL + postID
}