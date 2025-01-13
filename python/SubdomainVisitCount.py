class Solution:
    # Runtime: 98.17%  Memory Usage: 22.14%
    def subdomainVisits(self, cpdomains):
        result = []
        counter = {}
        for cpdomain in cpdomains:
            [count, domain] = cpdomain.split(" ")
            subdomains = domain.split(".")
            current_subdomain = ""
            for subdomain in reversed(subdomains):
                current_subdomain = subdomain + "." + current_subdomain
                if current_subdomain not in counter:
                    counter[current_subdomain] = int(count)
                else:
                    counter[current_subdomain] += int(count)
        for key in counter:
            result.append(str(counter[key]) + " " + key[:-1])
        return result


if __name__ == "__main__":
    test_cases = [
        ["9001 discuss.leetcode.com"],
        ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
    ]
    s = Solution()
    for test_case in test_cases:
        result = s.subdomainVisits(test_case)
        print(result)
