package main

import "github.com/gin-gonic/gin"

func RegAccountsRouter(r *gin.Engine, routeName string) {
	accountRouter := r.Group(routeName)
	accountRouter.POST("", createAcc)
	accountRouter.POST("/:id/deposit", deposit)
	accountRouter.POST("/:id/withdraw", withdraw)
	accountRouter.GET("/:id/balance", balance)
}
