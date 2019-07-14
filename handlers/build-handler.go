package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projects/nizhal-build-automater/states"
)

type RequestedBuildPayload struct {
	ProjectId string `json:"projectId"`
	EnvId     string `json:"envId"`
	UserId    string `json:"userId"`
}

func Build(res http.ResponseWriter, req *http.Request) {

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	var reqestedData RequestedBuildPayload
	err = json.Unmarshal(b, &reqestedData)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	var buildSetupStata = states.GetBuilderSetupState()
	for i := 0; i < len(buildSetupStata.Users); i++ {
		if buildSetupStata.Users[i].Id == reqestedData.UserId {
          cmds := buildSetupStata.Users[i].Projects[reqestedData.ProjectId].Environments[reqestedData.EnvId].Sh;
          log.Println(cmds)
		}
	}

	fmt.Fprintf(res, "%s", "YOU HAVE BEING HACKED")

	return
}
