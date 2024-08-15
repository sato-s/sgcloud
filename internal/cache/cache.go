package cache

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/sato-s/sgcloud/internal/projects"
)

type Cache struct {
	Projects projects.Projects
	CachedAt time.Time
}

func NewCache() (c *Cache, isCacheExist bool, err error) {
	file := cachefile()
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return &Cache{}, false, nil
	} else if {
	}
	ioutil.ReadFile(cachefile())
}

func cachefile() string {
	return path.Join(os.Tempdir, "sgcloud-cache-file.json")
}
