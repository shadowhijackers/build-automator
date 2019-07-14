package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"../states"
)

type RequestedUserPayload struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func Login(res http.ResponseWriter, req *http.Request) {

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	var reqestedData RequestedUserPayload
	err = json.Unmarshal(b, &reqestedData)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	log.Println(reqestedData)
	var buildSetupStata = states.GetBuilderSetupState()
	log.Println(buildSetupStata)

	for i := 0; i < len(buildSetupStata.Users); i++ {
		if buildSetupStata.Users[i].Name == reqestedData.User && buildSetupStata.Users[i].Password == reqestedData.Password {
			log.Println("login Sucess")
			var projects, _ = json.Marshal(buildSetupStata.Users[i].Projects)
			res.Header().Set("content-type", "application/json")
			_, _ = res.Write(projects);
			return
		}
	}
	loginFailureResData := make(map[string]interface{})
	status := &struct {
		Code    int `json:"code"`
		Message string `json:"message"`
	}{
		1,
		"password is incorrect",
	}

	loginFailureResData["status"] = status

	log.Println(loginFailureResData["status"])
	loginFailureRes, err := json.Marshal(loginFailureResData)
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), 500)
		return
	}
	_, _ = res.Write(loginFailureRes)

}
