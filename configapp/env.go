package configapp

import "os"

func readEnv() {
	readEnvDb()
}

func readEnvDb() {

	configData = configDataStruct{
		Db: dbData{
			User: os.Getenv("DB_USER"),
			Pwd:  os.Getenv("DB_PWD"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
