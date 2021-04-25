package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	Winner1Speed = 0.25 //Скорость передвижения надписи "Победа"
	Winner1Size  = 100  //Размер надписи "Победа"
)

type winner1 struct { //Структура надписи "Победа"
	tex  *sdl.Texture
	x, y float64 //Координаты для перемещения надписи "Победа"
}

func newWinner1(renderer *sdl.Renderer) (w winner1, err error) {
	img, err := sdl.LoadBMP("sprites/Winner1.bmp") //Загрузка картинки надписи "Победа" (в bmp формате) (синей машинки)
	if err != nil {
		return winner1{}, fmt.Errorf("Загрузка спрайта игрока №2: %v", err)

	}
	defer img.Free()
	w.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return winner1{}, fmt.Errorf("Создание текстуры игрока №2: %v", err)
	}

	w.x = Winner1Size - 1200
	w.y = screenHeight/2.0 - 50

	return w, nil
}

func (w *winner1) draw(renderer *sdl.Renderer) {
	//Конвертируем координаты надписи "Победа" за переделы окна
	x := w.x - Winner1Size/2.0
	y := w.y - Winner1Size/2.0
	renderer.Copy(w.tex,
		&sdl.Rect{X: 0, Y: 0, W: 600, H: 227},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 600, H: 227}) //Координаты перемещения и размеры надписи "Победа"
}

func (w *winner1) update() {
	rand.Seed(time.Now().UTC().UnixNano())
	keys := sdl.GetKeyboardState() //Привязка клавиш для управления движением надписи "Победа"
	if keys[sdl.SCANCODE_LEFT] == 1 {
		w.x -= player1Speed //Движение надписи "Победа" влево
	} else if keys[sdl.SCANCODE_1] == 1 {
		gaz1 := rand.Intn(2)
		if gaz1 == 0 {
			w.x += Winner1Speed - 0.04 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_2] == 1 {
		gaz2 := rand.Intn(2)
		if gaz2 == 0 {
			w.x += Winner1Speed - 0.1 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_3] == 1 {
		gaz3 := rand.Intn(2)
		if gaz3 == 0 {
			w.x += Winner1Speed - 0.09 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_4] == 1 {
		gaz4 := rand.Intn(2)
		if gaz4 == 0 {
			w.x += Winner1Speed - 0.15 //Движение надписи "Победа" вправо
		}
	} else if keys[sdl.SCANCODE_5] == 1 {
		gaz5 := rand.Intn(2)
		if gaz5 == 0 {
			w.x += Winner1Speed - 0.06 //Движение надписи "Победа" вправо
		}
	}

}
