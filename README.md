# Sorting algorithms

The best way to learn a language is to study it in real tasks, or in areas that are not well known. So I resolutely implement the language google go Most of the sorting algorithms.

## TEST 10000 random int elements

```sh
go test -v

=== RUN TestInit
--- PASS: TestInit (0.00 seconds)
=== RUN TestStandartGoogleSort
--- PASS: TestStandartGoogleSort (0.01 seconds)
=== RUN TestSelectionSort
--- PASS: TestSelectionSort (0.70 seconds)
=== RUN TestBubbleSort
--- PASS: TestBubbleSort (1.26 seconds)
=== RUN TestSelectionBubbleSort
--- PASS: TestSelectionBubbleSort (1.06 seconds)
=== RUN TestShakerSort
--- PASS: TestShakerSort (1.19 seconds)
=== RUN TestGnomeSort
--- PASS: TestGnomeSort (0.70 seconds)
=== RUN TestInsertSort
--- PASS: TestInsertSort (1.07 seconds)
=== RUN TestMergeSort
--- PASS: TestMergeSort (0.06 seconds)
PASS
ok      ./sort-algorithms 6.182s
```