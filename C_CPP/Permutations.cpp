#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void printVV(vector<vector<int>> vv)
{
    for (vector<vector<int>>::iterator it = vv.begin(); it != vv.end(); ++it)
    {
        for (vector<int>::iterator it1 = it->begin(); it1 != it->end(); ++it1)
        {
            cout << *it1 << " ";
        }
        cout << endl;
    }
    cout << endl;
}

void printV(vector<int> v)
{
    for (vector<int>::iterator it = v.begin(); it != v.end(); ++it)
    {
        cout << *it << " ";
    }
    cout << endl;
}

class Solution
{
public:
    vector<vector<int>> subPermute(vector<int> &targets) //proved poor performance
    {
        vector<vector<int>> results;
        if (targets.size() == 1) 
        {
            vector<int> result(1, targets.at(0));
            results.push_back(result);
        }
        else
        {
            for (int i = 0; i < targets.size(); i++)
            {
                int target = targets.at(i);
                targets.erase(targets.begin() + i);
                vector<vector<int>> temp_results = subPermute(targets);
                for (vector<vector<int>>::iterator it = temp_results.begin(); it != temp_results.end(); ++it)
                {
                    it->insert(it->begin(), target);
                    results.push_back(*it);
                }
                targets.insert(targets.begin() + i, target);
            }
        }
        return results;
    }

    vector<vector<int>> permute(vector<int> &targets)
    {
        sort(targets.begin(), targets.end());
        return subPermute(targets);
    }
};

int main()
{
    Solution *s = new Solution();
    int arr[] = {1}; 
    vector<int> input(arr, arr + sizeof(arr) / sizeof(int));
    vector<vector<int>> results = s->permute(input);
    printVV(results);
    return 0;
}