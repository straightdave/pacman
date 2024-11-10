package game

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type sceneLoader func() (Scene, error)

type Scene interface {
	IsActive() bool
	Activate()
	Deactivate()
	Draw(*Context, *ebiten.Image)
	Update(*Context) error
}

var (
	m                sync.Mutex
	registeredScenes map[string]sceneLoader
)

func init() {
	registeredScenes = make(map[string]sceneLoader)
}

func RegisterScene(name string, loader sceneLoader) {
	m.Lock()
	registeredScenes[name] = loader
	log.Printf("scene %s registered!", name)
	m.Unlock()
}

func NewScene(name string) (Scene, error) {
	if loader, ok := registeredScenes[name]; !ok {
		return nil, NewGameError(name, "no such scene")
	} else {
		return loader()
	}
}
