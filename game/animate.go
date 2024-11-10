package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animate struct {
	imgRef          *ebiten.Image
	framesInSprites []image.Rectangle
	ticks           int
	ticksPerFrame   int
}

func (a *Animate) AnimeTick()      { a.ticks++ }
func (a *Animate) AnimeTickReset() { a.ticks = 0 }

func (a *Animate) ImageToDraw() *ebiten.Image {
	i := (a.ticks / a.ticksPerFrame) % len(a.framesInSprites)
	return a.imgRef.SubImage(a.framesInSprites[i]).(*ebiten.Image)
}
