package internal

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapCore struct {
	level  zapcore.Level
	fields []zapcore.Field // With 附加的常驻字段(如 node/app_id/env),Write 时并入输出
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(global.GVA_CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return entity
}

// WriteSyncer 返回只写文件的 syncer（不含控制台）
// formats 用于 business/folder/directory 子目录路由
func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.GVA_CONFIG.Zap.Director,
		z.level.String(),
		global.GVA_CONFIG.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	return zapcore.AddSync(cutter)
}
