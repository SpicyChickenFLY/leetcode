#include <iostream>
#include <string>
using namespace std;

class Solution
{
public:
    // RunTime: 100.00%, MemoryUsage: 64.31%
    int titleToNumber(string s)
    {
        int result = 0;
        for (int i = 0; i < s.length() - 1; i++)
            result = (result + (s.at(i) - 'A' + 1)) * 26;
        result += s.at(s.length() - 1) - 'A' + 1;
        return result;
    }
};

int main(int argc, char const *argv[])
{
    /* code */
    Solution *s = new Solution();
    cout << s->titleToNumber("A") << endl;
    cout << s->titleToNumber("Z") << endl;
    cout << s->titleToNumber("AA") << endl;
    cout << s->titleToNumber("AB") << endl;
    cout << s->titleToNumber("AZ") << endl;
    cout << s->titleToNumber("BA") << endl;
    cout << s->titleToNumber("ZA") << endl;
    return 0;
}