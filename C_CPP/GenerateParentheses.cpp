#include <iostream>
#include <vector>
#include <set>
#include <unordered_map>
#include <string>
using namespace std;

class Solution
{
private:
    unordered_map<int, vector<string>> record;

public:
    vector<string> recursion(int n)
    {
        vector<string> result;
        if (n == 1)
            result.push_back("()");
        else if (record.find(n) == record.end()) {
            vector<string> subResult = recursion(n - 1);
            set<string> temp;
            for (int i = 0; i < subResult.size(); i++)
            {
                temp.insert("()" + subResult.at(i));
                temp.insert("(" + subResult.at(i) + ")");
                temp.insert(subResult.at(i) + "()");
            }
            
            if (n >= 4) {
                for (int i = 2; i <= n / 2; i++) {
                    vector<string> subResult1 = recursion(i);
                    vector<string> subResult2 = recursion(n - i);
                    for (int j = 0; j < subResult1.size(); j++)
                        for (int k = 0; k < subResult2.size(); k++)
                        {
                            temp.insert(subResult1.at(j) + subResult2.at(k));
                            temp.insert(subResult2.at(k) + subResult1.at(j));
                        }
                }
            }
            result.assign(temp.begin(), temp.end());
            record.insert({{n, result}});
        }
        else
        {
            result = record.find(n)->second;
        } 
        return result;
    }

    vector<string> generateParenthesis(int n) {
        return recursion(n);
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    for (int i = 1; i < 5; i++)
    {
        cout << i << endl;
        vector<string> result = s->generateParenthesis(i);
        for (int j = 0; j < result.size(); j++)
        {
            cout << result.at(j) << endl;
        }
        cout << endl;
    }
    return 0;
}
