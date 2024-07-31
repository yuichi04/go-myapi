package testdata

import (
	"go-myapi/models"
)

var ArtcileTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  4,
	},
}

var InsertArticleData = []models.Article{
	models.Article{
		Title:    "insertTest1",
		Contents: "testtest1",
		UserName: "saki",
	},
	models.Article{
		Title:    "insertTest2",
		Contents: "testtest2",
		UserName: "saki",
	},
}
