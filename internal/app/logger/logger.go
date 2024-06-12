package logger

import (
	"github.com/borismarvin/loyalty-system.git/internal/app/config"
	"go.uber.org/zap"
)

var sugar zap.SugaredLogger
var Log *zap.Logger = zap.NewNop()

func Initialize(config config.Config) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()

	// делаем регистратор SugaredLogger
	sugar = *logger.Sugar()

	// записываем в лог, что сервер запускается
	return err
}
