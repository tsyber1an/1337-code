/* Tsyren Ochirov (c) 2021
   LeedCode solution https://leetcode.com/problems/binary-tree-level-order-traversal/submissions/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {    
    res := [][]int{}
    level := 0
    traverse(root, &res, level)
    
    return res
}

func traverse(node *TreeNode, levelNodes *[][]int, level int) {
    if node == nil {
        return
    }
    
    if len(*levelNodes) == level {
        *levelNodes = append(*levelNodes, []int{node.Val})
    } else {
        (*levelNodes)[level] = append((*levelNodes)[level], node.Val)
    }
    
    traverse(node.Left, levelNodes, level+1)    
    traverse(node.Right, levelNodes, level+1)
}
