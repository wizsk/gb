package main

import "path/filepath"

type Config struct {
	RootDir string // root dir
	encExt  string
	decExt  string
	Key     string
	Editor  string
}

func (c *Config) addEncExt(n string) string {
	return n + c.encExt
}

// adds c.encExt with the name
func (c *Config) fullEncFilePath(n string) string {
	return filepath.Join(c.RootDir, c.addEncExt(n))
}

// func (c *Config) fullEncFilePath(n string) string {
// 	return filepath.Join(c.RootDir, n)
// }
