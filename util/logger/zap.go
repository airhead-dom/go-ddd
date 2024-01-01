package logger

import "go.uber.org/zap"

// ### SHARED ZAP INSTANCE ###

type ZapInstance struct {
	zap *zap.Logger
}

var zapInstance = ZapInstance{}

func GetZap() *zap.Logger {
	if zapInstance.zap == nil {
		logger := zap.NewExample()

		zapInstance.zap = logger
	}

	return zapInstance.zap
}

// ### SHARED ZAP INSTANCE ###

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger() Logger {
	return &ZapLogger{
		zap: GetZap(),
	}
}

func (z *ZapLogger) Log(msg string) {
	z.zap.Info(msg)
}
