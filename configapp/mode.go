package configapp

import (
	"fmt"
	"os"
	"strings"
)

const (
	modeDev     string = "dev"
	modeHomolog string = "homolog"
	modeProd    string = "prod"
)

// Ambiente que o app está rodando: [dev, homolog, prod].
var modeProperties string

//****************************************************************

// GetMode retorna em qual ambiente o app está rodando: [dev, homolog, prod].
func GetMode() string {
	return modeProperties
}

func readMode() (string, error) {
	const fileProp string = "mode.properties"

	result, err := os.ReadFile(fileProp)
	if err != nil {
		return "", err
	}

	modeProperties = strings.TrimSpace(string(result))
	return modeProperties, nil
}

func validMode(modeRead string) error {

	ok := modeRead == modeDev ||
		modeRead == modeHomolog ||
		modeRead == modeProd

	if !ok {
		return fmt.Errorf("mode.properties invalid: %v", modeRead)
	}
	return nil
}
