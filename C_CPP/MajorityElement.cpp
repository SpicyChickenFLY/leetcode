#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    // RunTime: 25.88%, MemoryUsage: 96.97%
    int majorityElement(vector<int>& nums) {
        int size = nums.size();
        if (size == 1)
            return nums.at(0);
            
        map<int, int> count;
        for (vector<int>::iterator it = nums.begin(); it != nums.end(); it++)
        {
            if (count.find(*it) == count.end())
                count[*it] = 1;
            else
                count[*it] += 1;

            if (count[*it] > (size / 2))
                return *it;    
        }
        return 0;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    int test_arr1[] = {6, 5, 5};
    vector<int> test_case(test_arr1, test_arr1 + sizeof(test_arr1) / sizeof(int));
    cout << s->majorityElement(test_case) << endl;
    int test_arr2[] = {2, 2, 1, 1, 1, 2, 2};
    test_case = vector<int>(test_arr2, test_arr2 + sizeof(test_arr2) / sizeof(int));
    cout << s->majorityElement(test_case) << endl;
    return 0;
}
