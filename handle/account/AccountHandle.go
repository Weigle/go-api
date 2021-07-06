package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weigle/http/domain"
	"github.com/weigle/http/service"
)

type Data struct {
	votes string `json:"votes"`
	count string `json:"count,omitempty"`
}

func CreateHandle(r *mux.Router) {
	// maps := make(map[string]http.Handler)

	// maps["GET"] = "Maria"

	// r.HandleFunc("/hellow", hellow).Methods("GET")

	r.HandleFunc("/save", saveApplication).Methods("POST")

}

func hellow(w http.ResponseWriter, request *http.Request) {
	var data Data
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(body), data)
	fmt.Printf("%v => \n", data)
	// decoder := json.NewDecoder(request.Body)
}

func saveApplication(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var m domain.PropertyValue
	err = json.Unmarshal([]byte(body), &m)

	service.Save(m)

}
