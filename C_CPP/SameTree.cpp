#include <iostream>
#include <vector>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x)
    {
        val = x;
        left = NULL;
        right = NULL;
    }
};

class Solution
{
public:

    bool isSameTree(TreeNode *p, TreeNode *q)
    {
        if (p == NULL || q == NULL) 
            return p == q ? true : false;
        return p->val == q->val && isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
    }
};

int main(int argc, char const *argv[])
{
    Solution *s = new Solution();
    TreeNode *h1 = new TreeNode(1);
    h1->left = new TreeNode(2);
    TreeNode *h2 = new TreeNode(1);
    h2->right = new TreeNode(2);
    cout << s->isSameTree(h1, h2) << endl;
    return 0;
}
