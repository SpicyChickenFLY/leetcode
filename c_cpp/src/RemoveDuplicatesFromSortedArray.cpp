#include <iostream>
#include <vector>
using namespace std;

class Solution {
   public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() != 0) {
            int last = nums.at(0) + 1;
            for (vector<int>::iterator it = nums.begin(); it != nums.end();) {
                if (*it == last)
                    it = nums.erase(it);
                else{
                    last = *it;
                    it++; 
                }
            }
        }
        return nums.size();
    }
};

int main() {
    vector<vector<int>> testCases = {{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}};

    Solution* s = new Solution();

    for (vector<int>& testCase : testCases) {
        int result = s->removeDuplicates(testCase);
        cout << "length: " << result << endl;
        cout << "resultArr: " << endl;
        for (int val : testCase) {
            cout << val << " "; 
        }
        cout << endl;
    }
    return 0;
}