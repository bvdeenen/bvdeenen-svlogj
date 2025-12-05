package config

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"maps"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"svlogj/pkg/svlog"
	"svlogj/pkg/types"
	"svlogj/pkg/utils"

	"github.com/adrg/xdg"
)

func ParseAndStoreConfig() {
	config := generateConfig()
	storeConfig(config)
}

func generateConfig() types.Config {
	facilities := make(map[string]struct{})
	levels := make(map[string]struct{})
	what := make(map[string]struct{})
	entities := make(map[string]struct{})
	services := make([]string, 10)
	parse := func(line string) {
		re := regexp.MustCompile(`^([^.]+)\.([^:]+)(?::?(.*))$`)
		if 0 == len(line) || line == "*" {
			return
		}
		m := re.FindStringSubmatch(line)
		facilities[m[1]] = struct{}{}
		levels[m[2]] = struct{}{}
		what[m[3]] = struct{}{}
	}

	root := "/var/log/socklog"
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		utils.Check(err)
		if d.IsDir() && path != root {
			services = append(services, d.Name())
		}
		if d.Name() == "config" && !d.IsDir() {
			file, err := os.Open(path)
			utils.Check(err)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				switch line[0] {
				case '#':
					break
				case '+':
					parse(strings.TrimSpace(line[1:]))
					break
				case '-':
					parse(strings.TrimSpace(line[1:]))
					break
				case '!':
					break
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	})

	// parse through all of the socklog files to find all the entities
	svlog.ParseLog(true, func(info types.Info, parse_config types.ParseConfig) {
		entities[info.Entity] = struct{}{}
	}, types.ParseConfig{})
	config := types.Config{
		Facilities: utils.RemoveEmptyStrings(slices.Collect(maps.Keys(facilities))),
		Levels:     utils.RemoveEmptyStrings(slices.Collect(maps.Keys(levels))),
		Entities:   utils.RemoveEmptyStrings(slices.Collect(maps.Keys(entities))),
		Services:   utils.RemoveEmptyStrings(services),
	}
	return config
}

func LoadConfig() types.Config {
	bytes, err := ioutil.ReadFile(configFile())
	utils.Check(err)
	var config types.Config
	err = json.Unmarshal(bytes, &config)
	utils.Check(err)
	return config
}

func storeConfig(config types.Config) {
	b, err := json.MarshalIndent(config, "", "  ")
	utils.Check(err)
	configFile := configFile()
	err = os.MkdirAll(path.Dir(configFile), 0700)
	utils.Check(err)
	f, err := os.Create(configFile)
	utils.Check(err)
	defer f.Close()
	f.Write(b)
}

func configFile() string {
	configFile, _ := xdg.ConfigFile("svlogj.json")
	return configFile
}
