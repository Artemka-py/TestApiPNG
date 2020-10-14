package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"

)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println("server started!!!")
	r.POST("/api/:text", func(c *gin.Context) {
		var obj interface{}

		bytes, _ := json.Marshal(c.Param("text"))
		err := json.Unmarshal(bytes, &obj)
		if err != nil {fmt.Println("Can't deserislize", err)}
		bytes2, _ := json.Marshal(obj)
		err2 := json.Unmarshal(bytes2, &obj)
		if err2 != nil {fmt.Println("Can't deserislize", err2)}
		c.JSON(http.StatusOK, obj)
	})
	addr := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	r.Run(addr)
}
