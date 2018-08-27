package service

import "github.com/makki0205/kyotohack/model"

var articles []model.Article

func init() {
	StoreArticle(model.Article{
		Title:  "title",
		Body:   "makki no kizi",
		UserID: 1,
	})
	StoreArticle(model.Article{
		Title:  "title",
		Body:   "jin no kizi",
		UserID: 2,
	})
}
func StoreArticle(article model.Article) {
	articles = append(articles, article)
}

func GetArticle() []model.Article {
	return articles
}
