package internal

type Post struct {
	Title string
	Author string
}

type Result struct {
	Name string
	HotPosts []Post
	TopPosters []TopUsers
}

type TopUsers struct {
	Author   string
	NumberOfPosts int
}