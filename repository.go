package restapisample

import "context"

//Repository implements the behaviour to get, create, update and delete information from a general repository.
type Repository interface {
	Get(ctx context.Context, id string) (Job, bool, error)
	Create(ctx context.Context, job Job) (string, error)
	Update(ctx context.Context, id string, job Job) error
	Delete(ctx context.Context, id string) error
}
