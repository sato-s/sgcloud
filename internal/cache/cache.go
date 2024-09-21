package cache

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"github.com/sato-s/sgcloud/internal/projects"
)

const expirationDuration = 1 * time.Hour * 24

type Cache struct {
	Projects projects.Projects
	CachedAt time.Time
}

func NewCache() (*Cache, error) {
	file := cachefile()
	if _, err := os.Stat(file); err == nil {
		return readCacheJsonFile(file)
	} else if errors.Is(err, os.ErrNotExist) {
		return &Cache{}, nil
	} else {
		return nil, err
	}
}

func (c *Cache) Save() error {
	c.CachedAt = time.Now()
	if content, err := json.Marshal(c); err == nil {
		return os.WriteFile(cachefile(), content, 0600)
	} else {
		return err
	}
}

func (c *Cache) IsExpired() bool {
	now := time.Now()
	return now.After(c.CachedAt.Add(expirationDuration))
}

func cachefile() string {
	return path.Join(os.TempDir(), "sgcloud-cache-file.json")
}

func readCacheJsonFile(file string) (*Cache, error) {
	c := &Cache{}
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, c)
	return c, err
}
