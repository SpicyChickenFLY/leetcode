#include <map>
#include <vector>
#include <cstring>

class Solution
{
public:
    vector<string> findRepeatedDnaSequences(string s)
    {
        map<string, int> mp;
        vector<string> result;
        if (s.length() < 10)
            return result;
        for (int i = 0; i <= s.length() - 10; i++)
        {
            mp[s.substr(i, 10)]++;
        }

        for (auto it : mp)
        {
            if (it.second > 1)
                result.push_back(it.first);
        }

        return result;
    }
};