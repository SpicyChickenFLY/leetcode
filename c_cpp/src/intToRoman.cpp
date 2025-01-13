#include <iostream>
using namespace std;

class Solution {
	public:
	struct Node {
		int value;
		char symbol;
		struct Node *next;
		Node(int value, char symbol, struct Node *next = NULL) {
			this->value = value;
			this->symbol = symbol;
			this->next = next;
		}
	};

	string intToRoman(int num) {
		//initalize linklist
		struct Node *I = new Node(1, 'I'),
		*V = new Node(5, 'V', I), *X = new Node(10, 'X', I),
		*L = new Node(50, 'L', X), *C = new Node(100, 'C', X),
		*D = new Node(500, 'D', C), *M = new Node(1000, 'M', C),
		*h = *M;
		string symbol = "";
		//extract each number
		while (num > 0) {
			struct Node *temp = h;
			if (num / temp->value >= 1)
		}
	}
};


int main() {
	int test[5] = {3, 4, 9, 58, 1994};
	Solution s = new Solution();
	for (int i  = 0; i < 5;i++) {
		cout << s.intToRoman(test[i]) << endl;
	}
}
