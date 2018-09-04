package controller
import (
	"miniproject/model"
	"net/http"
	"encoding/json"
)



	func GetAll(w http.ResponseWriter, r *http.Request) {

		var contohs = []model.Contoh{
			{ID: 1, Code: "ME001", Name: "Menu1", CreatedBy:"admin1"},
			{ID: 1, Code: "ME001", Name: "Menu1", CreatedBy:"admin1"},
			{ID: 1, Code: "ME001", Name: "Menu1", CreatedBy:"admin1"},
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(contohs)
	}
	
