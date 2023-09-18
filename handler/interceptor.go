package handler

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func NewInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				slog.Error(
					err.Error(),
					slog.String("func", req.Spec().Procedure),
					slog.Any("req", req.Any()),
				)
			}

			return res, err
		}
	}
}
