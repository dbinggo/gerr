package gerr

import "context"

func LogError(logger Logger, error Error) {
	logger.Error(error)
}

func LoggerCtx(ctx context.Context, logger Logger, err Error) {
	logger.ErrorCtx(ctx, err)
}

func (e *CodeErr) Log(logger Logger) {
	logger.Error(e)
}
func (e *CodeErr) LogCtx(ctx context.Context, logger Logger) {
	logger.ErrorCtx(ctx, e)
}
