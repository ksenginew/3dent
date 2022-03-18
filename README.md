# 3dent

A  **lightweight**  runtime for  **JavaScript**  and  **TypeScript**.

3dent is a simple runtime for JavaScript and TypeScript that uses your preinstalled web browser and a server written in Go-lang.

### Installation

3dent ships as a single `.go` with no external dependencies. You need to [install Go compiler](https://go.dev/doc/install) to run it. Then just download [`main.go`](https://github.com/ksenginew/3dent/blob/main/main.go).

### Getting Started

Try running a simple program:

```sh
go run main.go https://deno.land/std/examples/welcome.ts
```

Then open the url you got.

> Note: You can pass a local file name instead of an URL.

### Examples

Here are some examples that you can use to get started immediately.

1.  [Hello World](https://github.com/ksenginew/3dent/blob/main/examples/hello-world)
2.  [Importing & Exporting](https://github.com/ksenginew/3dent/blob/main/examples/import-export)

For more examples, check out  [examples folder](https://examples.deno.land/).
