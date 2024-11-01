package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneError struct{}

func NewSceneError() *SceneError {
	return &SceneError{}
}

func (s *SceneError) IsActive() bool { return false }
func (s *SceneError) Activate()      {}
func (s *SceneError) Deactivate()    {}

func (s *SceneError) Update(_ *Context) error { return nil }

func (s *SceneError) Draw(ctx *Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Error: "+ctx.Message)
}
