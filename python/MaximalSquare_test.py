from typing import List

class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        result = int(matrix[0][0])

        vertical = [[0] * len(matrix[0]) for _ in range(len(matrix))]
        horizontal = [[0] * len(matrix[0]) for _ in range(len(matrix))]
        temp = [[0] * len(matrix[0]) for _ in range(len(matrix))]
        for i in range(len(matrix)):
            val =  int(matrix[i][0])
            horizontal[i][0] = val
            temp[i][0] = val
            if val > result:
                result = val
        for j in range(len(matrix[0])):
            val =  int(matrix[0][j])
            vertical[0][j] = val
            temp[0][j] = val
            if val > result:
                result = val
        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[0])):
                val =  int(matrix[i][j])
                vertical[i][j] = val * (vertical[i-1][j] + 1)
                horizontal[i][j] = val * (horizontal[i][j-1] + 1)
                min_val = min(vertical[i][j], horizontal[i][j])
                temp[i][j] = 0 if min_val == 0 else min(temp[i-1][j-1] + 1, min_val)
                if temp[i][j] > result:
                    result = temp[i][j]
        # print(vertical)
        # print(horizontal)
        # print(temp)

        return result * result

    def __init__(self) -> None:
        pass

def test_all():
    s = Solution()
    test_cases = [
            {"input": [[["1","0","1","0","0"], ["1","0","1","1","1"], ["1","1","1","1","1"], ["1","0","0","1","0"]]], "expect": 4},
            {"input": [[["1","0"], ["0","1"]]], "expect": 1},
            {"input": [[["0","1"], ["1","0"]]], "expect": 1},
            {"input": [[["0"]]], "expect": 0},
    ]
    for case in test_cases:
        print(case["input"])
        assert s.maximalSquare(*case["input"]) == case["expect"]
        print("\n")


if __name__ == '__main__':
    test_all()

