package user

type UserSvc interface {
	// GetUser(ctx context.Context) error
}

type (
	Handler struct {
		userSvc UserSvc
	}
)

func New(iu UserSvc) *Handler {
	return &Handler{
		userSvc: iu,
	}
}
