package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type recoverBinarySearchTreeTestCase struct {
    root *TreeNode
    expect *TreeNode
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRecoverBinarySearchTree(t *testing.T) {
    testCases := []recoverBinarySearchTreeTestCase {
        {
            root: newTree([]*TreeNode{newTreeNode(1), newTreeNode(3), nil, nil, newTreeNode(2)}),
            expect: newTree([]*TreeNode{newTreeNode(3), newTreeNode(1), nil, nil, newTreeNode(2)}),
        },
    }

    for _, testCase := range testCases {
        recoverTree(testCase.root)
        assert.True(t, compareTwoTree(testCase.expect, testCase.root), "结果不一致")
    }
}

