func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	var dfs func(r, c , index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[index] {
			return false
		}
		
		temp := board[r][c]
		board[r][c] = '#'

		found := dfs(r, c + 1, index + 1) || 
				dfs(r, c - 1, index + 1) ||
				dfs(r + 1, c , index + 1) ||
				dfs(r - 1, c , index + 1)


		board[r][c] = temp
		return found		
	}

	for r := 0; r < rows ; r++ {
		for c := 0; c < cols ; c++ {
			if board[r][c] == word[0] {
				if dfs(r, c, 0) {
					return true
				}
			}
		}
	}

	return false
}
