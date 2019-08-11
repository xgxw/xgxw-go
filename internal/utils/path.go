package utils

import "path"

// CleanPath is 清洗路径
func CleanPath(p string) string {
	p = path.Clean(p)
	return p
}
