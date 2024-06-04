package store

type Array struct {
	cursor [2]int
	store  [][]string
}

var array_text_editor Array

const DIM_COL int = 400

func insert(row int, col int, text string) {
	if col == DIM_COL {
		//	add a new row since were at the end of the line
		var new_row = []string{""}
		array_text_editor.store = append(array_text_editor.store, new_row)

		row += 1
		col = 0
	}
	array_text_editor.store[row][col] = text

}

func delete(row int, col int) {
	var temp = array_text_editor.store[row]
	array_text_editor.store[row] = append(temp[:col], temp[col+1:]...)

	if len(array_text_editor.store[row]) == 0 {
		array_text_editor.store = append(array_text_editor.store[:row])
	}
}

func up(row int) {
	if row == 0 {
		return
	}
	var row_above_size = len(array_text_editor.store[row-1])
	var current_row_size = len(array_text_editor.store[row])
	if current_row_size > row_above_size {
		array_text_editor.cursor = [2]int{row - 1, row_above_size - 1}
	} else {
		array_text_editor.cursor[0] = row - 1
	}
}

func down(row int) {
	var current_editor_size = len(array_text_editor.store)
	if row == current_editor_size {
		return
	}
	var row_below_size = len(array_text_editor.store[row+1])
	var current_row_size = len(array_text_editor.store[row])
	if current_row_size > row_below_size {
		array_text_editor.cursor = [2]int{row + 1, row_below_size - 1}
	} else {
		array_text_editor.cursor[0] = row + 1
	}
}

// todo just wrapping if you reach the end
func left(row int, col int) {
	// two special case; first column
	if row == 0 && col == 0 {
		return
	}

	// wrap effect (wrap up)
	if col == 0 {
		row -= 1
		col = len(array_text_editor.store[row]) - 1
		array_text_editor.cursor = [2]int{row, col}
	} else {
		array_text_editor.cursor[1] -= 1
	}

}

func right(row int, col int) {
	// similar special case as left cmd
	var current_editor_size = len(array_text_editor.store)
	var last_line_width = len(array_text_editor.store[current_editor_size-1])
	if row == current_editor_size && col == last_line_width {
		return
	}

	// wrap effect
	if col == DIM_COL-1 {
		col = 0
		row += 1
		array_text_editor.cursor = [2]int{row, col}
	} else {
		array_text_editor.cursor[1] += 1
	}
}

// todo undo/redo commands. store each action as an event
// in chronological order (linked list) event (action : insert/delete, text, location)
