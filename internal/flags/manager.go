package flags

import (
	"flag"
	"os"

	"github.com/Helltale/vk-parser-program/internal/logger"
)

type FlagManager struct {
	WallHandler *WallFlagHandler
	AllHandler  *AllFlagHandler
	UserHandler *UserFlagHandler
}

func NewFlagManager() *FlagManager {
	return &FlagManager{
		WallHandler: &WallFlagHandler{},
		AllHandler:  &AllFlagHandler{},
		UserHandler: &UserFlagHandler{},
	}
}

func (fm *FlagManager) FlagHandler(logger *logger.CombinedLogger) *Entry {
	wall := StringSliceFlag{}
	all := StringSliceFlag{}
	user := StringSliceFlag{}

	// Определяем флаги
	flag.Var(&wall, "wall", "get information about wall")
	flag.Var(&all, "all", "get all information")
	flag.Var(&user, "user", "get all information for users")

	flag.Parse()

	// Устанавливаем значения в обработчики
	fm.WallHandler.Wall = wall
	fm.AllHandler.All = all
	fm.UserHandler.Users = user

	// Проверяем, были ли переданы какие-либо флаги
	if len(wall) == 0 && len(all) == 0 && len(user) == 0 {
		logger.Error("no entry flags")
		os.Exit(1)
	}

	// Обрабатываем флаги и возвращаем первый найденный
	if entry := fm.WallHandler.Handle(logger); entry != nil {
		return entry
	}

	if entry := fm.AllHandler.Handle(logger); entry != nil {
		return entry
	}

	if entry := fm.UserHandler.Handle(logger); entry != nil {
		return entry
	}

	return nil
}
