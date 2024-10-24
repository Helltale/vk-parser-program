package flags

import (
	"github.com/Helltale/vk-parser-program/internal/logger"
)

type FlagHandler interface {
	Handle(logger *logger.CombinedLogger) *Entry
}

type UserFlagHandler struct {
	Users StringSliceFlag
}

type WallFlagHandler struct {
	Wall StringSliceFlag
}

type AllFlagHandler struct {
	All StringSliceFlag
}

func (u *UserFlagHandler) Handle(logger *logger.CombinedLogger) *Entry {
	if len(u.Users) > 0 {
		entry := NewEntry("user", u.Users)
		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}
	return nil
}

func (w *WallFlagHandler) Handle(logger *logger.CombinedLogger) *Entry {
	if len(w.Wall) > 0 {
		entry := NewEntry("wall", w.Wall)
		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}
	return nil
}

func (a *AllFlagHandler) Handle(logger *logger.CombinedLogger) *Entry {
	if len(a.All) > 0 {
		entry := NewEntry("all", a.All)
		logger.Info("get input", "flag", entry.Flag, "value", entry.Value)
		return entry
	}
	return nil
}
