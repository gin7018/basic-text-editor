package store

type ArrayStore struct {
	Cursor []int
	Store  [][]string
}

type EditorActions interface {
	insert(text string)
	delete(row int, col int)
	up(row int, col int)
	down(row int, col int)
	left(row int, col int)
	right(row int, col int)
}

const DIM_COL int = 400

func (editor *ArrayStore) Insert(text string, new_line bool) {
	var row, col = editor.Cursor[0], editor.Cursor[1]
	if col == DIM_COL || new_line {
		//	add a new row since were at the end of the line
		var new_row = []string{text}
		//fmt.Printf(" address of store before adding row %p", &editor.Store)

		editor.Store = append(editor.Store, new_row)
		// todo if enter pressed and cursor is mid sentence, move rest of sentence to new line
		//fmt.Printf("after %p ", &editor.Store)

		row += 1
		col = 0
	} else {
		if col == len(editor.Store[row])-1 {
			editor.Store[row] = append(editor.Store[row], text)
		} else {
			editor.Store[row] = append(editor.Store[row][:col+1], editor.Store[row][col:]...)
			editor.Store[row][col+1] = text
		}
		col += 1
	}

	editor.Cursor[0] = row
	editor.Cursor[1] = col
}

func (editor *ArrayStore) Delete() {
	var row, col = editor.Cursor[0], editor.Cursor[1]
	var temp = editor.Store[row]
	editor.Store[row] = append(temp[:col], temp[col+1:]...)

	//if len(editor.Store[row]) == 0 {
	//	editor.Store = append(editor.Store[:row])
	//
	//} // todo revisit this logic

	editor.Cursor[0] = row
	editor.Cursor[1] = col - 1
}

func (editor *ArrayStore) Up() {
	var row = editor.Cursor[0]
	if row == 0 {
		return
	}
	var row_above_size = len(editor.Store[row-1])
	var current_row_size = len(editor.Store[row])
	if current_row_size > row_above_size {
		editor.Cursor[0], editor.Cursor[1] = row-1, row_above_size-1
	} else {
		editor.Cursor[0] = row - 1
	}
}

func (editor *ArrayStore) Down() {
	var row = editor.Cursor[0]
	var current_editor_size = len(editor.Store)
	if row == current_editor_size {
		return
	}
	var row_below_size = len(editor.Store[row+1])
	var current_row_size = len(editor.Store[row])
	if current_row_size > row_below_size {
		editor.Cursor[0], editor.Cursor[1] = row+1, row_below_size-1
	} else {
		editor.Cursor[0] = row + 1
	}
}

// todo just wrapping if you reach the end
func (editor *ArrayStore) Left() {
	var row, col = editor.Cursor[0], editor.Cursor[1]
	// two special case; first column
	if row == 0 && col == 0 {
		return
	}

	// wrap effect (wrap up)
	if col == 0 {
		row -= 1
		col = len(editor.Store[row]) - 1
		editor.Cursor[0], editor.Cursor[1] = row, col
	} else {
		editor.Cursor[1] = col - 1
	}

}

func (editor *ArrayStore) Right() {
	var row, col = editor.Cursor[0], editor.Cursor[1]
	// similar special case as left cmd
	var current_editor_size = len(editor.Store)
	var last_line_width = len(editor.Store[current_editor_size-1])
	if row == current_editor_size && col == last_line_width {
		return
	}

	// wrap effect
	if col == DIM_COL-1 {
		col = 0
		row += 1
		editor.Cursor[0], editor.Cursor[1] = row, col
	} else {
		editor.Cursor[1] += 1
	}
}

// todo undo/redo commands. Store each action as an event
// in chronological order (linked list) event (action : insert/delete, text, location)
