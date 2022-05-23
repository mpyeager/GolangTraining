## Exercises in this level
1. [ex01](10.ex01.go)
    > Make the code below work. Use a `func literal` [anonymous self-executing func] as well as a buffered channel.
    ```
        func main()  {
          c := make(chan int)
            c <- 42
            fmt.Println(<-c)
          }

    ```
2. [ex02](10.ex02.go)
    > Get the code blocks below to run properly.
    ```
      func main()  {
        cs := make(chan<- int)

        go func() {
          cd <- 42
        }()

      fmt.Println(<-cs)

      fmt.Printf("-------\n")
      fmt.Printf("cs\t%T\n", cs)
    }


      func main()  {
        cr := make(chan<- int)

        go func() {
          cd <- 42
        }()

      fmt.Println(<-cr)

      fmt.Printf("-------\n")
      fmt.Printf("cr\t%T\n", cr)
    }
    ```
3. [ex03](10.ex03.go)
    > Get the code below to run properly.
    ```
    func main()  {
    c := gen()
    receive(c)

    fmt.Println("About to exit.")
    }

    func gen() <-chan int {
      c := make(chan int)

        for i := 0; i < 100; i++ {
          c <-i
        }
      }()
      return c
      }
   ```
  4. [ex04](10.ex04.go)
    > Starting with the code below, pull the values off the channel using the `select` statement.
    ```
    func main() {
      q := make(chan int)
      c := gen(q)

      receive(c, q)

      fmt.Println("About to exit.")
      }

    func gen(q <-chan int) <-chan int {
      c := make(chan int)

      for i := 0; i < 100; i++ {
        c <- i
      }

      return c
      }
    ```
5. [ex05](10.ex05.go)
    > Show the comma `ok` idiom starting with the code below.
    ```
    func main() {
      c := make(chan int)

      v, ok := <-c
      fmt.Println(v, ok)

      close(c)

      v, ok = <-c
      fmt.Println(v, ok)
    }
    ```
6. [ex06](10.ex06.go)
    > Write a program that; puts 100x numbers to a `channel`, pulls the numbers off the `channel` and prints them.

7. [ex07](10.ex07.go)
    > Write a program that; launches 10x goroutines where each goroutine adds 10x numbers to a `channel`, pulls the numbers off the `channel` and prints them.


