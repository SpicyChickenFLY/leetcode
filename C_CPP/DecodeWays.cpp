#include <iostream>
#include <vector>
#include <cmath>
#include <string>
using namespace std;

class Solution
{
public:
    int numDecodings(string s)
    {
        int sta = 1, dyn = 0, temp = 0;
        char last = '0';
        for (int i = 0; i < s.length(); i++)
        {
            temp = sta + dyn;
            sta = dyn;
            dyn = temp;
            if (s[i] == '0')
            {
                if (last > '2')
                    return 0;
                dyn = 0;
            }
            else if ((s[i] > '6' && last == '2') || last > '2')
                sta = 0;
            else if (s[i] > '2')
            {
                sta += dyn;
                dyn = 0;
            }
            last = s[i];
            //cout << "sta-" << sta << " dyn-" << dyn << " ";
        }
        
        return sta + dyn;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    string input[] = {
        "222622",
        "2342",
        "101020",
        "27",
        "2839",
        "230",
        "4960"
    };
    for (int i = 0; i < 7; i++)
        cout << input[i] << ": " << s->numDecodings(input[i]) << endl;
    return 0;
}
