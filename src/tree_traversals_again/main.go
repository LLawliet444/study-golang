package main

import(
	"fmt"
	// "os"
)

/* An inorder binary tree traversal can be implemented in a non-recursive way with a stack.
For example, suppose that when a 6-node binary tree (with the keys numbered from 1 to 6) is
traversed, the stack operations are: push(1); push(2); push(3); pop(); pop(); push(4); pop()
;pop(); push(5); push(6); pop(); pop(). Then a unique binary tree (shown in Figure 1) can be
 generated from this sequence of operations. Your task is to give the postorder traversal
 sequence of this tree. */
func main(){
	preorder := [6]int{1,2,3,4,5,6}
	inorder := [6]int{3,2,4,1,6,5}
	postorder :=make([]int,len(preorder),len(preorder))
	// fmt.Printf("%T\n",postorder)
	// fmt.Printf("%T\n",preorder)
	// fmt.Printf("%T\n",inorder)
	solve( 0,0,0,6,preorder,inorder,postorder )
	fmt.Println(postorder)
}

func solve( preL int,inL int,postL int,n int,preorder [6]int,inorder [6]int,postorder []int ){
	if n == 0{
		return
	}
	if n == 1{
		postorder[postL] = preorder[preL]
		return
	}

	root := preorder[preL]
	postorder[postL + n - 1] = root
	// fmt.Println(postorder)
	// 左子树
	left := 0
	for i := inL;i < len(inorder);i++{
		if inorder[i] == root {
			break
		}
		left++
	}
	// fmt.Println(left)
	// fmt.Println(postorder)

	solve( preL + 1,inL,postL,left,preorder,inorder,postorder )

	// 右子树
	right := n - left - 1
	solve( preL + 1 + left,inL + left + 1,postL + left,right,preorder,inorder,postorder )
}