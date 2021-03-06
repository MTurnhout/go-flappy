package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Could not initialize TTF: %v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	_ = r

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("Couldn not draw title: %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/FiraCode-Bold.ttf", 20)
	if err != nil {
		return fmt.Errorf("Could not load font: %v", err)
	}
	defer f.Close()

	s, err := f.RenderUTF8_Solid("Go Flappy", sdl.Color{R: 255, G: 100, B: 0, A: 255})
	if err != nil {
		return fmt.Errorf("Could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v", err)
	}

	r.Present()

	return nil
}
