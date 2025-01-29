package threads

// Represents a thread
type Thread struct {
	Name  string
	Posts []Post
}

type Post struct {
	Name string
	Body string
}
