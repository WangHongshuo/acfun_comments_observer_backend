package commentsob

type observeNextCommentsPage struct {
	aid       int64
	oldFloor  int64
	newFloor  int64
	nextPage  int
	totalPage int
	proxyAddr string
}

type observeNextArticle struct{}
