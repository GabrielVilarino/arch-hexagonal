package response

type NewsClientResponse struct {
	Status       string
	TotalResults int
	Articles     []ArticleResponse
}

type ArticleResponse struct {
	Source      ArticleSourceResponse
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponse struct {
	Id   string
	Name string
}
