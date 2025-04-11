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

func (r *Result) Equals(o *Result) bool {
	if r.Name != o.Name {
		return false
	}
	if len(r.HotPosts) != len(o.HotPosts){
        return false
    }
	for i, hobby := range r.HotPosts {
        if hobby != o.HotPosts[i]{
            return false
        }
    }
	if len(r.TopPosters) != len(o.TopPosters){
        return false
    }
    for i, hobby := range r.TopPosters {
        if hobby != o.TopPosters[i]{
            return false
        }
    }
	return true
}

type TopUsers struct {
	Author   string
	NumberOfPosts int
}