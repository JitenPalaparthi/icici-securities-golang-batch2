package main

import (
	"demo/models"
	"demo/utils"
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var PORT string

//var DB_CONN string

func init() {
	// args := os.Args
	// fmt.Println(args)
	// --port=8084 -port= 8081
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8083", "--port=8083 or -port=8083 or --port 8083")
	}
	//flag.StringVar(&DB_CONN, "db", "some db", "something")
	flag.Parse()
}
func main() {
	router := gin.Default()

	// how to use gin for cross origin requests
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello ICICI")
	})

	router.POST("/users", func(ctx *gin.Context) {
		user := new(models.User)
		err := ctx.Bind(user)
		if err != nil {
			ctx.String(400, err.Error())
			ctx.Abort()
			return
		}
		user.Id = uint(rand.Intn(10000))
		user.Status = "active"
		user.LastModified = time.Now().Unix()
		utils.ChUser <- user // blocked ?
		//ctx.JSON(201, gin.H{"user-id": user.Id})
		ctx.JSON(201, user)

	})

	router.Run(":" + PORT)

}

// func Runall(str ...string) error {
// 	for _, st := range str {
// 		go http.ListenAndServe(st, nil)
// 	}
// 	return nil
// }
