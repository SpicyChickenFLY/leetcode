class Solution:
    def expander(self, s, left, right):
        if left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
            while left >= 0 and right < len(s) and s[left] == s[right]:
                left -= 1
                right += 1
            return s[left + 1: right]
        else:
            return ""


    def longestPalindrome(self, s: str) -> str:
        result = ""
        if len(s) <= 0:
            return result
        elif len(s) == 1:
            return s
        else:
            for i in range(len(s) - 1):
                even_expander = self.expander(s, i, i + 1)
                if len(even_expander) > len(result):
                    result = even_expander
            for i in range(1, len(s)):
                odd_expander = self.expander(s, i - 1, i + 1)
                if len(odd_expander) > len(result):
                    result = odd_expander
            if result == "":
                return s[0]
            else:
                return result

if __name__ == "__main__":
    s = Solution()
    print(s.longestPalindrome("abbc"))