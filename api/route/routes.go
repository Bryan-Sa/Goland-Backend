package route

import (
	"fmt"

	"github.com/Bryan-Sa/Goland-Backend/internal/controllers/users"
	"github.com/julienschmidt/httprouter"
)

func SetupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/login", users.RegisterUser)
	return router
}

func Tester() {
	fmt.Println("Félicitation Bryan")
}
