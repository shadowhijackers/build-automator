package states

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"projects/nizhal-build-automater/models"
)

func GetBuilderSetupState() models.BuildSetup {

	filePath := "./static/builder-setup.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}

	var BuilderSetupState models.BuildSetup
	_ = json.Unmarshal(file, &BuilderSetupState)

	return BuilderSetupState

}
