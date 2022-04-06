package configapp

type configDataStruct struct {
	Db dbData
}

type dbData struct {
	User string
	Pwd  string
	Host string
	Port string
	Name string
}

var configData configDataStruct

//*****************************************************

// GetData retorna todos os dados de configuração.
func GetData() configDataStruct {
	return configData
}
