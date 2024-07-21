package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"` // コメントID
	ArticleID int       `json:"article_id"` // コメントの対象となった記事ID
	Message   string    `json:"message"`    // コメント本文
	CreatedAt time.Time `json:"created_at"` // 投稿日時
}

type Article struct {
	ID          int       `json:"article_id"` // 記事ID
	Title       string    `json:"title"`      // 記事タイトル
	Contents    string    `json:"contents"`   // 記事内容
	UserName    string    `json:"user_name"`  // 投稿者
	NiceNum     int       `json:"nice"`       // いいね数
	CommentList []Comment `json:"comments"`   // コメントリスト
	CreatedAt   time.Time `json:"created_at"` // 投稿時間
}
