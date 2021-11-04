/* Tsyren Ochirov (c) 2021
  solution to https://leetcode.com/problems/binary-tree-level-order-traversal-ii
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
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
        *levelNodes = append([][]int{[]int{node.Val}}, (*levelNodes)...)
    } else {
        n := len(*levelNodes) - level - 1
        (*levelNodes)[n] = append((*levelNodes)[n], node.Val)        
    }
    
    if node.Left != nil {
        traverse(node.Left, levelNodes, level+1)
    }
    
    if node.Right != nil {
        traverse(node.Right, levelNodes, level+1)    
    }
}
