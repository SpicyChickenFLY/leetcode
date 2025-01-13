class Solution:
    def isValidSudoku(self, board: 'List[List[str]]') -> 'bool':
        for x, row in enumerate(range(len(board))):
            for y, col in enumerate(range(len(board[0]))):
                if board[row][col] == ".":
                    continue
                '''check row'''
                count = 0 
                for num in board[row]:
                    
                    if num == board[row][col]:
                        count += 1
                if count >= 2:
                    #print("row false")
                    return False
                '''check row'''
                count = 0 
                for num in [x[col] for x in board]:
                    if num == board[row][col]:
                        count += 1
                if count >= 2:
                    #print("col false")
                    return False
                '''check sub'''
                count = 0
                for sub_row in range(int(x / 3) * 3, int(x / 3) * 3 + 3):
                    for sub_col in range(int(y / 3) * 3, int(y / 3) * 3 + 3):
                        if board[sub_row][sub_col] == board[row][col]:
                            count += 1
                if count >= 2:
                    #print("sub false")
                    return False
        return True