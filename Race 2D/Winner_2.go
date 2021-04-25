package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	Winner2Speed = 0.25 //Скорость передвижения надписи "Победа"
	Winner2Size  = 100  //Размер надписи "Победа"
)

type winner2 struct { //Структура надписи "Победа"
	tex  *sdl.Texture
	x, y float64 //Координаты для перемещения надписи "Победа"
}

func newWinner2(renderer *sdl.Renderer) (w winner2, err error) {
	img, err := sdl.LoadBMP("sprites/Winner2.bmp") //Загрузка картинки надписи "Победа" (красной машинки)
	if err != nil {
		return winner2{}, fmt.Errorf("Загрузка спрайта игрока №2: %v", err)

	}
	defer img.Free()
	w.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return winner2{}, fmt.Errorf("Создание текстуры игрока №2: %v", err)
	}

	w.x = Winner2Size - 1196
	w.y = screenHeight/2.0 + 150

	return w, nil
}

func (w *winner2) draw(renderer *sdl.Renderer) {
	//Конвертируем координаты надписи "Победа" за пределы окна
	x := w.x - Winner2Size/2.0
	y := w.y - Winner2Size/2.0
	renderer.Copy(w.tex,
		&sdl.Rect{X: 0, Y: 0, W: 600, H: 227},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 600, H: 227}) //Координаты перемещения и размеры надписи "Победа"
}

func (w *winner2) update() {
	rand.Seed(time.Now().UTC().UnixNano())
	keys := sdl.GetKeyboardState() //Привязка клавиш для управления движением надписи "Победа"
	if keys[sdl.SCANCODE_LEFT] == 1 {
		w.x -= Winner2Speed //Движение надписи "Победа" влево
	} else if keys[sdl.SCANCODE_G] == 1 {
		gaz11 := rand.Intn(2)
		if gaz11 == 1 {
			w.x += Winner2Speed - 0.06 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_H] == 1 {
		gaz22 := rand.Intn(2)
		if gaz22 == 1 {
			w.x += Winner2Speed - 0.1 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_J] == 1 {
		gaz33 := rand.Intn(2)
		if gaz33 == 1 {
			w.x += Winner2Speed - 0.15 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_K] == 1 {
		gaz44 := rand.Intn(2)
		if gaz44 == 1 {
			w.x += Winner2Speed - 0.04 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_L] == 1 {
		gaz55 := rand.Intn(2)
		if gaz55 == 1 {
			w.x += Winner2Speed - 0.09 //Движение надписи "Победа" вправо
		}
	}

}
