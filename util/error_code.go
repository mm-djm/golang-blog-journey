package util

import (
	"golang-blog-journey/controller"
)

func InitErrorMap() {
	controller.RegisterError("NAME_EXIST", "user name already exist,cannot create account", 1001)
	controller.RegisterError("ADD_ADMIN_ERROR", "create admin error,pls try again", 1002)
	controller.RegisterError("REMOVE_ADMIN_ERROR", "delete admin failure,pls try again", 1003)
	controller.RegisterError("UPDATE_ADMIN_ERROR", "update admin failure,pls try again", 1004)
	controller.RegisterError("GET_ADMIN_DATA_ERROR", "describe admin error,pls try again", 1005)

	controller.RegisterError("ADD_BLOG_ERROR", "create blog error,pls try again", 2001)
	controller.RegisterError("REMOVE_BLOG_ERROR", "delete blog failure,pls try again", 2002)
	controller.RegisterError("UPDATE_BLOG_ERROR", "update blog failure,pls try again", 2003)
	controller.RegisterError("GET_BLOG_DATA_ERROR", "describe blog error,pls try again", 2004)
	controller.RegisterError("BLOG_ID_NOT_EXIST", "blog id not exist,cannot update blog", 2005)

	controller.RegisterError("LOGIN_ERROR", "login error,pls check email or password", 3001)
	controller.RegisterError("LOGOUT_ERROR", "logout error,pls check log", 3002)
}
