class Solution:
    record = {}

    def recursion(self, candidates, target):
        result = []
        for index, candidate in enumerate(candidates):
            if target - candidate > 0:
                sub_results = self.recursion(candidates[index + 1: ], target - candidate)
                if len(sub_results) > 0:
                    for index in range(len(sub_results)):
                        sub_results[index].insert(0, candidate)
                        result.append(sub_results[index])
            elif target - candidate == 0:
                result.append([target])
                break
            else:
                break
        return result


    def combinationSum2(self, candidates, target):
        self.record = {}
        candidates.sort()
        results = self.recursion(candidates, target)
        new_results = []
        for result in results:
            if result not in new_results:
                new_results.append(result)
        return new_results

if __name__ == "__main__":
    s = Solution()
    inputs = [
        [[10,1,2,7,6,1,5], 8]
    ]
    for i in inputs:
        print(s.combinationSum2(i[0], i[1]))