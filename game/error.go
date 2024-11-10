package game

import "fmt"

type GameError struct {
	scene string
	msg   string
}

func NewGameError(scene, msg string) error {
	return &GameError{scene: scene, msg: msg}
}

func (g *GameError) Error() string {
	return fmt.Sprintf("[scene: %s] %s", g.scene, g.msg)
}
