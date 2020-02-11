package app
import ("fmt"
		"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication(){
	fmt.Println("ok")
	mapUrls()
	router.Run(":8080")
}