package main

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Name() string
	Draw(*Context, *ebiten.Image)
	Update(*Context) error
}
