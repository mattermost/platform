package pluginenv

import (
	"io/ioutil"
	"path/filepath"
)

// Performs a full scan of the given path.
//
// This function will return info for all subdirectories that appear to be plugins (i.e. all
// subdirectories containing plugin manifest files, regardless of whether they could actually be
// parsed).
//
// Plugins are found non-recursively and paths beginning with a dot are always ignored.
func ScanSearchPath(path string) ([]*BundleInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var ret []*BundleInfo
	for _, file := range files {
		if !file.IsDir() || file.Name()[0] == '.' {
			continue
		}
		if info := BundleInfoForPath(filepath.Join(path, file.Name())); info.ManifestPath != "" {
			ret = append(ret, info)
		}
	}
	return ret, nil
}
