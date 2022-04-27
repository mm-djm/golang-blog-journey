package controller

import (
	"github.com/gin-gonic/gin"

	"golang-blog-journey/util/middleware"
)

var ArticlePath string

func RunServer(addr string, port string) error {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()
	r.Use(middleware.Cors())

	app := r.Group("app")

	router := app.Group("api")
	{
		router.POST("/", API)

		router.POST("/login", LoginAdmin)

		router.POST("/register", RegisterAdmin)

		router.Use(middleware.JwtAPIToken()).POST("/logout", LogoutAdmin)
	}

	auth := app.Group("/admin")
	auth.Use(middleware.JwtToken())
	{
		auth.POST("/", Admin)
	}

	registerAdminHandler("DeleteAdmin", DeleteAdminHandler)
	registerAdminHandler("UpdateAdmin", UpdateAdminHandler)
	registerAdminHandler("DescribeSimpleAdmin", DescribeSimpleAdminHandler)
	registerAdminHandler("DescribeAdmin", DescribeAdminHandler)
	registerAdminHandler("CreateBlog", CreateBlogHandler)
	registerAdminHandler("UpdateBlog", UpdateBlogHandler)
	registerAdminHandler("DeleteBlog", DeleteBlogHandler)

	registerAPIHandler("DescribeBlog", DescribeBlogHandler)
	registerAPIHandler("DescribeSimpleBlog", DescribeSimpleBlogHandler)

	r.Run(port)
	return nil
}
