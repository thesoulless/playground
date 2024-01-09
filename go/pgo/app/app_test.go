package app

import (
	"context"
	"log/slog"
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	ctx := context.Background()
	app := New(slog.Default())
	if err := app.addNoise(ctx, b.N); err != nil {
		b.Errorf("generateLoad got err %v want nil", err)
	}
}
