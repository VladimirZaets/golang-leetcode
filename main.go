package main

import (
	"github.com/VladimirZaets/golang-leetcode/array"
	"github.com/VladimirZaets/golang-leetcode/graph/disjointset"
	str "github.com/VladimirZaets/golang-leetcode/string"
)

func main() {
	executeArrays()
	executeGraphDisjointSet()
	executeString()
}

func executeString() {
	str.AddBinary("11", "1")
	str.StrStr("hello", "ll")
	str.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	str.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	str.ReverseWords("the sky is blue")
	str.ReverseWordsInSting2("Let's take LeetCode contest")
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
}
