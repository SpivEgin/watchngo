package conf

import (
	"fmt"

	"github.com/Leryan/watchngo/pkg/watcher"

	"github.com/go-ini/ini"
)

// WatchersFromPath returns configuration from file at path
func WatchersFromPath(path string) ([]*watcher.Watcher, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("conf: from path: %s: %v", path, err)
	}

	// we only have the DEFAULT section
	if len(cfg.Sections()) == 1 {
		return nil, fmt.Errorf("conf: no configuration")
	}

	watchers := make([]*watcher.Watcher, 0)

	defaultSection := cfg.Section(ini.DEFAULT_SECTION)

	debug := false
	if defaultSection.HasKey("debug") {
		debug, err = defaultSection.Key("debug").Bool()
		if err != nil {
			return nil, fmt.Errorf("conf: debug is not a bool: %v", err)
		}
	}

	// exclude the DEFAULT section, which comes first
	for _, section := range cfg.Sections()[1:] {
		name := section.Name()
		match := ""
		command := ""
		filter := ""
		wdebug := debug
		withShell := true

		if section.HasKey("match") {
			match = section.Key("match").String()
		} else {
			return nil, fmt.Errorf("conf: missing required match key")
		}

		if section.HasKey("command") {
			command = section.Key("command").String()
		} else {
			return nil, fmt.Errorf("conf: missing required command key: %v", err)
		}

		if section.HasKey("filter") {
			filter = section.Key("filter").String()
		}

		if section.HasKey("debug") {
			wdebug, err = section.Key("debug").Bool()
			if err != nil {
				return nil, fmt.Errorf("conf: debug is not a bool: %v", err)
			}
		}

		w, err := watcher.NewWatcher(
			name,
			match,
			filter,
			command,
			withShell,
			wdebug,
		)

		if err != nil {
			return nil, fmt.Errorf("conf: new watcher: %s: %v", name, err)
		}

		watchers = append(watchers, w)
	}

	return watchers, nil
}
