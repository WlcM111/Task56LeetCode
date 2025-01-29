# Merge Intervals

This program merges overlapping intervals from a given list of intervals. Each interval is represented as a pair of integers `[start, end]`. The program ensures that all overlapping intervals are combined into a single interval.

## Function

- **mergeIntervals(intervals [][]int) [][]int**: Merges overlapping intervals and returns a list of non-overlapping intervals.

## How It Works

1. **Sorting**:
   - The intervals are first sorted based on their start values using `sort.Slice`. This ensures that overlapping intervals are adjacent in the list.

2. **Merging**:
   - A result list is initialized with the first interval.
   - The program iterates through the sorted intervals starting from the second one.
   - If the current interval overlaps with the last interval in the result list (i.e., the start of the current interval is less than or equal to the end of the last interval in the result), the intervals are merged by updating the end value of the last interval in the result list to the maximum of the two end values.
   - If there is no overlap, the current interval is added to the result list as a new interval.

3. **Return**:
   - The program returns the list of merged, non-overlapping intervals.

## Example

For the input `[[1, 3], [2, 6], [8, 10], [15, 18]]`, the program will return `[[1, 6], [8, 10], [15, 18]]`, as the intervals `[1, 3]` and `[2, 6]` overlap and are merged into `[1, 6]`.
