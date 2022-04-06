package configapp

// Start carrega todas as configurações necessárias.
func Start() error {

	// STEP 1: identifica qual ambiente rodará a aplicação.
	modeRead, err := readMode()
	if err != nil {
		return err
	}

	if err := validMode(modeRead); err != nil {
		return err
	}

	// STEP 2: recebe variáveis de ambiente.
	readEnv()

	return nil
}
