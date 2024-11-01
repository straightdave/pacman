package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneWin struct{}

func NewSceneWin() *SceneWin {
	return &SceneWin{}
}

func (s *SceneWin) IsActive() bool { return false }
func (s *SceneWin) Activate()      {}
func (s *SceneWin) Deactivate()    {}

func (s *SceneWin) Update(_ *Context) error { return nil }

func (s *SceneWin) Draw(_ *Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "WIN!")
}
