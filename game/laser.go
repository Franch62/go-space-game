package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite
	bound := image.Bounds()
	halfW := float64(bound.Dx()) / 2 //metade da largura da imagem do laser
	halfH := float64(bound.Dy()) / 2 //metade da altura da imagem do laser

	position.X -= halfW
	position.Y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0
	l.position.Y += -speed

}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.image, op)

}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()
	return NewRect(l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}
