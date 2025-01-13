class Solution:
    def spiralOrder(self, matrix):
        direction = 0 #right:0 down:1 left:2 up:3
        left = 0
        right = len(matrix[0])
        up = 0
        down = len(matrix)
        x = 0
        y = 0
        while left != right or up != down:
            if direction == 0:
                if x < right:
                    pass
                else:
                    pass

