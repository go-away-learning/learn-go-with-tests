# Learn Go with Tests

<p align="center">
  <figure>
  <img src=".github/red-green-blue-gophers-smaller.png" />
  <figcaption>

Inspired by the TDD approach, our diligent gophers demonstrate each phase: The eager red gopher envisions a delectable cupcake, the focused green gopher mixes ingredients to pass the taste test, while the thoughtful blue gopher tidies up the kitchen, symbolizing the iterative Red, Green, Refactor cycle of development. Art by: [Denise](https://twitter.com/deniseyu21)

  </figcaption>
  </figure>
</p>

## What

This is my public journey towards getting better at Go by reading [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go).

I'll be registering every chapter in a branch named after it's corresponding chapter.

## Why

- Explore the Go language by writing tests
- Get a grounding with TDD. Go is a good language for learning TDD because it is a simple language to learn and testing is built-in
- Be confident that you'll be able to start writing robust, well-tested systems in Go

## Who this is for

- People who are interested in picking up Go.
- People who already know some Go, but want to explore testing with TDD.

## What you'll need

- A computer!
- [Installed Go](https://golang.org/)
- A text editor
- Some experience with programming. Understanding of concepts like `if`, variables, functions etc.
- Comfortable with using the terminal

## Table of contents

To make this easier to follow for future me, I'll be using [this specific tree](https://github.com/quii/learn-go-with-tests/tree/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa) from the source repository. Every resource source here will be linking to that `tree`.

### Go fundamentals

1. [Install Go](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/install-go.md) - Set up environment for productivity.
2. [Hello, world](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/hello-world.md) - Declaring variables, constants, if/else statements, switch, write your first go program and write your first test. Sub-test syntax and closures.
3. [Integers](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/integers.md) - Further Explore function declaration syntax and learn new ways to improve the documentation of your code.
4. [Iteration](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/iteration.md) - Learn about `for` and benchmarking.
5. [Arrays and slices](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/arrays-and-slices.md) - Learn about arrays, slices, `len`, varargs, `range` and test coverage.
6. [Structs, methods & interfaces](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/structs-methods-and-interfaces.md) - Learn about `struct`, methods, `interface` and table driven tests.
7. [Pointers & errors](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/pointers-and-errors.md) - Learn about pointers and errors.
8. [Maps](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/maps.md) - Learn about storing values in the map data structure.
9. [Dependency Injection](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/dependency-injection.md) - Learn about dependency injection, how it relates to using interfaces and a primer on io.
10. [Mocking](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/mocking.md) - Take some existing untested code and use DI with mocking to test it.
11. [Concurrency](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/concurrency.md) - Learn how to write concurrent code to make your software faster.
12. [Select](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/select.md) - Learn how to synchronise asynchronous processes elegantly.
13. [Reflection](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/reflection.md) - Learn about reflection
14. [Sync](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/sync.md) - Learn some functionality from the sync package including `WaitGroup` and `Mutex`
15. [Context](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/context.md) - Use the context package to manage and cancel long-running processes
16. [Intro to property based tests](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/roman-numerals.md) - Practice some TDD with the Roman Numerals kata and get a brief intro to property based tests
17. [Maths](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/math.md) - Use the `math` package to draw an SVG clock
18. [Reading files](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/reading-files.md) - Read files and process them
19. [Templating](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/html-templates.md) - Use Go's html/template package to render html from data, and also learn about approval testing
20. [Generics](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/generics.md) - Learn how to write functions that take generic arguments and make your own generic data-structure
21. [Revisiting arrays and slices with generics](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/revisiting-arrays-and-slices-with-generics.md) - Generics are very useful when working with collections. Learn how to write your own `Reduce` function and tidy up some common patterns.

### Build an application

- [HTTP server](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/http-server.md) - We will create an application which listens to HTTP requests and responds to them.
- [JSON, routing and embedding](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/json.md) - We will make our endpoints return JSON and explore how to do routing.
- [IO and sorting](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/io.md) - We will persist and read our data from disk and we'll cover sorting data.
- [Command line & project structure](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/command-line.md) - Support multiple applications from one code base and read input from command line.
- [Time](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/time.md) - using the `time` package to schedule activities.
- [WebSockets](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/websockets.md) - learn how to write and test a server that uses WebSockets.

### Testing fundamentals

- [Introduction to acceptance tests](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/intro-to-acceptance-tests.md) - Learn how to write acceptance tests for your code, with a real-world example for gracefully shutting down a HTTP server
- [Scaling acceptance tests](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/scaling-acceptance-tests.md) - Learn techniques to manage the complexity of writing acceptance tests for non-trivial systems.

### Questions and answers

- [OS exec](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/os-exec.md) - An example of how we can reach out to the OS to execute commands to fetch data and keep our business logic testable/
- [Error types](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/error-types.md) - Example of creating your own error types to improve your tests and make your code easier to work with.
- [Context-aware Reader](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/context-aware-reader.md) - Learn how to TDD augmenting `io.Reader` with cancellation. Based on [Context-aware io.Reader for Go](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer)
- [Revisiting HTTP Handlers](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/http-handlers-revisited.md) - Testing HTTP handlers seems to be the bane of many a developer's existence. This chapter explores the issues around designing handlers correctly.

### Meta / Discussion

- [Why unit tests and how to make them work for you](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/why.md) - Watch a video, or read about why unit testing and TDD is important
- [Anti-patterns](https://github.com/quii/learn-go-with-tests/blob/61b4b6e0ffb9e655df5fe499b254e1b6cdbf79aa/anti-patterns.md) - A short chapter on TDD and unit testing anti-patterns

[MIT license](LICENSE.md)
