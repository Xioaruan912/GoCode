package app

import (
	"gogrep/command"
	"gogrep/file"
)

func Run() {
	config := flag.Checkflag()
	switch {
	case config.Path != "":
		file.PATH(config.Path, config.KeyWord)
	case config.FilePath != "":
		file.File(config.FilePath, config.KeyWord)
	case config.EXE != "":
		file.EXE(config.EXE, config.KeyWord)
	}

}
