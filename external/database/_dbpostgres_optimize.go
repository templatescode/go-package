package dbpostgres

import (
	"os"
	"strconv"
	"time"
)

/*
Otimiza pool de conexões do banco de dados
*/
func (d dbpostgreser) Optimize() error {

	// valores default
	const (
		defaultDbMaxOpenConns           int = 1000
		defaultDbMaxIdleConns           int = 500
		defaultDbConnMaxLifetimeMinutes int = 30
	)

	// Se não existir variável de ambiente, preenche com default.
	_, ok := os.LookupEnv("DB_POSTGRES_MAX_OPEN_CONNS")
	if !ok {
		db.SetMaxOpenConns(defaultDbMaxOpenConns)
		db.SetMaxIdleConns(defaultDbMaxIdleConns)
		db.SetConnMaxLifetime(time.Minute * time.Duration(defaultDbConnMaxLifetimeMinutes))

	} else {
		// Busca variáveis de ambiente já parseadas para int
		envDbMaxOpenConns, err := getEnvDbMaxOpenConns()
		if err != nil {
			return err
		}
		envDbMaxIdleConns, err := getenvDbMaxIdleConns()
		if err != nil {
			return err
		}
		envDbConnMaxLifetimeMinutes, err := getenvDbConnMaxLifetime()
		if err != nil {
			return err
		}

		db.SetMaxOpenConns(envDbMaxOpenConns)
		db.SetMaxIdleConns(envDbMaxIdleConns)
		db.SetConnMaxLifetime(time.Minute * time.Duration(envDbConnMaxLifetimeMinutes))
	}
	return nil
}

/*
Busca variável de ambiente MaxOpenConns
MaxOpenConns define o número máximo de conexões abertas com o banco de dados.
*/
func getEnvDbMaxOpenConns() (int, error) {
	const envDbMaxOpenConns string = "DB_POSTGRES_MAX_OPEN_CONNS"
	v := os.Getenv(envDbMaxOpenConns)
	return strconv.Atoi(v)
}

/*
Busca variável de ambiente MaxIdleConns
MaxIdleConns define o número máximo de conexões no pool de conexões inativas.
*/
func getenvDbMaxIdleConns() (int, error) {
	const envDbMaxIdleConns string = "DB_POSTGRES_MAX_IDLE_CONNS"
	v := os.Getenv(envDbMaxIdleConns)
	return strconv.Atoi(v)
}

/*
Busca variável de ambiente ConnMaxLifetime
ConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.
*/
func getenvDbConnMaxLifetime() (int, error) {
	const envDbConnMaxLifetime string = "DB_POSTGRES_CONN_MAX_LIFE_TIME"
	v := os.Getenv(envDbConnMaxLifetime)
	return strconv.Atoi(v)
}
