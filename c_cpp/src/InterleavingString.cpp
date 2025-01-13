#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
	bool isInterleave(string s1, string s2, string s3) {
		int s1len =  s1.length();
		int s2len =  s2.length();
		int s3len =  s3.length();
		if (s1len + s2len != s3len) {
			return false;
		}
		if (s1len == 0) return s3 == s2;
		if (s2len == 0) return s3 == s1;

		vector<bool> temp(s1len + 1);
		temp[0] = true;
		for (size_t j = 1; j <= s1len; j++) {
			temp[j] = temp[j-1] && s3[j - 1] == s1[j-1];
		}

		for (size_t i = 1; i <= s2len; i++) {
			temp[0] = temp[0] && s3[i-1] == s2[i-1];
			for (size_t j = 1; j <= s1.length(); j++) {
				temp[j] = (temp[j-1] && s3[i+j-1] == s1[j-1])
					|| (temp[j] && s3[i+j-1] == s2[i-1]);

			}
		}
		return temp[s1len];
	}
};


int main() {
	Solution *s = new Solution();
	cout << s->isInterleave("db", "b", "cbb") << endl;
}
