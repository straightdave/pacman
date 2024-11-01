package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneIntro struct {
	isActive bool
}

func NewSceneIntro() *SceneIntro {
	return &SceneIntro{
		isActive: true,
	}
}

func (s *SceneIntro) IsActive() bool { return s.isActive }
func (s *SceneIntro) Activate()      { s.isActive = true }
func (s *SceneIntro) Deactivate()    { s.isActive = false }

func (s *SceneIntro) Update(ctx *Context) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		ctx.NextScene = "play"
		s.Deactivate()
	}
	return nil
}

func (s *SceneIntro) Draw(_ *Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press SPACE To Start")
}
