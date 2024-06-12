package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	RunAddress    string `env:"RUN_ADDRESS"`
	SystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DSN           string `env:"DATABASE_URI"`
}

// загружаем данные среды окружения
func LoadConfig() (config Config) {

	flag.StringVar(&config.RunAddress, "a", "localhost:8080", "address and port to run server")
	// адрес и порт куда отправлять сокращатель
	flag.StringVar(&config.SystemAddress, "r", "localhost:8080", "address for accural system")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.StringVar(&config.DSN, "d", "", "access to DBMS")
	// if !strings.Contains(config.DSN, "sslmode") {
	// 	config.DSN = fmt.Sprintf("%s?sslmode=disable", config.DSN)
	// }
	flag.Parse()
	if err := env.Parse(&config); err != nil {
		panic(fmt.Errorf("error while parsing config: %w", err))
	}
	return
}
