#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    // RunTime: 75.83%, MemoryUsage: 80.00%
    int findMin(vector<int> &nums)
    {
        for (int i = 0; i < nums.size() - 1; i++)
        {
            if (nums.at(i) > nums.at(i + 1))
                return nums.at(i + 1);
        }
        return nums.at(0);
    }

};

int main(int argc, char const *argv[])
{
    /* code */
    return 0;
}
