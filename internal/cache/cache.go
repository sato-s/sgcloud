package cache

import (
	"encoding/json"
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

func NewCache() (*Cache, error) {
	file := cachefile()
	c = &Cache{}
	if _, err := os.Stat(file); err == nil {
		content = ioutil.ReadFile(file)
		err := json.Unmarshal(c, c)
		return c, err
	} else if errors.Is(err, os.ErrNotExist) {
		return c, nil
	} else {
		return nil, err
	}
}

func (c *Cache) Save() error {
	c.CachedAt = time.Now()
	if content, err = json.Marshal(c); err != nil {
		os.WriteFile()
	}
}

func cachefile() string {
	return path.Join(os.Tempdir, "sgcloud-cache-file.json")
}
