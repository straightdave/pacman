package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneIntro struct{}

func NewSceneIntro() *SceneIntro {
	return &SceneIntro{}
}

func (s *SceneIntro) Name() string {
	return "intro"
}

func (s *SceneIntro) Update(ctx *Context) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		ctx.NextScene = "play"
	}
	return nil
}

func (s *SceneIntro) Draw(_ *Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press SPACE To Start")
}
