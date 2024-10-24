package flags

import (
	"flag"
	"os"
	"strings"

	"github.com/Helltale/vk-parser-program/internal/logger"
)

type Entry struct {
	Flag  string
	Value string
}

func NewEntry(flag_, value_ string) *Entry {
	return &Entry{
		Flag:  flag_,
		Value: value_,
	}
}

func FlagHandler(logger *logger.CombinedLogger) *Entry {
	wall := flag.String("wall", "", "get information about wall")
	all := flag.String("all", "", "get all information")
	user := flag.String("user", "", "get all information")

	flag.Parse()

	if *wall == "" && *all == "" && *user == "" {
		logger.Error("no entery flags")
		os.Exit(1)
	}

	if *wall != "" {
		wallSlice := strings.Split(*wall, " ")
		entry := NewEntry("wall", wallSlice[0])

		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}

	if *all != "" {
		allSlice := strings.Split(*all, " ")
		entry := NewEntry("wall", allSlice[0])

		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}

	if *user != "" {
		userSlice := strings.Split(*user, " ")
		entry := NewEntry("user", userSlice[0])

		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}

	return nil
}
