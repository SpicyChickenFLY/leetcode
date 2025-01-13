package main

func searchMatrix_v1(matrix [][]int, target int) bool {
    for row := len(matrix)-1; row >= 0 ; row-- {
        if target < matrix[row][0] {
            continue
        }
        for col := len(matrix[0]) - 1; col >= 0; col-- {
            if target == matrix[row][col] {
                return true
            }
        }

    }
    return false
}

// 叶师傅,切他二分法
func searchMatrix_v2(matrix [][]int, target int) bool {
    rs := 0
    re := len(matrix)
    cs := 0
    ce := len(matrix[0])
    // 先定位行
    for rs < re - 1 {
        rm := ( rs + re ) / 2
        if matrix[rm][0] == target {
            return true
        }
        if matrix[rm][0] > target {
            re = rm
        } else {
            rs = rm
        }
    }
    // 在定位列
    for cs < ce - 1 {
        cm := ( cs + ce ) / 2
        if matrix[rs][cm] == target {
            return true
        }
        if matrix[rs][cm] > target {
            ce = cm
        } else {
            cs = cm
        }
    }
    return matrix[rs][cs] == target
}


// 叶师傅,切他二叉搜索树
func searchMatrix(matrix [][]int, target int) bool {
    // 将数组看作搜索二叉树
    // 右上角为树根
    // 同一行为节点的左子树杈
    // 下面的行为节点的右子树杈,最右节点为子树的根
    row := 0
    col := len(matrix[0]) - 1

    for (row < len(matrix) && col >= 0) {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            row++
        } else {
            col--
        }
    }
    return false
}


type TestCase struct {
    matrix [][]int
    target int
    expect bool
}

func main() {
    testCases := []TestCase {
        { [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}} , 3, true },
        { [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}} , 11, true },
        { [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}} , 13, false },
        { [][]int{{1}} , 13, false },
        { [][]int{{1}} , 1, true },
    }
    for _, testCase := range testCases {
        println(searchMatrix(testCase.matrix, testCase.target), testCase.expect)
    }
}
