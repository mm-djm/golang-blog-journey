package model

import (
	"golang-blog-journey/util/db"
	"golang-blog-journey/util/log"
	"time"

	sb "github.com/dropbox/godropbox/database/sqlbuilder"
)

type Admin struct {
	Email    string
	Name     string
	Password string
	Role     string
}

func getSimpleAdminData(email, requestId string) ([]*Admin, error) {
	a := AdminTable()
	q := a.Select(
		a.C("email"),
		a.C("user_name"),
		a.C("user_password"),
		a.C("role"),
	).Where(
		sb.And(
			sb.EqL(a.C("is_deleted"), 0),
			sb.EqL(a.C("email"), email),
		),
	)

	req, err := q.String("test_db")
	if err != nil {
		log.Errorw("getSimpleAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}

	log.Infow("getSimpleAdminData", "RequestID", requestId, "sql", req)
	rows, err := db.QuerySQL(req)
	if err != nil {
		log.Errorw("getSimpleAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}

	res := make([]*Admin, 0)
	for rows.Next() {
		r := new(Admin)
		err := rows.Scan(&r.Email, &r.Name, &r.Password, &r.Role)
		if err != nil {
			log.Errorw("getSimpleAdminData", "RequestID", requestId, "err", err)
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func getAdminData(requestId string) ([]*Admin, error) {
	a := AdminTable()
	q := a.Select(
		a.C("email"),
		a.C("user_name"),
		a.C("role"),
	).Where(
		sb.And(
			sb.EqL(a.C("is_deleted"), 0),
		),
	)

	req, err := q.String("test_db")
	if err != nil {
		log.Errorw("getAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}

	log.Infow("getAdminData", "RequestID", requestId, "sql", req)
	rows, err := db.QuerySQL(req)
	if err != nil {
		log.Errorw("getAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}

	res := make([]*Admin, 0)
	for rows.Next() {
		r := new(Admin)
		err := rows.Scan(&r.Email, &r.Name, &r.Role)
		if err != nil {
			log.Errorw("getAdminData", "RequestID", requestId, "err", err)
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func addAdminData(name, password, role, email, requestId string) error {
	a := AdminTable()
	q := a.Insert(
		a.C("user_name"),
		a.C("user_password"),
		a.C("add_time"),
		a.C("update_time"),
		a.C("role"),
		a.C("is_deleted"),
		a.C("email"),
	)

	s := q.Add(
		sb.Literal(name),
		sb.Literal(password),
		sb.Literal(time.Now().Unix()),
		sb.Literal(time.Now().Unix()),
		sb.Literal(role),
		sb.Literal(0),
		sb.Literal(email),
	)

	req, err := s.String("test_db")
	if err != nil {
		log.Errorw("addAdminData", "RequestID", requestId, "err", err)
		return err
	}

	log.Infow("addAdminData", "RequestID", requestId, "sql", req)
	_, err = db.ExecSQL(req)
	if err != nil {
		log.Errorw("addAdminData", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func updateAdminData(name, password, role, email, requestId string, isDeleted int) error {
	a := AdminTable()
	q := a.Update()
	if password != EmptyContent {
		q.Set(a.C("user_password"), sb.Literal(password))
	}
	if role != EmptyContent {
		q.Set(a.C("role"), sb.Literal(role))
	}
	if name != EmptyContent {
		q.Set(a.C("user_name"), sb.Literal(name))
	}
	q.Set(a.C("update_time"), sb.Literal(time.Now().Unix()))
	q.Set(a.C("is_deleted"), sb.Literal(isDeleted))

	s := q.Where(
		sb.And(
			sb.EqL(a.C("email"), email),
			sb.EqL(a.C("is_deleted"), 0),
		),
	)

	req, err := s.String("test_db")
	if err != nil {
		log.Errorw("updateAdminData", "RequestID", requestId, "err", err)
		return err
	}

	log.Infow("updateAdminData", "RequestID", requestId, "sql", req)
	_, err = db.ExecSQL(req)
	if err != nil {
		log.Errorw("updateAdminData", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}
