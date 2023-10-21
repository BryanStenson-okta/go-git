package filesystem

import (
	"github.com/bryanstenson-okta/go-git/v5/plumbing/cache"
	"github.com/bryanstenson-okta/go-git/v5/storage"
	"github.com/bryanstenson-okta/go-git/v5/storage/filesystem/dotgit"
)

type ModuleStorage struct {
	dir *dotgit.DotGit
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	fs, err := s.dir.Module(name)
	if err != nil {
		return nil, err
	}

	return NewStorage(fs, cache.NewObjectLRUDefault()), nil
}
