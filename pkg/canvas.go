package canvas

import "github.com/nsf/termbox-go"

func drawMessage(x, y int, msg string) {
	for pos, r := range msg {
		termbox.SetCell(x, y+pos, r, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawBox() {

	message := "hello world"
	drawMessage(0, 0, message)
}
func DrawCanvas() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawBox()

	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				break loop
			}
		}
	}
}
