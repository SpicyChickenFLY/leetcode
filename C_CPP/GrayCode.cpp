#include <iostream>
#include <vector>
#include <cmath>
using namespace std;

class Solution
{
public:
    vector<int> grayCode(int n)
    {
        vector<int> result;
        result.push_back(0);
        int pow = 1;
        for (int i = 1; i <= n; i++)
        {
            for (int j = result.size() - 1; j >= 0; j--)
                result.push_back(result.at(j) + pow);
            pow <<= 1;
        } 
        return result;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    for (int i = 1; i < 4; i++)
    {
        vector<int> result = s->grayCode(i);
        cout << "i: " << i << endl;
        for (int i = 0; i < result.size(); i++)
        {
            cout << result.at(i) << endl;
        }
        cout << endl;
    }
    return 0;
}
