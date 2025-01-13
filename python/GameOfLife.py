class Solution:
    def count(self, board, y, x) -> int:
        count = 0
        for i in range(0 if y - 1 < 0 else y - 1, y + 2):
            for j in range(0 if x - 1 < 0 else x - 1, x + 2):
                if i != y or j != x:
                    try:
                        if board[i][j] == 1 or board[i][j] == 3:
                            count += 1
                    except:
                        pass
        return count

    def gameOfLife(self, board) -> None:
        # Runtime: 91.07% MemoryUsage: 76.25%
        """ 
        state: 
            0: die  -> die
            1: live -> live 
            2: die  -> live
            3: live -> die
        """
        for y in range(len(board)):
            for x in range(len(board[0])):
                if board[y][x] == 1:
                    count = self.count(board, y, x)
                    if count > 3 or count < 2:
                        board[y][x] = 3
                if board[y][x] == 0 and self.count(board, y, x) == 3:
                    board[y][x] = 2
        for y in range(len(board)):
            for x in range(len(board[0])):
                if board[y][x] == 3:
                    board[y][x] = 0
                if board[y][x] == 2:
                    board[y][x] = 1
        print(board)

if __name__ == "__main__":
    s = Solution()
    test_cases= [
        [
            [0, 1, 0],
            [0, 0, 1],
            [1, 1, 1],
            [0, 0, 0]
        ]
    ]
    for test_case in test_cases:
        s.gameOfLife(test_case)