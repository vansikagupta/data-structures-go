**Arrays**

Arrays are widely used and form the building block for most other data structures.

Array is a fixed length container that holds n items that are indexable. Elements of an are contigously stored in memeory.\
Dynamic arrays can grow and shrink in size.

**Arrays in Go**
* The type of element and length are both part of the array’s type.
* Arrays are values. Assigning one array to another copies all the elements.
* Slices are used more often than arrays. Slices hold references to an underlying array.

Slices wrap arrays to give a more convenient and powerful interface to sequences. They are the dynamic arrays. 
Slices are typed by the type of elements alone and not the number of elements. 

It is the run-time data structure holding the pointer, length, and capacity.

len(s) : The length of a slice is the number of elements it contains. \
cap(s) : The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

*Example*
s1 := []int{1,2,3,4,5,6}
s2 := s1[2:5]
fmt.Println(s2) // [3,4,5]
len(s) // 3
cap(s) // 4
 
append the slice is reallocated when the data exceeds capacity


**When and where to use arrays?**
* store and access sequential data
* heavily used in dynamic programming to cache sub-problem solutions.

**Complexity of common operations**\
Access any element - O(1)\
Linear Search - O(n)\
Binary Search - O(logn)\
Append - O(1)\
Insertion - O(n)\
Deletion - O(n)

Insert and Delete are expensive operations as they require to shift the elements and potentially need to re-copy elements into new larger array.

**Common array problems with solution**