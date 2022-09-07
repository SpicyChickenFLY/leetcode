class Solution:
    # Runtime - 93.78%, Memory Usage - 100.00%
    def findKthLargest(self, nums: List[int], k: int) -> int:
        nums = sorted(nums)
        return nums[-k]