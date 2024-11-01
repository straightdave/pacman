package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	IsActive() bool
	Activate()
	Deactivate()

	Draw(*Context, *ebiten.Image)
	Update(*Context) error
}

func NewScene(name string) Scene {
	switch name {
	case "intro":
		return NewSceneIntro()

	case "play":
		return NewScenePlay()

	case "error":
		return NewSceneError()

	case "win":
		return NewSceneWin()

	default:
		return nil
	}
}
