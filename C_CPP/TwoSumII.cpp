#include <map>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    // RunTime: 56.68%, MemoryUsage: 41.14%
    vector<int> twoSum(vector<int>& numbers, int target) 
    {
        map<int, int> numberMap;
        vector<int> result;
        int i = 0;
        for (vector<int>::iterator it = numbers.begin(); it != numbers.end(); it++)
        {
            int sub = target - *it;
            if (numberMap.find(sub) != numberMap.end()) 
            {
                result.push_back(numberMap[sub] + 1);
                result.push_back(i + 1);
                return result;
            }
            else
            {
                numberMap[*it] = i;
            }
            i++;
        }
    }
};

int main(int argc, char const *argv[])
{
    /* code */
    Solution *s = new Solution();
    int temp[] = {2, 7, 11, 15};
    vector<int> numbers(temp, temp + sizeof(temp) / sizeof(int));
    int target = 9;
    vector<int> result = s->twoSum(numbers, target);
    for (vector<int>::iterator it = result.begin(); it != result.end(); it++)
    {
        cout << *it << " ";
    }
    return 0;
}