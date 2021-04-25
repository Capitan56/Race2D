package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type fon struct { //Структура трассы (фона)
	tex *sdl.Texture
}

func newFON(renderer *sdl.Renderer) (f fon, err error) {
	img, err := sdl.LoadBMP("sprites/Track.bmp") //Загрузка картинки Трассы (в bmp формате)
	if err != nil {
		return fon{}, fmt.Errorf("Загрузка спрайта Трассы: %v", err)

	}
	defer img.Free()
	f.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return fon{}, fmt.Errorf("Создание текстуры Трассы: %v", err)
	}

	return f, nil
}

func (f *fon) draw(renderer *sdl.Renderer) {

	renderer.Copy(f.tex,
		&sdl.Rect{X: 0, Y: 0, W: 1200, H: 500},
		&sdl.Rect{X: 0, Y: 0, W: 1200, H: 500})
}
