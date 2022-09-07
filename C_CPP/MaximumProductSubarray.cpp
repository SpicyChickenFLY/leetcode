#include <vector>

class Solution {
public:
    int maxProduct(std::vector<int>& nums) {
        int maxEnd = nums.at(0);
        int minEnd = nums.at(0);
        int maxTotal = nums.at(0);
        for (int i = 1; i < nums.size(); i++) {
            // std::cout << "before:" << nums.at(i) << " " << 
            // maxEnd << " " << minEnd << " " << maxTotal << endl; 
            if (nums.at(i) < 0) {
                int temp = maxEnd;
                maxEnd = minEnd;
                minEnd = temp;
            }
            // maxEnd
            if (maxEnd * nums.at(i) > nums.at(i))
                maxEnd *= nums.at(i);
            else
                maxEnd = nums.at(i);
            // minEnd
            if (minEnd * nums.at(i) < nums.at(i))
                minEnd *= nums.at(i);
            else
                minEnd = nums.at(i);
            // maxTotal
            if (maxEnd > maxTotal)
                maxTotal = maxEnd;
            // std::cout << "after:" << nums.at(i) << " " << 
            // maxEnd << " " << minEnd << " " << maxTotal << endl; 
        }
        return maxTotal;
    }
};

int main()
{
    Solution *solution = new Solution();
    // vector<vector<int>> testCases;
    int arrTestCase1[] = {2, 3, -2 ,4};
    vector<int> testCase1(arrTestCase1, arrTestCase1 + 4); 
    std::cout << solution->maxProduct(testCase1) << endl;
    return 0;
}