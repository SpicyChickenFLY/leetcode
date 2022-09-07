 
# Use Dynamic Programming(DP) Algorithm to solve the problem
# The K will go only Right and Down, so there is only two choice
# in each step.

class Solution:
    def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
        
        directs = [(0,1), (1,0)]
        dest = (len(dungeon)-1, len(dungeon[0])-1)

        @lru_cache(None)
        def recursion(r,c):
            if (r,c) == dest:   
                return max(1,1-dungeon[r][c])
            
            ret = float('inf')
            for i,j in directs:
                x,y = i+r, c+j
                if 0<=x<len(dungeon) and 0<=y<len(dungeon[0]):
                    ret = min(ret, max(1, recursion(x,y) - dungeon[r][c]))
            return ret
        
        return recursion(0,0)


testCase = [
    [-2, -3, 3],
    [-5, -10, 1],
    [10, 30, -5]
]

s = Resolution()
print(s.calculateMinimumHP(testCase))