package main

import (
	"fmt"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&list=WL&index=13

func main() {
	// Fibonacci
	// fmt.Println(fibWithMemo(50, map[int]int{}))
	// fmt.Println(fib(50))

	// Travel
	// fmt.Println(travelWithMemo(18,18, map[string]int{}))
	// fmt.Println(travel(18,18))

	// CanSum
	// fmt.Println(canSumWithMemo(7, []int{5, 3, 4, 7}, map[int]bool{}))
	// fmt.Println(canSumWithMemo(7, []int{2,4}, map[int]bool{}))
	// fmt.Println(canSumWithMemo(300, []int{7,21}, map[int]bool{}))
	// fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	// fmt.Println(canSum(7, []int{2,4}))
	// fmt.Println(canSum(300, []int{7,21}))

	// HowSum
	// fmt.Println(howSumWithMemo(300, []int{7,14}, map[int][]int{}))
	// fmt.Println(howSum(300, []int{7,14}))
	// fmt.Println(howSumWithMemo(300, []int{7,15}, map[int][]int{}))
	// fmt.Println(howSum(300, []int{7,15}))

	// BestSum
	// fmt.Println(bestSumWithMemo(35, []int{1,2,5,10}, map[int][]int{}))
	// fmt.Println(bestSum(35, []int{1,2,5,10}))

	// CanConstruct
	// fmt.Println(canConstructWithMemo("abcdef", []string{"ab","abc","cd","def","abcd"}, map[string]bool{}))
	// fmt.Println(canConstructWithMemo("skateboard", []string{"bo","rd","ate","t","ska","sk","boar"}, map[string]bool{}))
	// fmt.Println(canConstructWithMemo("enterapotentpot", []string{"a","p","ent","enter","ot","o","t"}, map[string]bool{}))
	// fmt.Println(canConstructWithMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e","ee","eee","eeee","eeeee"}, map[string]bool{}))
	// fmt.Println(canConstruct("abcdef", []string{"ab","abc","cd","def","abcd"}))
	// fmt.Println(canConstruct("skateboard", []string{"bo","rd","ate","t","ska","sk","boar"}))
	// fmt.Println(canConstruct("enterapotentpot", []string{"a","p","ent","enter","ot","o","t"}))
	// fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e","ee","eee","eeee","eeeee"}))

	// CountConstruct
	fmt.Println(countConstructWithMemo("abcdef", []string{"ab","abc","cd","def","ef","ab"}, map[string]int{}))
	fmt.Println(countConstructWithMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"e","ee","eee","eeee","eeeee"}, map[string]int{}))
	fmt.Println(countConstruct("abcdef", []string{"ab","abc","cd","def","ef","ab"}))
	fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"e","ee","eee","eeee","eeeee"}))
}
