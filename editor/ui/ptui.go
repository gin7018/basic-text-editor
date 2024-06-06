package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	"slices"
)
import "editor/store"

func clear_terminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	//fmt.Print("\033[1A\033[K")
}

func hide_builtin_cursor() {
	_, err := fmt.Fprintf(os.Stderr, "\033[?25l")
	if err != nil {
		fmt.Println("having trouble hiding the built in cursor")
		panic(err)
	}
}

func render(editor *store.ArrayStore) {
	for i := 0; i < len(editor.Store); i++ {
		fmt.Print("\n", "[", i+1, "] ")
		for j := 0; j < len(editor.Store[i]); j++ {
			fmt.Print(editor.Store[i][j])
			if slices.Equal(editor.Cursor, []int{i, j}) {
				fmt.Print("|")
			}
		}
	}
}

func main() {
	fmt.Println("welcome to the basic text editor (plain text ver)")
	fmt.Println("Press ESC to quit")
	fmt.Println("--------")
	hide_builtin_cursor()

	// todo UNDO command
	// todo REDO command (can only redo DIRECTLY after undo command)

	var editor = store.ArrayStore{Store: make([][]string, 1), Cursor: []int{0, 0}}
	editor.Store[0] = make([]string, 1)

	render(&editor)

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	arrow_key_map := map[keyboard.Key]func(){
		keyboard.KeyArrowUp:    editor.Up,
		keyboard.KeyArrowDown:  editor.Down,
		keyboard.KeyArrowLeft:  editor.Left,
		keyboard.KeyArrowRight: editor.Right,
	}

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeySpace {
			editor.Insert(" ", false)
		} else if key == keyboard.KeyBackspace {
			editor.Delete()
		} else if arrow_action := arrow_key_map[key]; arrow_action != nil {
			arrow_action()
		} else if key == keyboard.KeyCtrlZ {
			fmt.Println("here in cz")
			editor.Undo()
		} else {
			if key == keyboard.KeyEnter {
				editor.Insert("", true)
			} else {
				editor.Insert(string(char), false)
			}
		}

		clear_terminal()
		render(&editor)
		if key == keyboard.KeyEsc {
			break
		}
	}

}
