package main

import (
	"fmt"
	"github.com/VladimirZaets/golang-leetcode/array"
	dynamic_programming "github.com/VladimirZaets/golang-leetcode/dynamic-programming"
	"github.com/VladimirZaets/golang-leetcode/graph"
	"github.com/VladimirZaets/golang-leetcode/graph/disjointset"
	str "github.com/VladimirZaets/golang-leetcode/string"
)

func main() {
	executeDynamicProgramming()
	executeArrays()
	executeGraphDisjointSet()
	executeString()
	executeGraph()
}

func executeString() {
	str.AddBinary("11", "1")
	str.StrStr("hello", "ll")
	str.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	str.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	str.ReverseWords("the sky is blue")
	str.ReverseWordsInSting2("Let's take LeetCode contest")
}

func executeGraph() {
	r0 := graph.Preorder(&graph.Node{
		Val: 1,
		Children: []*graph.Node{
			{
				Val: 3,
				Children: []*graph.Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	})
	r1 := graph.PreorderTraversal(&graph.TreeNode{
		Val: 1,
		Left: &graph.TreeNode{
			Val: 2,
		},
		Right: &graph.TreeNode{
			Val: 3,
		},
	})
	r2 := graph.ClosedIsland([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	})
	r3 := graph.ClosedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	})
	r4 := graph.UpdateBoard([][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}, []int{3, 0})
	r5 := graph.NumEnclaves([][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	})
	r6 := graph.NumIslands([][]byte{
		{49, 49, 49, 49, 48},
		{49, 49, 48, 49, 48},
		{49, 49, 48, 48, 48},
		{48, 48, 48, 48, 48},
	})
	r7 := graph.VerticalOrder(&graph.TreeNode{
		Val: 3,
		Left: &graph.TreeNode{
			Val: 9,
			Left: &graph.TreeNode{
				Val: 4,
			},
			Right: &graph.TreeNode{
				Val: 0,
			},
		},
		Right: &graph.TreeNode{
			Val: 8,
			Left: &graph.TreeNode{
				Val: 1,
			},
			Right: &graph.TreeNode{
				Val: 7,
			},
		},
	})
	r8 := graph.InvertTree(&graph.TreeNode{
		Val: 4,
		Left: &graph.TreeNode{
			Val: 2,
			Left: &graph.TreeNode{
				Val: 1,
			},
			Right: &graph.TreeNode{
				Val: 3,
			},
		},
		Right: &graph.TreeNode{
			Val: 7,
			Left: &graph.TreeNode{
				Val: 6,
			},
			Right: &graph.TreeNode{
				Val: 9,
			},
		},
	})
	r9 := graph.InOrderTraversal(&graph.TreeNode{
		Val: 1,
		Right: &graph.TreeNode{
			Val: 2,
			Left: &graph.TreeNode{
				Val: 3,
			},
		},
	})
	r10 := graph.Postorder(&graph.Node{
		Val: 1,
		Children: []*graph.Node{
			{
				Val: 3,
				Children: []*graph.Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	})
	r11 := graph.LevelOrder(&graph.Node{
		Val: 1,
		Children: []*graph.Node{
			{
				Val: 3,
				Children: []*graph.Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	})
	fmt.Println(r0)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)
	fmt.Println(r6)
	fmt.Println(r7)
	fmt.Println(r8)
	fmt.Println(r9)
	fmt.Println(r10)
	fmt.Println(r11)
}

func executeGraphDisjointSet() {
	numberOfConnectedComponents := disjointset.NewNumberOfConnectedComponents(5)
	numberOfConnectedComponents.Union(0, 1)
	pathCompressionUnionFind := disjointset.NewPathCompressionUnionFind(5)
	pathCompressionUnionFind.Union(0, 1)
	smallestStringWIthSwaps := disjointset.NewSmallestStringWIthSwaps(5)
	smallestStringWIthSwaps.Union(0, 1)
	disjointset.ValidTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}})
	disjointset.FindCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	quickFind := disjointset.NewQuickFind(5)
	quickFind.Union(0, 1)
	quickUnion := disjointset.NewQuickUnion(5)
	quickUnion.Union(0, 1)
	disjointset.EarliestAcq([][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6)
	unionByRank := disjointset.NewUnionByRank(5)
	unionByRank.Union(0, 1)
}

func executeArrays() {
	array.MaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1, 1})
	array.FindMaxConsecutiveOnes2([]int{1, 0, 1, 1, 0, 1, 1})
	array.HeightChecker([]int{1, 1, 4, 2, 1, 3})
	array.ValidMountainArray2([]int{0, 3, 2, 1})
	array.ValidMountainArray([]int{0, 3, 2, 1})
	array.CheckIfNandItsDoubleExist([]int{10, 2, 5, 3})
	array.DuplicateZerosSimple([]int{1, 0, 2, 3, 0, 4, 5, 0})
	array.DuplicateZerosComplex([]int{1, 0, 2, 3, 0, 4, 5, 0})
	array.FindDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	array.ArrayPairSum([]int{1, 4, 3, 2})
	array.FindNumbers([]int{12, 345, 2, 6, 7896})
	array.FindPivotIndex([]int{1, 7, 3, 6, 5, 6})
	array.DominantIndex([]int{3, 6, 1, 0})
	array.FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	array.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	array.PascaleTriangle(5)
	array.PascaleTriangle2(5)
	array.PlusOne([]int{1, 2, 3})
	array.RemoveDuplicates([]int{1, 1, 2})
	array.RemoveDuplicates2([]int{1, 1, 2})
	array.RemoveElement([]int{3, 2, 2, 3}, 3)
	array.RemoveElement2([]int{3, 2, 2, 3}, 3)
	array.RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	array.SpiralMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	array.SortedSquaresFirst([]int{-4, -1, 0, 3, 10})
	array.TwoSum([]int{2, 7, 11, 15}, 9)
	array.MergeSortArray([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	array.MergeSortArray([]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1)
	array.MergeSortArray([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	array.ReplaceElementWithGreatestElementOnRightSide([]int{17, 18, 5, 4, 6, 1})
	array.MoveZeroes([]int{0, 1, 0, 3, 12})
	array.SortArrayByParity([]int{3, 1, 2, 4})
	array.FindAllNumbersDisappearedInArray([]int{4, 3, 2, 7, 8, 2, 3, 1})
	array.FindDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	array.ThirdMaximumNumber([]int{3, 2, 1})
}

func executeDynamicProgramming() {
	r1 := dynamic_programming.Fib(10)
	r2 := dynamic_programming.PascalsTriangleGenerate(5)
	r3 := dynamic_programming.ClimbStairs(3)
	r4 := dynamic_programming.MaxProfit([]int{2, 1, 2, 1, 0, 1, 2})
	r5 := dynamic_programming.CombinationSum4([]int{1, 2, 3}, 4)
	r6 := dynamic_programming.LongestPalindrome("aaaa")
	r7 := dynamic_programming.MaxSubArray([]int{-2, -1})
	r8 := dynamic_programming.CheckValidString("(*)")
	r9 := dynamic_programming.GenerateParenthesis(3)
	r10 := dynamic_programming.GenerateParenthesisMemo(3)
	r11 := dynamic_programming.PpredictTheWinner([]int{1, 5, 100, 7})
	r12 := dynamic_programming.PredictTheWinner([]int{1, 5, 100, 7})
	r13 := dynamic_programming.IsValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
	r14 := dynamic_programming.Rob([]int{1, 2, 3, 1})
	r15 := dynamic_programming.MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	r16 := dynamic_programming.Tribonacci(4)
	r17 := dynamic_programming.DeleteAndEarn([]int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10})
	r18 := dynamic_programming.LongestCommonSubsequence("abcde", "ace")
	r19 := dynamic_programming.MaximumScore([]int{1, 2, 3}, []int{3, 2, 1})
	dynamic_programming.SolveSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
	fmt.Println(r1, r2, r3, r4, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19)
}
