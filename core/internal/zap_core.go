package internal

import (
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	level  zapcore.Level
	fields []zapcore.Field // With 附加的常驻字段(如 node/app_id/env),Write 时并入输出
	zapcore.Core
}

//func NewZapCore(level zapcore.Level) *ZapCore {
//	entity := &ZapCore{level: level}
//	syncer := entity.WriteSyncer()
//	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
//		return l == level
//	})
//	entity.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, levelEnabler)
//	return entity
//}
