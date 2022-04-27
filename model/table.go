package model

import sb "github.com/dropbox/godropbox/database/sqlbuilder"

const DATABASE_NAME = "test_db"

var EmptyContent = ""

func AdminTable() *sb.Table {
	t1 := sb.NewTable(
		"t_admin",
		sb.IntColumn("id", sb.NotNullable),
		sb.StrColumn("user_name", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.StrColumn("user_password", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.IntColumn("add_time", sb.NotNullable),
		sb.IntColumn("update_time", sb.NotNullable),
		sb.StrColumn("role", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.IntColumn("is_deleted", sb.NotNullable),
		sb.StrColumn("email", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
	)

	return t1
}

func BlogTable() *sb.Table {
	t1 := sb.NewTable(
		"t_blog",
		sb.IntColumn("id", sb.NotNullable),
		sb.StrColumn("name", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.StrColumn("tag", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.IntColumn("add_time", sb.NotNullable),
		sb.IntColumn("update_time", sb.NotNullable),
		sb.IntColumn("is_deleted", sb.NotNullable),
		sb.StrColumn("article_id", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.IntColumn("read_count", sb.NotNullable),
		sb.StrColumn("content", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
		sb.StrColumn("abstract", sb.UTF8, sb.UTF8CaseInsensitive, sb.NotNullable),
	)

	return t1
}
