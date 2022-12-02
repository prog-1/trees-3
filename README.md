# Trees, trees, trees

## 1. Kefa and the park

https://codeforces.com/problemset/problem/580/C

Just reference your solution in the PR.

> **Note**
>
> You need n-trees in this problem (trees that have an arbitraty number of children).

## 2. Duplicates

Create a function that verifies whether there are duplicate elements in a binary tree.

## 3. Level Max

Create a function that finds max element on each tree level of a binary tree.

The function should return a slice, which contains max elements on each level. Slice
indices should be treated as corresponding levels (level 0 corresponds to the tree root).

### Examples

```
Example 1:
     5
   /   \
  2     7
 / \      \
1   3       9
           /
          8
Returns [5 7 9 8]

Example 2:
     5
   /   \
  2     7
 / \
1   3
Returns [5 7 3]
```

## 4. Count Words

Create a program that uses BST to count words in a file.

> **Note**
>
> Use custom split function. You can get inspired by the [standard split implementation].
> And feel free to treat every rune for which [unicode.IsLetter] returns $false$ as a separator.

[standard split implementation]: https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/strings/strings.go;l=236;drc=e7c56fe9948449a3710b36c22c02d57c215d1c10

[unicode.Letter]: https://pkg.go.dev/unicode#IsLetter

## 5. Ingus koeficients

Ingus koeficients problem as described here https://lio.lv/arhivs/arhivs2/2019_3_d1_uzd.pdf

Test data can be taken from here https://lio.lv/arhivs/arhivs2/2019_3_d1_testi.zip
