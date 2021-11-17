package back_track

import "fmt"

//问题：在n×n格的棋盘上放置彼此不受攻击的n个皇后。按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
//N皇后问题等价于在n×n格的棋盘上放置n个皇后，任何2个皇后不放在同一行或同一列或同一斜线上。
//
//分析：从n×n个格子中选择n个格子摆放皇后。可见解空间树为子集树。
//
//使用Board[N][N]来表示棋盘，Board[i][j]=0 表示(I,j)位置为空，Board[i][j]=1 表示(I,j)位置摆放有一个皇后。
//
//全局变量way表示总共的摆放方法数目。
//
//使用Queen(t)来摆放第t个皇后。Queen(t) 函数符合子集树时的递归回溯范式。当t>N时，说明所有皇后都已经摆   放完成，这是一个可行的摆放方法，输出结果；否则，遍历棋盘，找皇后t所有可行的摆放位置，Feasible(i,j) 判断皇后t能否摆放在位置(i,j)处，如果可以摆放则继续递归摆放皇后t+1，如果不能摆放，则判断下一个位置。
//
//Feasible(row,col)函数首先判断位置(row,col)是否合法，继而判断(row,col)处是否已有皇后，有则冲突，返回0，无则继续判断行、列、斜方向是否冲突。斜方向分为左上角、左下角、右上角、右下角四个方向，每次从（row,col）向四个方向延伸一个格子，判断是否冲突。如果所有方向都没有冲突，则返回1，表示此位置可以摆放一个皇后。
//————————————————
//版权声明：本文为CSDN博主「JarvisChu」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
//原文链接：https://blog.csdn.net/JarvisChu/article/details/16067319

//回溯问题关键点：
//1.解决路径问题很有效
//2.用递归可以使代码比较好理解，递归的层数是路径的层数
//3.回溯的的recover很重要，每一次路径选择完成后需要回复上前一状态，这样才能做到回溯

func QueensSolutions(max int, isPrintIncomplete bool) int {
	GoodSolutions, IncompleteSolutions, curI := new(int), new(int), new(int)
	*GoodSolutions = 0
	*IncompleteSolutions = 0
	*curI = 0
	board := initBoard(max)
	queen(1, GoodSolutions, IncompleteSolutions, curI, board, isPrintIncomplete)
	fmt.Printf("total goold solutions: %d\n", *GoodSolutions)
	if isPrintIncomplete {
		fmt.Printf("total Incomplete Solutions %d\n", *IncompleteSolutions)
	}
	return 0
}

func initBoard(max int) [][]int {
	board := make([][]int, max)
	for i := 0; i < max; i++ {
		board[i] = make([]int, max)
		for j := 0; j < max; j++ {
			board[i][j] = 0
		}
	}
	return board
}

func queen(num int, goodSolutions, incompleteSolutions, curI *int, board [][]int, isPrintIncomplete bool) {
	max := len(board)
	if num > max {
		*goodSolutions++
		printSolution(goodSolutions, board, true)
		return
	}
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if feasible(i, j, board, curI) {
				board[i][j] = 1
				*curI = i
				queen(num+1, goodSolutions, incompleteSolutions, curI, board, isPrintIncomplete)
				board[i][j] = 0 //recover board after each solution
				*curI = 0
			}
			if isPrintIncomplete && i == max-1 && j == max-1 {
				*incompleteSolutions++
				printSolution(incompleteSolutions, board, false)
			}
		}
	}
}

func printSolution(solutions *int, board [][]int, isGoodSolution bool) {
	if isGoodSolution {
		fmt.Println("good solution: ", *solutions)
	} else {
		fmt.Println("incomplete solution: ", *solutions)
	}
	max := len(board)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("================")
}

func feasible(i, j int, board [][]int, curI *int) bool {
	max := len(board)
	//to avoid repeat, only feasible for below slot
	if i < *curI {
		return false
	}
	//row and col no any queen
	for k := 0; k < max; k++ {
		if board[k][j] != 0 || board[i][k] != 0 {
			return false
		}
	}
	//left top line no any queen
	for r, c := i-1, j-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] != 0 {
			return false
		}
	}
	//left bottom line no any queen
	for r, c := i-1, j+1; r >= 0 && c < max; r, c = r-1, c+1 {
		if board[r][c] != 0 {
			return false
		}
	}
	//right top line no any queen
	for r, c := i+1, j-1; r < max && c >= 0; r, c = r+1, c-1 {
		if board[r][c] != 0 {
			return false
		}
	}
	//right bottom line no any queen
	for r, c := i+1, j+1; r < max && c < max; r, c = r+1, c+1 {
		if board[r][c] != 0 {
			return false
		}
	}
	return true
}
