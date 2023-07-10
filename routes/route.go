package routes

import (
	"serendipity_backend/controllers"
	"serendipity_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// User Route -> Register, Login
	router.POST("/api/v1/user/create", controllers.CreateUser())
	router.POST("/api/v1/user/signin", controllers.SignInUser())
	router.POST("/api/v1/user/social/google_auth", controllers.GoogleAuth())
	router.POST("/api/v1/user", controllers.CheckUser())
	router.POST("/api/v1/forgotPassword", controllers.SendMail())
	router.POST("/api/v1/reset", controllers.ResetCode())
	router.POST("/api/v1/password_reset", controllers.ResetPassword())
	router.POST("/api/v1/set_logout", controllers.SetLogoutWithID())
}

func AdminRoutes(router *gin.Engine) {
	// User Route
	router.GET("/api/v1/admin/users", controllers.GetAllUsers())                   //-> Super Admin, Admin
	router.PUT("/api/v1/admin/user/:userId", controllers.UpdateUserInfo())         //-> Super Admin, Admin
	router.PUT("/api/v1/admin/user/role/:userId", controllers.UpdateUserRole())    //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/user/:userId", controllers.DeleteUser())          //-> Super Admin, Admin
	router.POST("/api/v1/admin/users/:userId", controllers.UpdateUserActivation()) //-> Super Admin, Admin

	// Toolkit Route
	router.POST("/api/v1/admin/toolkit/create", controllers.AddNewToolkit())                     //-> Super Admin, Admin
	router.POST("/api/v1/admin/toolkit/post/create", controllers.AddNewToolkitPost())            //-> Super Admin, Admin
	router.PUT("/api/v1/admin/toolkit/:toolkitId", controllers.UpdateToolkit())                  //-> Super Admin, Admin
	router.PUT("/api/v1/admin/toolkit/posts/:postId", controllers.UpdateToolkitPost())           //-> Super Admin, Admin
	router.PUT("/api/v1/admin/toolkit/posts/weekly/:postId", controllers.SetToolkitWeeklyPost()) //-> Super Admin, Admin
	router.POST("/api/v1/admin/toolkit/posts/today_activity", controllers.SetTodayActivities())  //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/toolkit/:toolkitId", controllers.DeleteToolkitWithId())         //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/toolkit/posts/:postId", controllers.DeleteToolkitPost())        //-> Super Admin, Admin

	// Forum Route
	router.POST("/api/v1/admin/forum/create", controllers.AddForumType())                      //-> Super Admin, Admin
	router.PUT("/api/v1/admin/forum/:forumId", controllers.UpdateForumTypeWithID())            //-> Super Admin, Admin
	router.DELETE("/api/admin/v1/forum/:forumId", controllers.DeleteForumWithId())             //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/forum/post/:postId", controllers.DeleteForumPostWithPostId()) //-> Super Admin, Admin

	// Foundation Route
	router.POST("/api/v1/admin/foundation/create", controllers.AddNewFoundation()) //-> Super Admin, Admin
	router.GET("/api/v1/admin/allFoundations", controllers.GetAllFoundations())    //-> Super Admin, Admin

	// Marketplace Route
	router.POST("/api/v1/admin/marketplace/add", controllers.CreatNewMarketplace())                  //-> Super Admin, Admin
	router.POST("/api/v1/admin/marketplace/new_item", controllers.CreateNewMarketplaceItem())        //-> Super Admin, Admin
	router.PUT("/api/v1/admin/marketplace/:marketplaceId", controllers.UpdateMarketPlace())          // -> Super Admin, Admin
	router.PUT("/api/v1/admin/marketplace/items/:itemId", controllers.UpdateMarketplacePost())       //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/marketplace/:marketplaceId", controllers.DeleteMarketplaceWithId()) //-> Super Admin, Admin
	router.DELETE("/api/v1/admin/marketplace/items/:itemId", controllers.DeleteMarketplaceItem())    //-> Super Admin, Admin
}

func SerendipityClientRoutes(router *gin.Engine) {

	// User Route -> Update
	router.PUT("/api/v1/user/:userId", controllers.UpdateProfile())                       //-> Super Admin, Admin, User
	router.POST("/api/v1/user/follow/:userId", controllers.HandleFollow())                //-> User
	router.POST("/api/v1/user/updateFoundation/:userId", controllers.SetUserFoundation()) //-> User
	router.GET("/api/v1/user/logout", controllers.LogOut())                               //-> User

	// Toolkit Route
	router.GET("/api/v1/toolkits", controllers.GetAllToolkitTypes())                     //-> User
	router.GET("/api/v1/toolkit/posts/:toolkitType", controllers.GetAllPostsInToolkit()) //-> User
	router.GET("/api/v1/toolkit/posts/today", controllers.GetToolkitPostsForToday())     //-> User
	router.GET("/api/v1/toolkit/posts/weekly_posts", controllers.GetWeeklyPosts())       //-> User

	// Forum Route
	router.GET("/api/v1/forums", controllers.GetAllForums())                                                      //-> Super Admin, Admin, User
	router.POST("/api/v1/forum/post/create", controllers.AddForumPost())                                          //-> User
	router.GET("/api/v1/forum/posts/:forumId", controllers.GetForumPostsWithId())                                 //-> Super Admin, Admin, User
	router.GET("/api/v1/forum/post/:postId", controllers.GetForumPostWithPostId())                                //-> Super Admin, Admin, User
	router.PUT("/api/v1/forum/post/:postId/:userId", controllers.UpdateForumPostWithPostIDUserID())               //-> User
	router.POST("/api/v1/forum/post/comment/:postId", controllers.AddCommentForForumPost())                       //-> User
	router.PUT("/api/v1/forum/post/comment/emotion/:postId/:commentId", controllers.AddForumPostCommentEmotion()) //-> User
	router.PUT("/api/v1/forum/post/emotions/:postId", controllers.UpdateForumPostWithEmotions())                  //-> User
	router.PUT("/api/v1/forum/post/visitcount/:postId", controllers.UpdateForumPostWithVisitCount())              //-> User

	// Marketplace Route
	router.GET("/api/v1/marketplace/all", controllers.GetAllMarketplaces())                              //-> Super Admin, Admin, User
	router.GET("/api/v1/marketplace/items/:marketplaceId", controllers.GetAllMarketplaceItemsWithType()) //-> Super Admin, Admin, User

}

func SerendipityRoute(router *gin.Engine) {
	UserRoutes(router)
	router.Use(middlewares.DeserializeUser())
	AdminRoutes(router)
	SerendipityClientRoutes(router)
}
