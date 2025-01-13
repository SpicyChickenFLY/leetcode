#include <iostream>
#include <vector>
#include <set>
#include <unordered_map>
#include <string>
using namespace std;

class Solution
{
public:
    int searchInsert(vector<int> &nums, int target)
    {
        int low = 0, high = nums.size() - 1;
        if (target < nums.at(low))
            return 0;
        if (target > nums.at(high))
            return nums.size();
        while (low + 1 < high) {
            int mid = (low + high) / 2;
            if (nums.at(mid) < target)
                low = mid;
            else if (nums.at(mid) > target)
                high = mid;
            else
                return mid;
        }
        if (target > nums.at(low))
            return high;
        else
            return low;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution;
    /* int array[] = {1, 3, 5, 6};
    vector<int> nums(array, array + sizeof(array) / sizeof(int));
    int target[] = {0, 2, 5, 7}; */
    int array[] = {1};
    vector<int> nums(array, array + sizeof(array) / sizeof(int));
    int target[] = {1};
    for (int t = 0; t < sizeof(target) / sizeof(int); t++)
    {
        cout << "target:" << target[t] << " result:" << s->searchInsert(nums, target[t]) << endl;
    }
    return 0;
}
