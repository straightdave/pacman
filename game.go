package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 320
)

type Game struct {
	ctx   *Context
	scene Scene
}

func NewGame() *Game {
	return &Game{
		ctx: &Context{
			NextScene: "intro",
		},
	}
}

func (g *Game) Update() error {
	if g.scene != nil {
		if err := g.scene.Update(g.ctx); err != nil {
			return err
		}
	}

	switch g.ctx.NextScene {
	case "intro":
		g.scene = NewSceneIntro()
		g.ctx.NextScene = ""
	case "play":
		g.scene = NewScenePlay()
		g.ctx.NextScene = ""
	case "win":
		g.scene = NewSceneWin()
		g.ctx.NextScene = ""
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}
	g.scene.Draw(g.ctx, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
