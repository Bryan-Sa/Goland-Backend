package users

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// il faudrait verifier le contenue du body request
	// il faudrait faire un appel a un dto ?

	// appeler le service
	// Renvoyer la réponse du service
}
