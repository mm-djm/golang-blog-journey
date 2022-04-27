package model

import (
	"errors"
	"golang-blog-journey/util/log"
	"strings"
)

func CheckBlogExist(id, requestId string) (bool, error) {
	data, err := getBlogSimpleData(id, requestId)
	if err != nil {
		log.Errorw("CheckBlogExist", "RequestID", requestId, "err", err)
		return false, err
	}
	if len(data) == 0 {
		return false, nil
	}
	return true, nil
}

func GetSimpleBlogData(id, requestId string) ([]*Blog, error) {
	data, err := getBlogSimpleData(id, requestId)
	if err != nil {
		log.Errorw("GetSimpleBlogData", "RequestID", requestId, "err", err)
		return nil, err
	}
	if len(data) == 0 {
		log.Errorw("GetSimpleBlogData", "RequestID", requestId, "err", "no data")
		return nil, errors.New("no data")
	}
	return data, nil
}

func GetBlogData(requestId string) ([]*Blog, error) {
	data, err := getBlogData(requestId)
	if err != nil {
		log.Errorw("GetBlogData", "RequestID", requestId, "err", err)
		return nil, err
	}
	if len(data) == 0 {
		log.Errorw("GetBlogData", "RequestID", requestId, "err", "no data")
		return nil, errors.New("no data")
	}
	return data, nil
}

func AddBlog(name string, tag []string, pageId, abstract, content, requestId string) error {
	tagSet := strings.Join(tag, ",")
	err := addBlogData(name, tagSet, pageId, abstract, content, requestId)
	if err != nil {
		log.Errorw("AddBlog", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func RemoveBlog(id, requestId string) error {
	err := updateBlogData("", "", id, "", "", requestId, 1)
	if err != nil {
		log.Errorw("RemoveBlog", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func UpdateBlog(name string, tag []string, articleId, abstract, content, requestId string) error {
	tagSet := strings.Join(tag, ",")
	err := updateBlogData(name, tagSet, articleId, abstract, content, requestId, 0)
	if err != nil {
		log.Errorw("UpdateBlog", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func UpdateReadCount(articleId, requestId string, count int) error {
	err := updateReadCount(articleId, requestId, count)
	if err != nil {
		log.Errorw("UpdateReadCount", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}
