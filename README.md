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
Using 𝑀 = 𝐸 − 𝑁 + 2𝑃, with:
- 𝐸 = 14 (edges)
- 𝑁 = 13 (nodes)
- 𝑃 = 1 (connected components)

Result: **Cyclomatic Complexity = 3**

## Asymptotic Complexity

- **Best / Worst / Average Case**: **O(n)**
- **Space Complexity**: **O(log n)** (due to recursion stack)

### Recurrence Relation
**C(n) = 2C(n/2) + 2**
- Base cases:
  - C(2) = 1
  - C(1) = 0

### Master Theorem Application
T(n) = 2T(n/2) + O(1)
### Parameters
- a = 2
- b = 2
- d = 0
- log₂(2) = 1
- Since d < log_b(a), falls under **Case 1**

**Time Complexity: O(n)**
