package main

var topicList = []string{
	"Arrays & Hashing",
	"Two Pointers",
	"Stack",
	"Binary Search",
	"Sliding Window",
	"Linked List",
	"Trees",
	"Tries",
	"Backtracking",
	"Heap & Priority Queue",
	"Graphs",
	"1-D DP",
	"Intervals",
	"Greedy",
	"Advanced Graphs",
	"2-D DP",
	"Bit Manipulation",
	"Math & Geometry",
}

var problemLists = map[string][]string{
	"Arrays & Hashing": {
		"Contains Duplicate",
		"Valid Anagram",
		"Two Sum",
		"Group Anagrams",
		"Top K Frequent Elements",
		"Product of Array Except Self",
		"Valid Sudoku",
		"Longest Consecutive Sequence",
	},
	"Two Pointers": {
		"Valid Palindrome",
		"Two Sum II Input Array Is Sorted",
		"3Sum",
		"Container With Most Water",
		"Trapping Rain Water",
	},
	"Stack": {
		"Valid Parentheses",
		"Min Stack",
		"Evaluate Reverse Polish Notation",
		"Generate Parentheses",
		"Daily Temperatures",
		"Car Fleet",
		"Largest Rectangle In Histogram",
	},
	"Binary Search": {
		"Binary Search",
		"Search a 2D Matrix",
		"Koko Eating Bananas",
		"Find Minimum In Rotated Sorted Array",
		"Search In Rotated Sorted Array",
		"Time Based Key Value Store",
		"Median of Two Sorted Arrays",
	},
	"Sliding Windows": {
		"Best Time to Buy And Sell Stock",
		"Longest Substring Without Repeating Characters",
		"Longest Repeating Character Replacement",
		"Permutation in String",
		"Minimum Window Substring",
		"Sliding Window Maximum",
	},
	"Linked List": {
		"Reverse Linked List",
		"Merge Two Sorted Lists",
		"Reorder List",
		"Remove Nth Node From End of List",
		"Copy List With Random Pointer",
		"Add Two Numbers",
		"Linked List Cycle",
		"Find The Duplicate Number",
		"LRU Cache",
		"Merge K Sorted Lists",
		"Reverse Nodes In K Group",
	},
	"Trees": {
		"Invert Binary Tree",
		"Maximum Depth of Binary Tree",
		"Diameter of Binary Tree",
		"Balanced Binary Tree",
		"Same Tree",
		"Subtree of Another Tree",
		"Lowest Common Ancestor of a Binary Search Tree",
		"Binary Tree Level Order Traversal",
		"Binary Tree Right Side View",
		"Count Good Notes In Binary Tree",
		"Validate Binary Search Tree",
		"Kth Smallest Element In a Binary Search Tree",
		"Construct Binary Tree From Preorder And Inorder Traversal",
		"Binary Tree Maximum Path Sum",
		"Serialize And Deserialize Binary Tree",
	},
	"Tries": {
		"Implement Trie Prefix Tree",
		"Design Add And Search Words Data",
		"Word Search II",
	},
	"Backtracking": {
		"Subsets",
		"Combination Sum",
		"Permutations",
		"Subsets II",
		"Combination Sum II",
		"Word Search",
		"Palindrome Partitioning",
		"Letter Combinations of a Phone Number",
		"N Queens",
	},
	"Heap & Priority Queue": {
		"Kth Largest Element In a Stream",
		"Last Stone Weight",
		"K Closest Points to Origin",
		"Kth Largest Element In An Array",
		"Task Scheduler",
		"Design Twitter",
		"Find Median From Data Stream",
	},
	"Graphs": {
		"Number of Islands",
		"Max Area of Island",
		"Clone Graph",
		"Rotting Oranges",
		"Pacific Atlantic Water Flow",
		"Surrounded Regions",
		"Course Schedule",
		"Course Schedule II",
		"Redundant Connection",
		"Word Ladder",
	},
	"1-D Dynamic Programming": {
		"Climbing Stairs",
		"Min Cost Climbing Stairs",
		"House Robber",
		"House Robber II",
		"Longest Palindromic Substring",
		"Palindromic Substrings",
		"Decode Ways",
		"Coin Change",
		"Maximum Product Subarray",
		"Word Break",
		"Longest Increasing Subsequence",
		"Partition Equal Subset Sum",
	},
	"Intervals": {
		"Insert Interval",
		"Merge Intervals",
		"Non Overlapping Intervals",
		"Minimum Interval to Include Each Query",
	},
	"Greedy": {
		"Maximum Subarray",
		"Jump Game",
		"Jump Game II",
		"Gas Station",
		"Hand of Straights",
		"Merge Triplets to Form Target Triplet",
		"Partition Labels",
		"Valid Parenthesis String",
	},
	"Advanced Graphs": {
		"Reconstruct Itinerary",
		"Min Cost to Connect All Points",
		"Network Delay Time",
		"Swim In Rising Water",
		"Cheapest Flights Within K Stops",
	},
	"2-D Dynamic Programming": {
		"Unique Paths",
		"Longest Common Subsequence",
		"Best Time to Buy and Sell Stock With Cooldown",
		"Coin Change II",
		"Target Sum",
		"Interleaving String",
		"Longest Increasing Path In a Matrix",
		"Distinct Subsequences",
		"Edit Distance",
		"Burst Balloons",
		"Regular Expression Matching",
	},
	"Bit Manipulation": {
		"Single Number",
		"Number of 1 Bits",
		"Counting Bits",
		"Reverse Bits",
		"Missing Number",
		"Sum of Two Integers",
		"Reverse Integer",
	},
	"Math & Geometry": {
		"Rotate Image",
		"Spiral Matrix",
		"Set Matrix Zeroes",
		"Happy Number",
		"Plus One",
		"Pow(x, n)",
		"Multiply Strings",
		"Detect Squares",
	},
}