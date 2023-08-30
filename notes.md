This chapter is really more about documentation than anything else

### Named Result Parameters

- Use if:

  - a function returns two or three parameters of the same type, or
  - the meaning of a result isn't clear from context

    ```go
    func (f *Foo) Location() (float64, float64, error)
    ```

    is less clear than:

    ```go
    // Location returns f's latitude and longitude.
    // Negative values mean south and west, respectively.
    func (f *Foo) Location() (lat, long float64, err error)
    ```

    but the following would be repetitive

    ```go
    func (n *Node) Parent1() (node *Node) {}
    func (n *Node) Parent2() (node *Node, err error) {}
    ```

    better to use:

    ```go
    func (n *Node) Parent1() *Node {}
    func (n *Node) Parent2() (*Node, error) {}
    ```

- Fun fact: A `return` statement without arguments returns the named return values. This is known as a "naked" return.
  ```go
  func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
      return
  }
  ```
  - Naked returns are okay if the function is a handful of lines. Once it's a medium sized function, be explicit with your return values. Corollary: it's not worth it to name result parameters just because it enables you to use naked returns. Clarity of docs is always more important than saving a line or two in your function.

### Godoc Examples

- Often code examples that can be found outside the codebase, such as a readme file often become out of date and incorrect compared to the actual code because they don't get checked.
- Examples are snippets of Go code that are displayed as package documentation and that are verified by running them as tests
- They can also be run by a user visiting the godoc web page for the package and clicking the associated “Run” button.
- Go examples are executed just like tests so we can be confident examples reflect what the code actually does.
- Running the package’s test suite, we can see the example function is executed

```zsh
$ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
```

- Output comments:

  - As it executes the example, the testing framework captures data written to standard output and then compares the output against the example’s “Output:” comment. The test passes if the test’s output matches its output comment.
  - To see a failing example we can change the output comment text to something obviously incorrect

    ```go
    func ExampleAdd() {
        sum := Add(1, 4)
        fmt.Println(sum)
        // Output: 9
    }
    ```

    and run the tests again:

    ```zsh
    $ go test
    === RUN   TestAdder
    --- PASS: TestAdder (0.00s)
    === RUN   ExampleAdd
    --- FAIL: ExampleAdd (0.00s)
    got:
    5
    want:
    9
    FAIL
    exit status 1
    FAIL    integers        0.004s
    ```

    If we remove the output comment entirely

    ```go
    func ExampleAdd() {
      sum := Add(1, 4)
      fmt.Println(sum)
    }
    ```

    then the example function is compiled but not executed:

    ```zsh
    $ go test -v
    === RUN   TestAdder
    --- PASS: TestAdder (0.00s)
    PASS
    ok      integers        0.003s
    ```

  - Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles.

### Godoc Documentation Server

- To start the documentation server
  - Install godoc: `go install golang.org/x/tools/cmd/godoc@latest`
  - Run this command: `godoc -http=:6060`
- godoc is aware of modules, no need to worry about `GOPATH`
- We can access this package documentation through http://localhost:6060/pkg/integers/
