package service

import (
	"fmt"

	"github.com/Bryan-Sa/Goland-Backend/internal/DTO"
	"github.com/Bryan-Sa/Goland-Backend/internal/entities"
	"github.com/Bryan-Sa/Goland-Backend/internal/repository"
)

func RegisterUser(loginDTO DTO.UserLoginRequestBodyDTO) {
	user := &entities.UserEntitie{}
	result, err := repository.FindAll("user", user)
	if err != nil {
		fmt.Println("Erreur lors de la récupération depuis la db")
		fmt.Println(err)
	}
	fmt.Println("test moi")
	fmt.Println(result)
}
