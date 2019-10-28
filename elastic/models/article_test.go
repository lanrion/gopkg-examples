package models_test

import (
	models "github.com/lanrion/gopkg-examples/elastic/models"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	title := "这里是中国"
	content := "这是内容"
	models.CreateArticle(title, content)
	article := models.Article{}
	models.GetDB().Last(&article)
	if article.Title != title {
		t.FailNow()
	}
}

func TestQueryArticleWithTitle(t *testing.T)  {
	models.QueryArticleWithTitle("中国")
}