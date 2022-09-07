class Solution:
    #Sort Solution - Runtime: 5.46%  Memory Usage: 5.21%
    def commonChars1(self, A):
        if len(A) == 0:
            return []
        result = []
        for char in A[0]:
            result.append(char)
        result.sort()
        for i in range(1, len(A)):
            p = 0
            q = 0
            new_list = []
            for char in A[i]:
                new_list.append(char)
            new_list.sort()
            while (p < len(result) and q < len(new_list)):
                if result[p] == new_list[q]:
                    p += 1
                    q += 1
                elif result[p] > new_list[q]:
                    q += 1
                else: 
                    result.remove(result[p])
            if q == len(new_list):
                result = result[:p]
        return result

    #Dict Solution - Runtime: 74.32%  Memory Usage: 5.21%
    def commonChars2(self, A):
        result = []
        count = {}
        for char in A[0]:
            count[char] = A[0].count(char)
        for word in A[1:]:
            for char in count:
                sub_count = word.count(char)
                if sub_count < count[char]:
                    count[char] = sub_count
                
        for key in count:
            for index in range(count[key]):
                result.append(key)
        return result

    # Solution of others - Runtime: 84.34%  Memory Usage: 5.21%
    def commonChars(self, A):
        first_word = A[0]
        for a in A[1:]:
            for s in first_word:
                if s not in a:
                    first_word = first_word.replace(s, '', 1)
                else:
                    a = a.replace(s, '', 1)
        return [s for s in first_word]

if __name__ == "__main__":
    test_cases = [
        ["bella", "label", "roller"],
        ["cool", "lock", "cook"]
    ]
    s = Solution()
    for test_case in test_cases:
        print(s.commonChars(test_case))
