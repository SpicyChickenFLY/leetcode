class Solution:
    def merge(self, intervals):
        if len(intervals) == 0:
            return intervals
        else:
            intervals.sort(key=lambda x: x[1])
            intervals.sort(key=lambda x: x[0])
            i = 0
            max_index = len(intervals)
            while i < max_index:
                if i + 1 >= len(intervals):
                    break
                if intervals[i+1][0] <= intervals[i][1]:
                    max_value = intervals[i][1]
                    for j in range(i + 1, len(intervals)):
                        if intervals[j][0] > intervals[i][1]:
                            intervals = intervals[: i] + [[intervals[i][0], max_value]] + intervals[j:]
                            break
                        max_value = max(intervals[j][1], max_value)
                    else:
                        intervals = intervals[: i] + [[intervals[i][0], max_value]]
                else:
                    i += 1
        return intervals
                    

if __name__ == "__main__":
    s = Solution()
    inputs = [
        [[1, 3], [2, 6], [8, 10], [15, 18]], 
        [[1, 4], [4, 5]], 
        [[1, 4], [0, 4]], 
        [[1, 4], [2, 3]], 
        [[1, 4], [0, 2], [3, 5]], 
        [[5, 5], [1, 3], [3, 5], [4, 6], [1, 1], [3, 3], [5, 6], [3, 3], [2, 4], [0, 0]], 
        [[1, 3], [0, 2], [2, 3], [4, 6], [4, 5], [5, 5], [0, 2], [3, 3]], 
        [[1, 3], [2, 6], [8, 10], [15, 18]]
    ]

    ground_truth = [
        [[1, 6], [8, 10], [15, 18]],
        [[1, 5]],
        [[0, 4]],
        [[1, 4]],
        [[0, 5]],
        [[0, 0], [1, 6]],
        [[0, 3], [4, 6]],
        [[1, 6], [8, 10], [15, 18]]
    ]

    for index in range(len(inputs)):
        if s.merge(inputs[index]) != ground_truth[index]:
            print('Error: ', inputs[index])
            break
    else:
        print('Pass')