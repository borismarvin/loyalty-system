// iter5
package main

import (
	"net/http"

	"github.com/borismarvin/loyalty-system.git/internal/app/config"
	"github.com/borismarvin/loyalty-system.git/internal/app/logger"
	// "github.com/borismarvin/loyalty-system.git/internal/app/handlers"
	// "github.com/borismarvin/shortener_url.git/internal/app/storage"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// // generate test body
	// // utils.GenerateBatchBody(100000)

	cfg := config.LoadConfig()
	err := logger.Initialize(cfg)
	if err != nil {
		panic(err)
	}
	// logger := handlers.NewLogger()
	// sugar := *logger.Sugar()

	// var store handlers.Storager
	// var err error

	// var storeMsg string
	// if len(cfg.DSN) > 0 {
	// 	store, err = storage.NewDBMS(cfg.DSN)
	// 	storeMsg = fmt.Sprintf("Store DBMS setup successfuly -> %s", cfg.DSN)
	// } else if len(cfg.FileStore) > 0 {
	// 	store, err = storage.NewFile(cfg.FileStore)
	// 	storeMsg = fmt.Sprintf("Store File setup successfuly -> %s", cfg.FileStore)
	// } else {
	// 	store, err = storage.NewMemory()
	// 	storeMsg = fmt.Sprintf("Store Memory setup successfuly -> %s", "default")
	// }

	// if err != nil {
	// 	sugar.Panicln("Store invalid")
	// 	os.Exit(1)
	// }

	// handler := handlers.New(cfg.BaseURL)

	// sugar.Infoln(storeMsg)
	// sugar.Infoln("Running server on", cfg.ServerAddress)
	return http.ListenAndServe(cfg.RunAddress, nil)
}
