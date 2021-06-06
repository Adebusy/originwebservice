package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/", AppStarted)
	route.Run("8016")

}

func AppStarted(c *gin.Context) {
	c.JSON(http.StatusOK, "Application is up.")

}

//if err := ctx.ShouldBindJSON(&visitorRequestBody); err != nil {
//	httputil.NewError(ctx, http.StatusBadRequest, err)
//return
//}
//r.GET("/user/:name", func(c *gin.Context) {
//	user := c.Params.ByName("name")
//value, ok := db[user]
//if ok {
//	c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//} else {
//	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//}
//})
