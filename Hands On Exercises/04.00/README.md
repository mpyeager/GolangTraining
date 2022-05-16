# Level 4
## Exercises in this folder
1. [ex01](04.ex01.go)
    > Using a `composite literal`; create an `array` which holds 5 values of type `int`, assign values to each `index position`. `Range` over the array and print the values out. Using `format printing`, print out the `type` of the `array`.
2. [ex02](04.ex02.go)
    > Using a `composite literal`; create a slice of `type` `int`, assign 10 values. `Range` over the slice and print the values out. Using `format printing`, print out the `type` of the slice.
3. [ex03](04.ex03.go)
    > Using the code from the previous exercise, use slicing to create and print the following new slices; `[42 43 44 45 46]`, `[47 48 49 50 51]`, `[45 46 47 48]`, `[43 44 45 46 47]`.
4. [ex04](04.ex04.go)
    > Write a program following these steps; start with the slice `x :=[]int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}`, `append` `38` to the slice and print, in one statement `append 39, 40, 41` and print, `append y := []int{42, 43, 44, 45, 46}` and print.
5. [ex05](04.ex05.go)
    > To `delete` from a slice, use `append` along with slicing. Write a program following these steps; start with slice `x := []int{28, 29, 30, 31, 32, 33, 34, 35, 36, 37}`, use `append` and slicing to derive values `[28, 29, 30, 34, 35, 36, 37]` and print.
6. [ex06](04.ex06.go)
    > Create a slice to store the names of all 25 finalists of Eurovision 2022. Use `make` and `append` to do this. Goal: Do not have the array that underlies the slice created more than once. Print all values including `index` position without using the `range` clause.
7. [ex07](04.ex07.go)
    > Create a slice of a slice of string `[][]string`. Store data in the multi-dimensional slice. `Range` over the records and then `range` of the data in each record.
8. [ex08](04.ex08.go)
    > Create a `map` with a `key` of `type` `string` which is the full name of a Bletchley Park researcher and a value of `type` `[]string` which stores some of their accomplishments. Store 5x records in the `map`, printing out all values with `index` position in the slice.
9. [ex09](04.ex09.go)
    > Using the code from the previous exercise, add a new record to the `map` and print the `map` using `range` loop.
10. [ex10](04.ex10.go)
    > Using the code from the previous exercise, `delete` a record from the `map`. Print the `map` using `range` loop.






