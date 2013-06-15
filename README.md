# Sorting algorithms

The best way to learn a language is to study it in real tasks, or in areas that are not well known. So I resolutely implement the language google go Most of the sorting algorithms.

# OPTIMIZATIONS

## GnomeSort => OptimizedGnomeSort

Was optimized by the addition of a spot check of the start and half sorted part that allows you to skip a lot of unnecessary checks.

    +--------+------------------+---...---+
    |        |                  |         |
    XxxxxxxxxXxxxxxxxxxxxxxxxxxxXxxxxxxxxxx

### Test with 100,000 items

    === RUN TestStandartGoogleSort
    --- PASS: TestStandartGoogleSort (0.07 seconds)
    === RUN TestGnomeSort
    --- PASS: TestGnomeSort (70.95 seconds)
    === RUN TestOptimizedGnomeSort
    --- PASS: TestOptimizedGnomeSort (0.89 seconds)

# TEST

go test -v
