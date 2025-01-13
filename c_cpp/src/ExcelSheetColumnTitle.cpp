#include <iostream>
#include <string>
using namespace std;

class Solution
{
public:
    // RunTime: 100.00%, MemoryUsage: 100.00%
    string convertToTitle(int n)
    {
        string result;
        while (n > 0)
        {
            if (n % 26 == 0)
            {
                result = 'Z' + result;
                n /= 26;
                n -= 1;
                continue;
            }
            else
            {
                result = (char)('A' - 1 + n % 26) + result;
                n /= 26;
            }
        }
        return result;
    }
};

int main(int argc, char const *argv[])
{
    /* code */
    Solution *s = new Solution();
    cout << s->convertToTitle(1) << endl;
    cout << s->convertToTitle(26) << endl;
    cout << s->convertToTitle(27) << endl;
    cout << s->convertToTitle(28) << endl;
    cout << s->convertToTitle(677) << endl;
    cout << s->convertToTitle(701) << endl;
    cout << s->convertToTitle(702) << endl;
    return 0;
}