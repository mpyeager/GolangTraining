## Exercises in this level
1. [ex01](09.ex01.go)
   > In addition to the `main` goroutine, launch 2x additional goroutines with each net additional goroutine printing something out. Use `waitgroup`s to ensure each goroutine finishes before your program exits.

2. [ex02](09.ex02.go)
   > Create a `type person struct` and attach a `method` `speak` to `type` `person` using a pointer receiver. Create a `type human interface`. Create `func saySomething` and have it take in a human as a parameter which calls the `speak` method.
3. [ex03](09.ex03.go)
   > Using goroutines, create an `incrementer` program. Have a `variable` to hold the `incrementer` value, and launch several goroutines with each goroutine; reading the incrementer value, store in a new `variable`, yield the processor with `runtime.Gosched()`, increment the new `variable`, write the value in the new `variable` back to the `incrementer variable`. Use `waitgroup`s to ensure all goroutines finish. Prove that you've just created a race condition using the `-race` flag.
4. [ex04](09.ex04.go)
   > Fix the `race` condition you created in the previous exercise by using a `mutex`.
5. [ex05](09.ex05.go)
   > Fix the `race` condition you created in [ex03](09.ex03.go) by using `package` `atomic`.
6. [ex06](09.ex06.go)
   > Create a program that prints out your `OS` and `ARCH`.




