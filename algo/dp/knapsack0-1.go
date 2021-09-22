/*** Problem ***
Given n items with associated weights and profits and a bag with capacity S,
fill the bag such that it gives maximum profit.
Items:   0 1 2
Profit:  6 8 7
Weight:  3 2 4
Capacity: 8

*** Solution ***
* For each item you have two choices, either to add to bag or exclude.
* Condition 1: You cannot add more items if you have no remaining capacity OR no more items left.
* Condition 2: if current item has weight > remaining capacity, you have to skip the item
* Condition 3: if current item has weight < remaining capacity,
 			you may or may not include item depending on which choice gives you max profit.
			i.e, max(inlude, exclude)

Recursive approach:
Memoization:
Tabulation:
*/

package dp
