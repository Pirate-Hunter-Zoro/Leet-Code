package leetcode

import (
	"leet-code/datastructures"
	"reflect"
	"testing"
)

func runTestHelper[I any, A any](t *testing.T, f func(i I) A, inputs []I, expected_outputs []A) {
	for idx, input := range inputs {
		output := f(input)
		expected_output := expected_outputs[idx]
		if !reflect.DeepEqual(output, expected_output) {
			t.Fatalf("Error - expected %v but got %v", expected_output, output)
		}
	}
}

func TestSuperEggDrop(t *testing.T) {
	type input struct {
		n int
		k int
	}

	inputs := []input{
		{n:1, k:2},
		{n:2, k:6},
		{n:3, k:14},
		{n:1, k:1},
		{n:7, k:5000},
	}

	expected_outputs := []int{
		2,3,4,1,13,
	}

	f := func(i input) int {
		return superEggDrop(i.n, i.k)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestLargestRectangleArea(t *testing.T) {
	type input struct {
		heights []int
	}
	inputs := []input{
		{heights : []int{2,1,5,6,2,3}},
		{heights : []int{2,4}},
	}
	
	expected_outputs := []int{
		10,
		4,
	}

	f := func(i input) int {
		return largestRectangleArea(i.heights)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestMaximalRectangle(t *testing.T) {
	type input struct {
		matrix [][]byte
	}
	inputs := []input{
		{matrix: [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}},
		{matrix: [][]byte{{'0'}}},
		{matrix: [][]byte{{'1'}}},
	}

	expected_outputs := []int{
		6,0,1,
	}

	f := func(i input) int {
		return maximalRectangle(i.matrix)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestMaxMove(t *testing.T) {
	type input struct {
		kx int
		ky int
		positions [][]int
	}

	inputs := []input{
		{1, 1, [][]int{{0,0}}},
		{0, 2, [][]int{{1,1},{2,2},{3,3}}},
		{0, 0, [][]int{{1,2},{2,4}}},
		{7, 3, [][]int{{2,2},{10,0}}},
		{0, 1, [][]int{{9,6},{8,0},{4,0}}},
	}

	expected_outputs := []int{
		4,
		8,
		3,
		8,
		12,
	}

	f := func(i input) int {
		return maxMoves(i.kx, i.ky, i.positions)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestCanFinish(t *testing.T) {
	type input struct {
		numCourses int
		prerequisites [][]int
	}
	inputs := []input{
		{2, [][]int{{1,0}}},
		{2, [][]int{{1,0},{0,1}}},
	}

	expected_outputs := []bool{true, false}

	f := func(i input) bool {
		return canFinish(i.numCourses, i.prerequisites)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestIsMatch(t *testing.T) {
	type input struct {
		s string
		p string
	}
	inputs := []input{
		{"aa", "a"},
		{"aa", "*"},
		{"cb", "?a"},
		{"adceb", "*a*b"},
	}

	expected_outputs := []bool{
		false,
		true,
		false,
		true,
	}

	f := func(i input) bool {
		return isMatch(i.s, i.p)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestFindSubstring(t *testing.T) {
	type input struct {
		s string
		words []string
	}
	inputs := []input{
		{"barfoothefoobarman", []string{"foo","bar"}},
		{"wordgoodgoodgoodbestword", []string{"word","good","best","word"}},
		{"barfoofoobarthefoobarman", []string{"bar","foo","the"}},
		{"aaaaaaaaaaaaaa", []string{"aa","aa"}},
	}

	expected_outputs := [][]int{
		{0,9},
		{},
		{6,9,12},
		{0,1,2,3,4,5,6,7,8,9,10},
	}

	f := func(i input) []int{
		return findSubstring(i.s, i.words)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestFirstMissingPositive(t *testing.T) {
	type input struct {
		nums []int
	}
	inputs := []input{
		{[]int{1,2,0}},
		{[]int{3,4,-1,1}},
		{[]int{7,8,9,11,12}},
	}

	expected_outputs := []int{
		3,
		2,
		1,
	}

	f := func(i input) int {
		return firstMissingPositive(i.nums)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestMergeKLists(t *testing.T) {
	type input struct {
		lists []*datastructures.ListNode
	}
	inputs := []input{
		{
			[]*datastructures.ListNode{
				datastructures.NewListNode([]int{1,4,5}),
				datastructures.NewListNode([]int{1,3,4}),
				datastructures.NewListNode([]int{2,6}),
			},
		},
		{
			[]*datastructures.ListNode{},
		},
		{
			[]*datastructures.ListNode{
				datastructures.NewListNode([]int{}),
			},
		},
	}

	expected_outputs := []*datastructures.ListNode{
		datastructures.NewListNode([]int{1,1,2,3,4,4,5,6}),
		nil,
		nil,
	}

	f := func(i input) *datastructures.ListNode {
		return mergeKLists(i.lists)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestLongestValidParentheses(t *testing.T) {
	type input struct {
		s string
	}
	inputs := []input{
		{"(()"},
		{")()())"},
		{""},
		{"()(())"},
		{")(())(()()))("},
	}

	expected_outputs := []int{
		2,
		4,
		0,
		6,
		10,
	}

	f := func(i input) int {
		return longestValidParentheses(i.s)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestIsScramble(t *testing.T) {
	type input struct {
		s1 string
		s2 string
	}
	inputs := []input{
		{"great","rgeat"},
		{"abcde", "caebd"},
		{"a", "a"},
	}

	expected_outputs := []bool{
		true,
		false,
		true,
	}

	f := func(i input) bool {
		return isScramble(i.s1, i.s2)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestNumDistinct(t *testing.T) {
	type input struct {
		s string
		t string
	}
	inputs := []input{
		{"rabbbit", "rabbit"},
		{"babgbag", "bag"},
		{"b", "a"},
	}

	expected_outputs := []int{
		3,
		5,
		0,
	}

	f := func(i input) int {
		return numDistinct(i.s, i.t)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestMaxStudents(t *testing.T) {
	type input struct {
		seats [][]byte
	}

	inputs := []input{
		{[][]byte{
			{'#','.','#','#','.','#'},
            {'.','#','#','#','#','.'},
            {'#','.','#','#','.','#'},
		}},
		{[][]byte{
			{'.','#'},
            {'#','#'},
            {'#','.'},
			{'#','#'},
			{'.','#'},
		}},
		{[][]byte{
			{'#','.','.','.','#'},
			{'.','#','.','#','.'},
			{'.','.','#','.','.'},
			{'.','#','.','#','.'},
			{'#','.','.','.','#'},
		}},
	}

	expected_outputs := []int{
		4,
		3,
		10,
	}

	f := func(i input) int {
		return maxStudents(i.seats)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}


func TestConnectTwoGroups(t *testing.T) {
	type input struct {
		cost [][]int
	}

	inputs := []input{
		{[][]int{{15, 96}, {36, 2}}},
		{[][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}},
		{[][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}}},
	}

	expected_outputs := []int{
		17,
		4,
		10,
	}

	f := func(i input) int {
		return connectTwoGroups(i.cost)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}


func TestMinCut(t *testing.T) {
	type input struct {
		s string
	}

	inputs := []input{
		{"aab"},
		{"a"},
		{"ab"},
		{"cabababcbc"},
	}

	expected_outputs := []int{
		1,
		0,
		1,
		3,
	}

	f := func(i input) int {
		return minCut(i.s)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}


func TestCandy(t *testing.T) {
	type input struct {
		ratings []int
	}

	inputs := []input{
		{[]int{1,0,2}},
		{[]int{1,2,2}},
	}

	expected_outputs := []int{
		5,
		4,
	}

	f := func(i input) int {
		return candy(i.ratings)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}


func TestGetPermutation(t *testing.T) {
	type input struct {
		n int
		k int
	}
	inputs := []input{
		{3,3},
		{4,9},
		{3,1},
		{2,2},
	}

	expected_outputs := []string{
		"213",
		"2314",
		"123",
		"21",
	}

	f := func(i input) string {
		return getPermutation(i.n, i.k)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}

func TestFindWords(t *testing.T) {
	type input struct{
		board [][]byte
		words []string
	}
	inputs := []input{
		{[][]byte{
			{'o','a','a','n'},
			{'e','t','a','e'},
			{'i','h','k','r'},
			{'i','f','l','v'},
		}, []string{"oat","oath","pea","eat","rain"}},
		{[][]byte{
			{'a','b'},
			{'c','d'},
		}, []string{"abcb"}},
	}

	expected_outputs := [][]string{
		{"eat","oat","oath"},
		{},
	}

	f := func(i input) []string {
		return findWords(i.board, i.words)
	}

	runTestHelper(t, f, inputs, expected_outputs)

}

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	type input struct {
		nums []int
		indexDiff int
		valueDiff int
	}
	inputs := []input{
		{[]int{1,2,3,1}, 3, 0},
		{[]int{1,2,1,1}, 1, 0},
		{[]int{1,5,9,1,5,9}, 2, 3},
		{[]int{-2,3}, 2, 5},
		{[]int{-3,3}, 2, 4},
		{[]int{7,2,8}, 2, 1},
	}

	expected_outputs := []bool{
		true,
		true,
		false,
		true,
		false,
		true,
	}

	f := func(i input) bool {
		return containsNearbyAlmostDuplicate(i.nums, i.indexDiff, i.valueDiff)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}


func TestGetSkyline(t *testing.T) {
	type input struct {
		buildings [][]int
	}
	inputs := []input{
		{[][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}},
		{[][]int{{0,2,3},{2,5,3}}},
		{[][]int{{1,2,1},{1,2,2},{1,2,3}}},
		{[][]int{{3,7,8},{3,8,7},{3,9,6},{3,10,5},{3,11,4},{3,12,3},{3,13,2},{3,14,1}}},
		{[][]int{{4547,253463,513907},{16593,154606,197418},
				{20536,587972,480271},{25155,135407,277021},
				{45693,480063,10484},{57392,70775,887782},
				{69423,760082,373899},{93047,222798,416217},
				{105421,832604,877583},{113118,797930,760771},
				{128302,234173,548113},{137036,889837,902341},
				{151470,195177,920702},{153020,398494,465441},
				{153980,501882,565647},{159005,416356,28559},
				{218765,535515,140331},{303595,627878,399999},
				{333489,449295,827098},{359242,830671,455392},
				{390291,400129,693584},{429774,450923,637981},
				{493315,563384,655884},{500860,718561,769319},
				{507299,817616,228},{537517,989102,300696},
				{543853,843011,695472},{550404,559944,90888},
				{617555,907369,523648},{622624,989745,894596},
				{640808,867707,963706},{697210,716449,705689},
				{804555,914745,764029}}},
	}

	expected_outputs := [][][]int{
		{{2,10},{3,15},{7,12},{12,0},{15,10},{20,8},{24,0}},
		{{0,3},{5,0}},
		{{1,3},{2,0}},
		{{3,8},{7,7},{8,6},{9,5},{10,4},{11,3},{12,2},{13,1},{14,0}},
		{{4547,513907},{57392,887782},
			{70775,513907},{105421,877583},
			{137036,902341},{151470,920702},
			{195177,902341},{640808,963706},
			{867707,902341},{889837,894596},
			{989745,0}}, // TODO - NOT SEEN - instead of {867707,902341},{889837,894596}, saw {867707,894596}
	}

	f := func(i input) [][]int {
		return getSkyline(i.buildings)
	}

	runTestHelper(t, f, inputs, expected_outputs)
}