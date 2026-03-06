package github

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"go.mattglei.ch/timber"
)

var (
	CLONE_DIRECTORY = "repositories"
	cloneLock       sync.Mutex
)

func SetupCloneFolder() error {
	_, err := os.Stat(CLONE_DIRECTORY)
	if !errors.Is(err, fs.ErrNotExist) {
		err = os.RemoveAll(CLONE_DIRECTORY)
		if err != nil {
			return fmt.Errorf("removing %s: %w", CLONE_DIRECTORY, err)
		}
	}

	err = os.MkdirAll(CLONE_DIRECTORY, 0755)
	if err != nil {
		return fmt.Errorf("creating directory %s: %w", CLONE_DIRECTORY, err)
	}
	return nil
}

func (r Repository) Clone() error {
	start := time.Now()

	cloneLock.Lock()
	destination := filepath.Join(CLONE_DIRECTORY, r.Name)
	_, err := os.Stat(destination)
	if !errors.Is(err, fs.ErrNotExist) {
		err = os.RemoveAll(destination)
		if err != nil {
			return fmt.Errorf("removing %s: %w", destination, err)
		}
	}

	repoURL, err := url.JoinPath("https://github.com/gleich", r.Name+".git")
	if err != nil {
		return fmt.Errorf("creating url: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	out, err := exec.CommandContext(ctx, "git", "clone", repoURL, destination).
		CombinedOutput()
	if err != nil {
		timber.Debug(string(out))
		return fmt.Errorf("running git clone: %w", err)
	}
	timber.DoneSince(start, "cloned", r.Name)
	cloneLock.Unlock()

	return nil
}

func (r Repository) EnsurePath(loc string) bool {
	destination := filepath.Join(CLONE_DIRECTORY, loc)
	_, err := os.Stat(destination)
	return !errors.Is(err, fs.ErrNotExist)
}
