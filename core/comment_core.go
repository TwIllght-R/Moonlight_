package core

import "Moonlight_/repo"

type commentCore struct {
	commentRepo repo.CommentRepo
}

func NewCommentCore(commentRepo repo.CommentRepo) CommentCore {
	return commentCore{commentRepo: commentRepo}
}

func (r commentCore) NewComment(req New_comment_req) (*Comment_resp, error) {
	return nil, nil
}

func (r commentCore) GetComment(id string) (*Comment_resp, error) {
	return nil, nil
}
