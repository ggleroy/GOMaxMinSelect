# GOMaxMinSelect  
Max-Min Select algorithm using Golang!  
A **Foundations of Algorithm Design and Analysis** assignment. 

## About  
The **Max-Min algorithm** is an efficient method for finding the maximum and minimum values in an array using **divide-and-conquer**,
 reducing the number of comparisons from 2n − 2 (naive approach) to approximately 1.5n. It recursively splits the array into halves, 
computes local max/min values, and merges results with just two comparisons per level. This makes it ideal for **large datasets,
 performance-critical systems, and resource-constrained applications.**
 
## Structure  
This repository consists of a single **Go** file: `main.go`.  

## How to Run  
Ensure **Go** is installed on your machine, then execute:  

```sh
go run GOMaxMinSelect.go
```

## Cyclomatic Complexity  
Using the formula 𝑀 = 𝐸 − 𝑁 + 2𝑃, with:  
- **𝐸** (edges) = 14  
- **𝑁** (nodes) = 13
- **𝑃** (connected components) = 1   

The **cyclomatic complexity** of the Max-Min algorithm is **3**.  

## Asymptotic Complexity Analysis: Operation Counting Method

### Algorithm Overview
The algorithm finds the maximum and minimum elements by recursively dividing the array and combining results efficiently.

### Recursive Division Strategy
* **Input Size Reduction**: Array divided into two equal (or nearly equal) halves
* **Division Cost**: Constant time O(1)

### Comparison Breakdown

#### Base Cases
* **Single Element**: 0 comparisons
* **Two Elements**: 1 comparison to determine max and min

#### Recursive Case (n > 2)
* Divide array into two halves
* Recursively find max and min in each half
* Combine results with 2 comparisons

### Recurrence Relation
```
C(n) = 2 * C(n/2) + 2, for n > 2
C(2) = 1
C(1) = 0
```

### Detailed Comparison Analysis

#### Comparison Counting
* Each recursive level contributes a constant number of comparisons
* Number of levels is log₂(n)
* Comparisons at each level remain constant (2 per subproblem)

### Complexity Conclusion
* **Time Complexity**: O(n)
* **Space Complexity**: O(log n) due to recursive call stack
