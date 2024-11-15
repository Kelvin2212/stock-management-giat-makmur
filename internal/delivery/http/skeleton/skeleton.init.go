package skeleton

import (
	"context"
)

type SkeletonSvc interface {
	GetSkeleton(ctx context.Context) error
}

type (
	Handler struct {
		skeletonSvc SkeletonSvc
	}
)

func New(is SkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}
