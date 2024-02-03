package commentsob

type observeNextCommentsPage struct {
	isNewAid  bool
	aid       int64
	oldFloor  int64
	nextPage  int
	proxyAddr string
}

type observeNextArticle struct{}
