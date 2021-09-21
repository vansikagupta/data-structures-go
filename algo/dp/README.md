# Dynamic programming
Dynamic programming is an programming approach in which we break down a problem into smaller overlapping sub-problems and construct the solution to larger problems using solution of the smaller sub-problems.
Since the sub-problems are overlapping in nature, in DP approach we avoid recomputing by memorizing answers to repeated sub problems.

The intuition behind dynamic programming is that we trade space for time.
As we are breaking a larger problem into smaller problems, DP is basically recursion with common sense of not recomputing same solution multiple times.

Majority of DP problems can be classified into two categories:
* Optimization: Finding max or min optimal solution.
* Combinatorial: Number of ways something can be done.

All DP problems must have these two properties:
* Optimal substructure: Optimal solution to the problem can be constructed with the optimal solutions to the sub-problems.
* Overlapping subproblems: Subproblems should be repeated so that we can store and reuse their solutions.

# Solving a DP problem

## Recursion:
Solve a problem by breaking it into smaller problems. Extra stack space is used due to function calls.

## Memoization DP:
Storing and reusing past results to reduce computation time is Memoization.
Avoids recomputaion of smaller overlapping sub-problems by storing their solutions for later reuse.
Function calls stack space + mem table space.
Recursion + Memoization = DP technique

1. Identify states.
    For a given subproblem, identify on what parameters does the solution depend. These parameters are the the states and we need to store values of these states in our memoization data structure(usually array or map).
2. Lookup: If value for particular state has already been calculated (look up the MEM table), use this value, avoid recomputation
3. Update: Else, calculate value using recursion and store in table for later use.

## Order of filling mem table:
We can find and store solutions to sub-problems in two ways:
* Top-down: calculating solution and storing it **on demand**.
    Recursion + Memoization = DP technique
    Recursion uses extra stack space.
    Since we solve sub-problems on demand, we avoid computing those solutions that won't contribute to the larger sub-problem.
* Bottom-up: calculating solution and storing it **in advance**. 
    Tabulation.
    Iterative, faster than recursive approach.
    All the sub-problems have to be calculated, which might not be required in some cases.

## Choosing type of mem table:
to-do
* n-d array, but if that is sparsely filled, string to value map would be better

## Tabulation DP
We fill our DP table with solutions to smaller subproblems in such an order that no extra function call is required to compute larger problem's solution.
Bottom-up approach
1. Break into sub-problems.
2. Compute solution to sub-problems.
3. Store into mem table.
4. Use above solution to solve large problems

Start with filling values of base case states.

# Identifying Dynamic Programming types

## Knapsack problem identification
1. You are given N items, for each item either you have to include or exclude. (Example: include in some set)
    No. of times we can include an item, defines type of knapsack problem(0/1, bounded, unbounded)
2. You have some constraint(like remaining bag capacity or remaining subset sum) to decide if item cam be included.

Examples:
* Subset Sum Problem

# Rough Work:
Homework: Use memoization to calculate fibonacci

knapsack explaination https://leetcode.com/discuss/study-guide/1152328/01-Knapsack-Problem-and-Dynamic-Programming
more on knapsack: https://leetcode.com/discuss/study-guide/1200320/Thief-with-a-knapsack-a-series-of-crimes
DP questions: https://leetcode.com/discuss/study-guide/1308617/Dynamic-Programming-Patterns

# References
* DP Newbie to Expert: https://www.youtube.com/playlist?list=PLEJXowNB4kPxBwaXtRO1qFLpCzF75DYrS
