package users
import ("github.com/gin-gonic/gin"
		"net/http"
		"github.com/titiksha05/book_store_users_api/domain/users"
		"fmt"
		//"io/ioutil"
		//"encoding/json"
		"strconv"
		"github.com/titiksha05/book_store_users_api/services"
		"github.com/titiksha05/book_store_users_api/utils/errors"
)

func CreateUser(c *gin.Context){
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle JSON error
		restErr := errors.NewBadRequestError("invalid json body")
		fmt.Println(err)
		c.JSON(restErr.Status, restErr)
		return
	}
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
		//TODO: Handle error
		//return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
		//fmt.Println(err.Error())
		//return
	//}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation
		c.JSON(saveErr.Status, saveErr)
	}
	
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO: handle user creation
		c.JSON(getErr.Status, getErr)
	}
	
	c.JSON(http.StatusOK, user)
	//c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me")
}

