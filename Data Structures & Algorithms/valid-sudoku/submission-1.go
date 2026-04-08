func isValidSudoku(board [][]byte) bool {
   row := make([]map[byte]bool, 9)
   col := make([]map[byte]bool, 9)
   square := make([]map[byte]bool, 9)

   for i:= 0 ; i < 9; i++ {
	row[i] = make(map[byte]bool)
	col[i] = make(map[byte]bool)
	square[i] = make(map[byte]bool)
   }


   for r:= 0 ; r < 9; r++ {
	for c:= 0 ; c < 9; c++ {

		if board[r][c] == '.' {
			continue
		}

		val := board[r][c]
		sqIdx := (r/3)*3 + c/3

		if square[sqIdx][val] || row[r][val] || col[c][val] {
			return false
		}

		square[sqIdx][val] = true
		row[r][val] = true 
		col[c][val] = true
   	 }
   }
   return true
}