class Solution:

    def recursion(self, nums, lb, rb):
        if rb == lb:
            return nums[lb]
        else:
            if str(nums[lb: rb + 1]) in self.record: 
                result = self.record[str(nums[lb: rb + 1])]
            else:
                mid = int((lb + rb) / 2)
                '''left'''
                left = self.recursion(nums, lb, mid)
                '''right'''
                right = self.recursion(nums, mid + 1, rb)
                '''mid-to-left'''
                l_max = -999
                l_sum = 0
                for l_index in range(mid, -1, -1):
                    l_sum += nums[l_index]
                    if l_sum > l_max:
                        l_max = l_sum
                r_max = -999
                r_sum = 0
                for r_index in range(mid + 1, len(nums)):
                    r_sum += nums[r_index]
                    if r_sum > r_max:
                        r_max = r_sum
                result = max(left, right, l_max + r_max)
                self.record[str(nums[lb: rb + 1])] = result 
            return result

    def maxSubArray(self, nums):
        for i in range(1, len(nums)):
            if nums[i-1] > 0:
                nums[i] += nums[i-1]
        return max(nums)



if __name__ == "__main__":
    s = Solution()
    inputs = [
        [1, -1 , 0],
        [-2, 1, -3, 4, -1, 2, 1, -5, 4],
        [-2,-1],
        [-2, -2, -3, 0, -3, 1, -3]
        
    ]
    for test_input in inputs:
        print(s.maxSubArray(test_input))