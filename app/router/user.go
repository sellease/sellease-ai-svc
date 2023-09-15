package router

// func (c *RouterContext) UserRoutes(r *gin.RouterGroup) {
// 	userHandler := c.Handler.UserHandler

// 	r.POST("/login",
// 		middleware.Validate[request.UserLoginRequest](consts.TagUserLoginRequest),
// 		userHandler.HandleUserLogin)
// 	r.POST("/image",
// 		middleware.Validate[request.AddUserImageRequest](consts.TagAddUserImageRequest),
// 		userHandler.HandleAddUserImage)
// 	r.GET("/image",
// 		middleware.Authorize(c.Usecase.User),
// 		userHandler.HandleGetUserImage)
// }
