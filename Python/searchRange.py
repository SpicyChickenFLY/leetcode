class Solution:
    def search_left(self, nums, target, left, right):
        middle = int((right + left) / 2)
        if right > left:
            if nums[middle] < target:
                #print(middle, nums[middle], "->")
                return self.search_left(nums, target, middle + 1, right)
            elif nums[middle] > target:
                #print(middle, nums[middle], "<-")
                return self.search_left(nums, target, left, middle - 1)
            else:
                #print(middle, nums[middle], "== <-")
                temp = self.search_left(nums, target, left, middle - 1)
                return temp if temp != -1 else middle
        elif right == left:
            if nums[left] == target:
                return left
            else:
                return -1
        else:
            return -1
    def search_right(self, nums, target, left, right):
        middle = int((right + left) / 2)
        if right > left:
            if nums[middle] < target:
                #print(nums[middle], "->")
                return self.search_right(nums, target, middle + 1, right)
            elif nums[middle] > target:
                #print(nums[middle], "<-")
                return self.search_right(nums, target, left, middle - 1)
            else:
                #print(nums[middle], "== ->")
                temp = self.search_right(nums, target, middle + 1, right)
                return temp if temp != -1 else middle
        elif right == left:
            if nums[left] == target:
                return left
            else:
                return -1
        else:
            return -1

    def searchRange(self, nums, target):
        return [
            self.search_left(nums, target, 0, len(nums) - 1), 
            self.search_right(nums, target, 0, len(nums) - 1)
        ]

if __name__ == "__main__":
    s = Solution()
    nums = [5,7,7,8,8,10]
    target = 5
    print(nums)
    print(s.searchRange(nums, target))