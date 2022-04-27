package model

import (
	"golang-blog-journey/util/db"
	"golang-blog-journey/util/log"
	"time"

	sb "github.com/dropbox/godropbox/database/sqlbuilder"
)

type Blog struct {
	Id       string
	Name     string
	Date     int64
	Tag      string
	Count    int
	Content  string
	Abstract string
}

func getBlogData(requestId string) ([]*Blog, error) {
	b := BlogTable()
	q := b.Select(
		b.C("name"),
		b.C("tag"),
		b.C("update_time"),
		b.C("article_id"),
		b.C("read_count"),
		b.C("abstract"),
	).Where(
		sb.EqL(b.C("is_deleted"), 0),
	)

	req, err := q.String("test_db")
	if err != nil {
		log.Errorw("getBlogData", "RequestID", requestId, "err", err)
		return nil, err
	}

	log.Infow("getBlogData", "RequestID", requestId, "sql", req)
	rows, err := db.QuerySQL(req)
	if err != nil {
		log.Errorw("getBlogData", "RequestID", requestId, "err", err)
		return nil, err
	}

	res := make([]*Blog, 0)
	for rows.Next() {
		r := new(Blog)
		err := rows.Scan(&r.Name, &r.Tag, &r.Date, &r.Id, &r.Count, &r.Abstract)
		if err != nil {
			log.Errorw("getBlogData", "RequestID", requestId, "err", err)
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func getBlogSimpleData(articleId, requestId string) ([]*Blog, error) {
	b := BlogTable()
	q := b.Select(
		b.C("name"),
		b.C("tag"),
		b.C("update_time"),
		b.C("article_id"),
		b.C("read_count"),
		b.C("content"),
	).Where(
		sb.And(
			sb.EqL(b.C("is_deleted"), 0),
			sb.EqL(b.C("article_id"), articleId),
		),
	)

	req, err := q.String("test_db")
	if err != nil {
		log.Errorw("getBlogSimpleData", "RequestID", requestId, "err", err)
		return nil, err
	}

	log.Infow("getBlogSimpleData", "RequestID", requestId, "sql", req)
	rows, err := db.QuerySQL(req)
	if err != nil {
		log.Errorw("getBlogSimpleData", "RequestID", requestId, "err", err)
		return nil, err
	}

	res := make([]*Blog, 0)
	for rows.Next() {
		r := new(Blog)
		err := rows.Scan(&r.Name, &r.Tag, &r.Date, &r.Id, &r.Count, &r.Content)
		if err != nil {
			log.Errorw("getBlogSimpleData", "RequestID", requestId, "err", err)
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func addBlogData(name, tag, pageId, abstract, content, requestId string) error {
	b := BlogTable()
	q := b.Insert(
		b.C("name"),
		b.C("tag"),
		b.C("add_time"),
		b.C("update_time"),
		b.C("is_deleted"),
		b.C("article_id"),
		b.C("read_count"),
		b.C("content"),
		b.C("abstract"),
	)

	s := q.Add(
		sb.Literal(name),
		sb.Literal(tag),
		sb.Literal(time.Now().Unix()),
		sb.Literal(time.Now().Unix()),
		sb.Literal(0),
		sb.Literal(pageId),
		sb.Literal(0),
		sb.Literal(content),
		sb.Literal(abstract),
	)

	req, err := s.String("test_db")
	if err != nil {
		log.Errorw("addBlogData", "RequestID", requestId, "err", err)
		return err
	}

	log.Infow("addBlogData", "RequestID", requestId, "sql", req)
	_, err = db.ExecSQL(req)
	if err != nil {
		log.Errorw("addBlogData", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func updateBlogData(name, tag, articleId, abstract, content, requestId string, isDeleted int) error {
	b := BlogTable()
	q := b.Update()
	if name != EmptyContent {
		q.Set(b.C("name"), sb.Literal(name))
	}
	if tag != EmptyContent {
		q.Set(b.C("tag"), sb.Literal(tag))
	}
	if content != EmptyContent {
		q.Set(b.C("content"), sb.Literal(content))
	}
	if abstract != EmptyContent {
		q.Set(b.C("abstract"), sb.Literal(abstract))
	}
	q.Set(b.C("update_time"), sb.Literal(time.Now().Unix()))
	q.Set(b.C("is_deleted"), sb.Literal(isDeleted))

	s := q.Where(
		sb.And(
			sb.EqL(b.C("article_id"), articleId),
		),
	)

	req, err := s.String("test_db")
	if err != nil {
		log.Errorw("updateBlogData", "RequestID", requestId, "err", err)
		return err
	}

	log.Infow("updateBlogData", "RequestID", requestId, "sql", req)
	_, err = db.ExecSQL(req)
	if err != nil {
		log.Errorw("updateBlogData", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func updateReadCount(articleId, requestId string, count int) error {
	b := BlogTable()
	q := b.Update()
	q.Set(b.C("read_count"), sb.Literal(count+1))
	s := q.Where(
		sb.And(
			sb.EqL(b.C("article_id"), articleId),
		),
	)

	req, err := s.String("test_db")
	if err != nil {
		log.Errorw("updateReadCount", "RequestID", requestId, "err", err)
		return err
	}

	log.Infow("updateReadCount", "RequestID", requestId, "sql", req)
	_, err = db.ExecSQL(req)
	if err != nil {
		log.Errorw("updateReadCount", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}
