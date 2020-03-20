package tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。

示例 2:

输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
*/

/*
我们可以使用递归的方法得到二叉树的前序遍历。在递归时，根据题目描述，我们需要加上额外的括号，会有以下 4 种情况：

    如果当前节点有两个孩子，那我们在递归时，需要在两个孩子的结果外都加上一层括号；

    如果当前节点没有孩子，那我们不需要在节点后面加上任何括号；

    如果当前节点只有左孩子，那我们在递归时，只需要在左孩子的结果外加上一层括号，而不需要给右孩子加上任何括号；

    如果当前节点只有右孩子，那我们在递归时，需要先加上一层空的括号 () 表示左孩子为空，再对右孩子进行递归，并在结果外加上一层括号。

考虑完上面的 4 种情况，我们就可以得到最终的字符串。

*/

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}

	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}

	// if t.Left==nil {
	//     return strconv.Itoa(t.Val)+"()"+"("+tree2str(t.Right)+")"
	// }

	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
}
