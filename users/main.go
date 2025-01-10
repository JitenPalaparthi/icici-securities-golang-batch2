package main

import (
	"demo/database"
	"demo/handlers"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT string

var DB_CONN string

func init() {
	// args := os.Args
	// fmt.Println(args)
	// --port=8084 -port= 8081
	PORT = os.Getenv("PORT")
	DB_CONN = os.Getenv("DB_CONN")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8083", "--port=8083 or -port=8083 or --port 8083")
	}
	if DB_CONN == "" {
		flag.StringVar(&DB_CONN, "db", `host=localhost user=admin password=admin123 dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--db=host=localhost user=admin password=admin123 dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	}
	flag.Parse()
}
func main() {
	// t1 := T1{}
	// t2 := T2{}
	//fmt.Println(unsafe.Sizeof(t1), "---", unsafe.Sizeof(t2))
	router := gin.Default()
	//db, err := database.GetConnection(DB_CONN)
	DB_CONN = `admin:admin123@tcp(localhost:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := database.GetMySQLConnection(DB_CONN)
	if err != nil {
		panic(err)
	}

	go database.NewAudit(db).Audit()
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

	userhandler := handlers.NewUserHandler(database.NewUser(db))

	router.POST("/user", userhandler.CreateUser)
	router.GET("/user/:id", userhandler.GetUserByID)
	router.DELETE("/user/:id", userhandler.DeleteByID)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello ICICI")
	})

	router.Run(":" + PORT)
}

// type T1 struct {
// 	ok   bool    //1
// 	no   float64 // 8
// 	done bool    //
// }

// type T2 struct {
// 	no   float32 // 4
// 	ok   bool    // 4
// 	done bool    // 1
// }
