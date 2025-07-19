package game

import (
	"math/rand"
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]
	speed := 2.0

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Star) Update() {
	m.position.Y += m.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(s.position.X, s.position.Y)
	screen.DrawImage(s.image, op)
}
