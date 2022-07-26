package restapisample

import (
	"context"
	"errors"
	"github.com/oklog/ulid/v2"
	cmap "github.com/orcaman/concurrent-map"

	"time"
)

//InMemoryRepo implements Repository in a very simple implementation (in memory).
type InMemoryRepo struct {
	m cmap.ConcurrentMap
}

//NewInMemoryRepo creates a new InMemoryRepo.
func NewInMemoryRepo() *InMemoryRepo {
	m := cmap.New()

	return &InMemoryRepo{m: m}
}

var errTypeNotExpected = errors.New("type not expected")

//Get gets a job from the repository. It will return the Job, a boolean indicating if found or not and an error.
func (repo InMemoryRepo) Get(_ context.Context, id string) (Job, bool, error) {
	data, found := repo.m.Get(id)

	if !found {
		return Job{}, false, nil
	}

	job, conversion := data.(Job)

	if !conversion {
		return Job{}, true, errTypeNotExpected
	}

	return job, true, nil
}

//Create creates a new entity in the repository.
func (repo InMemoryRepo) Create(_ context.Context, job Job) (string, error) {
	id := generateULID()

	repo.m.Set(id, job)

	return id, nil
}

//Update updates an entity in the repository, overwriting if found.
func (repo InMemoryRepo) Update(_ context.Context, id string, job Job) error {
	repo.m.Set(id, job)

	return nil
}

//Delete erases the entity in the repository.
func (repo InMemoryRepo) Delete(_ context.Context, id string) error {
	repo.m.Set(id, nil)

	return nil
}

func generateULID() string {
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, ulid.DefaultEntropy())

	if err != nil {
		panic(err)
	}

	return id.String()
}
