#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution
{
public:
    vector<string> restoreIpAddresses(string s)
    {
        vector<string> result;
        int length = s.length();
        if (length < 4 || length > 12)
            return result;
        for (int ss = 2; ss < length - 1; ss++)
        {
            vector<string> subResult1, subResult2;
            int fs = 0, ts = 0;
            if (ss > 6 || length - ss > 6)
                continue;
            if (ss == 6)
                if (s.substr(0, 3) > "255" || s.substr(0, 1) == "0" || s.substr(3, 3) > "255" || s.substr(3, 1) == "0")
                    continue;
                else
                {
                    fs = 3;
                    subResult1.push_back(
                        s.substr(0, fs) + "." +
                        s.substr(fs, ss - fs)
                    );
                }
            if (length - ss == 6)
                if (s.substr(ss, 3) > "255" || s.substr(ss, 1) == "0" || s.substr(ss + 3, 3) > "255" || s.substr(ss + 3, 1) == "0")
                    continue;
                else
                {
                    ts = ss + 3;
                    subResult2.push_back(s.substr(ss, ts - ss) + "." + s.substr(ts, length - ts));
                }
                    
            if (!fs)
            {
                for (int fs = 1; fs < ss; fs++)
                    if (
                        (fs > 1 && s.substr(0, 1) == "0") ||
                        (fs > 2 && s.substr(0, fs) > "255") ||
                        fs > 3 ||
                        (ss - fs > 1 && s.substr(fs, 1) == "0") ||
                        (ss - fs > 2 && s.substr(fs, ss - fs) > "255") ||
                        ss - fs > 3)
                        continue;
                    else
                        subResult1.push_back(s.substr(0, fs) + "." + s.substr(fs, ss - fs));
            }
            if (!ts)
            {
                for (int ts = ss + 1; ts < length; ts++)
                    if (
                        (ts - ss > 1 && s.substr(ss, 1) == "0") ||
                        (ts - ss > 2 && s.substr(ss, ts - ss) > "255") ||
                        ts - ss > 3 ||
                        (length - ts > 1 && s.substr(ts, 1) == "0") ||
                        (length - ts > 2 && s.substr(ts, length - ts) > "255") ||
                        length - ts > 3)
                        continue;
                    else
                        subResult2.push_back(s.substr(ss, ts - ss) + "." + s.substr(ts, length - ts));
            }
            for (int i = 0; i < subResult1.size(); i++)
                for (int j = 0; j < subResult2.size(); j++)
                    result.push_back(subResult1.at(i) +"." + subResult2.at(j));
        }
        return result;
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    string input[] = {
        // "010010",
        // "101020",
        // "27",
        // "25525511135",
        //"255255255255",
        "10101010"};
    for (int i = 0; i < 1; i++)
    {
        cout << input[i] << ": " << endl;
        vector<string> result = s->restoreIpAddresses(input[i]);
        for (int j = 0; j < result.size(); j++)
            cout << " " << result.at(j) << endl;
        cout << endl;
    }

    return 0;
}
