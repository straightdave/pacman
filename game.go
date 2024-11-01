package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ctx       *Context
	currScene Scene
	scenes    map[string]Scene
}

func NewGame() *Game {
	return &Game{
		ctx:    &Context{NextScene: "intro"},
		scenes: make(map[string]Scene),
	}
}

func (g *Game) Update() error {
	g.changeScene()

	for _, s := range g.scenes {
		if !s.IsActive() {
			continue
		}

		if err := s.Update(g.ctx); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.currScene != nil {
		g.currScene.Draw(g.ctx, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 320
}

func (g *Game) changeScene() {
	if s, ok := g.scenes[g.ctx.NextScene]; ok {
		g.currScene = s
		return
	}

	if !g.loadScene(g.ctx.NextScene) {
		g.setErrorPage(fmt.Sprintf("No scene with name %s.", g.ctx.NextScene))
	}
}

func (g *Game) loadScene(name string) bool {
	if _, ok := g.scenes[name]; ok {
		return true
	}

	if s := NewScene(name); s != nil {
		g.scenes[name] = s
		return true
	}

	return false
}

func (g *Game) setErrorPage(message string) {
	g.ctx = &Context{
		NextScene: "error",
		Message:   message,
	}
}
