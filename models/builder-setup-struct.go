package models


type Environment struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Sh   []string `json:"sh"`
}

type Project struct {
	Name         string `json:"name"`
	Id string `json:"id"`
	Environments map[string]Environment `json:"environments"`
}

type BuilderSetupStruct struct {
	Projects map[string]Project `json:"projects"`
}

type User struct {
	Id string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	BuilderSetupStruct
}

//BuildSetup @Deprecated
type BuildSetup struct {
	Users []User `json:"users"`
}
