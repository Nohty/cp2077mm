package events

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func SendLog(ctx context.Context, log string) {
	runtime.EventsEmit(ctx, "log:new", log)
}

func SendError(ctx context.Context, message string) {
	runtime.EventsEmit(ctx, "error:new", message)
}

func SendSuccess(ctx context.Context, message string) {
	runtime.EventsEmit(ctx, "success:new", message)
}

func SendRefreshMods(ctx context.Context) {
	runtime.EventsEmit(ctx, "mods:refresh")
}
