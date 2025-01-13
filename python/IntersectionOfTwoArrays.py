class Solution(object):
    # simple solution : 94.05%, 98.87%
    def intersection(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        result = []
        nums1.sort()
        nums2.sort()
        i = 0
        j = 0
        len_nums1 = len(nums1)
        len_nums2 = len(nums2)
        while (i < len_nums2 and j < len_nums1):
            if (i < len_nums2 - 1 and nums2[i] == nums2[i + 1]) or nums2[i] < nums1[j]:
                i += 1
                continue
            if (j < len_nums1 - 1 and nums1[j] == nums1[j + 1]) or nums2[i] > nums1[j]:
                j += 1
                continue
            result.append(nums1[j])
            i += 1
            j += 1
        return result

if __name__ == "__main__":
    inputs = [
        [[1, 2, 2, 1], [2, 2]],
        [[4, 9, 5], [9, 4, 9, 8, 4]]
    ]
    s = Solution()
    for i in inputs:
        print(i, ": ", s.intersection(i[0], i[1]))
    
