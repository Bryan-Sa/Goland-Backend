package users

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Bryan-Sa/Goland-Backend/internal/DTO"
	"github.com/Bryan-Sa/Goland-Backend/internal/service"
	"github.com/Bryan-Sa/Goland-Backend/internal/utils"
	"github.com/julienschmidt/httprouter"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginDTO DTO.UserLoginRequestBodyDTO
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	dtoerr := json.Unmarshal(data, &loginDTO)
	if dtoerr != nil {
		utils.SendJsonResponse(w, "Wrong type of data", http.StatusBadRequest)
		return
	}
	if errValidator := utils.ValidateStruct(loginDTO); errValidator != nil {
		fmt.Println(errValidator)
		utils.SendJsonResponse(w, errValidator.Error(), http.StatusBadRequest)
		return
	}
	service.RegisterUser(loginDTO)
	//appeler le service

}
