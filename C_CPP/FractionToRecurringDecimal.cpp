#include <iostream>
#include <string>
#include <map>

using namespace std;

class Solution
{
public:
    // RunTime: 58.15%, MemoryUsage: 100.00%
    string fractionToDecimal(int numerator, int denominator)
    {
        if (numerator == 0)
            return "0";
        long long a = numerator, b = denominator;
        map<int, int> record;
        string former = to_string(abs(a / b));
        string latter = "";
        long long mod = a % b;
        if (mod != 0)
        {
            former += ".";
            while (mod != 0) 
            {
                mod *= 10;
                map<int, int>::iterator it = record.find(mod);
                if (it != record.end())
                {
                    latter += ")";
                    latter.insert(it->second, "(");
                    break;
                }
                else
                {
                    record.insert(pair<int, int>(mod, latter.length()));
                    while (abs(mod) < (b))
                    {
                        mod *= 10;
                        latter += "0";
                    }
                    int value = mod / b;
                    if (value < 0) value *= -1;
                    latter += to_string(value);
                    mod %= b;
                }
            }
        }
        if (numerator > 0 && denominator > 0 || numerator < 0 && denominator < 0)
            return former + latter;
        else
            return "-" + former + latter;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();

    int test_cases[][2] = {

        {-1, -2147483648}
    };

    for (int i = 0; i < 1; i++) 
    {
        cout << s->fractionToDecimal(test_cases[i][0], test_cases[i][1]) << endl;
    }

    return 0;
}
