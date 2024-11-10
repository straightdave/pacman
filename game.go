package main

import (
	"log"

	g "github.com/straightdave/pacman/game"
	_ "github.com/straightdave/pacman/scene/error"
	_ "github.com/straightdave/pacman/scene/intro"
	_ "github.com/straightdave/pacman/scene/play"
	_ "github.com/straightdave/pacman/scene/win"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ctx       *g.Context
	currScene g.Scene
	scenes    map[string]g.Scene
}

func NewGame() *Game {
	game := &Game{
		ctx:    &g.Context{},
		scenes: make(map[string]g.Scene),
	}

	if err := game.loadScene("error"); err != nil {
		log.Fatal(err)
	}

	if err := game.loadScene("intro"); err != nil {
		log.Fatal(err)
	}
	game.ctx.NextScene = "intro"

	return game
}

func (game *Game) Update() error {
	game.changeScene()

	for _, s := range game.scenes {
		if !s.IsActive() {
			continue
		}

		if err := s.Update(game.ctx); err != nil {
			return err
		}
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	if game.currScene != nil {
		game.currScene.Draw(game.ctx, screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 320
}

func (game *Game) changeScene() {
	if s, ok := game.scenes[game.ctx.NextScene]; ok {
		game.currScene = s
		return
	}

	if err := game.loadScene(game.ctx.NextScene); err != nil {
		game.setErrorPage(err)
	}
}

func (game *Game) loadScene(name string) error {
	if _, ok := game.scenes[name]; ok {
		return nil
	}

	if s, err := g.NewScene(name); err != nil {
		return err
	} else {
		game.scenes[name] = s
		return nil
	}
}

func (game *Game) setErrorPage(err error) {
	game.ctx = &g.Context{
		NextScene: "error",
		Message:   err.Error(),
	}
}
