package zap

import (
	"github.com/xiecang/jarvis/log"
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	type fields struct {
		conf LoggerConfig
		log  *zap.Logger
		slog *zap.SugaredLogger
	}
	type args struct {
		v []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "string",
			fields: fields{
				conf: LoggerConfig{
					level: log.LevelDebug,
				},
			},
			args: args{
				v: []any{"test"},
			},
		},
		{
			name: "int",
			fields: fields{
				conf: LoggerConfig{
					level: log.LevelDebug,
				},
			},
			args: args{
				v: []any{1},
			},
		},
		{
			name: "zap",
			fields: fields{
				conf: LoggerConfig{
					level: log.LevelDebug,
				},
			},
			args: args{
				v: []any{zap.Any("test", "test"), zap.Any("test", 1), zap.String("test", "test"), zap.Int("test", 1)},
			},
		},
		{
			name: "struct",
			fields: fields{
				conf: LoggerConfig{
					level: log.LevelDebug,
				},
			},
			args: args{
				v: []any{struct {
					Name string `json:"name"`
					Age  int    `json:"age"`
				}{
					Name: "test",
					Age:  1,
				}},
			},
		},
		{
			name: "file",
			fields: fields{
				conf: LoggerConfig{
					level: log.LevelDebug,
					path:  "/tmp/1",
				},
			},
			args: args{
				v: []any{struct {
					Name string `json:"name"`
					Age  int    `json:"age"`
				}{
					Name: "test",
					Age:  1,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := NewLogger(log.LevelDebug, WithName(tt.fields.conf.name), WithPath(tt.fields.conf.path))
			z.Debug(tt.args.v...)
			z.Info(tt.args.v...)
			z.Info(172)
			z.slog.With("k", "v", "k1", "v1").Info("test")

			z.With(log.Field("a", "b")).Error("test")
		})
	}
}
