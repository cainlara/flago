# Flago ⚑

A lightweight and flexible flag/arguments handler for Go applications.

---
## 📚 Content

- [🚀 Features](#-features)
- [📦 Installation](#-installation)
- [🧩 Usage](#-usage)
- [🧪 Testing](#-testing)
- [📄 Documentation](#-documentation)
- [🛠️ Roadmap & Contribution](#️-roadmap--contribution)
- [🧾 License](#-license)
- [📌 About](#-about)
---

## 🚀 Features

- 📦 Intuitive command-line flag parsing  
- 🎯 Supports short (`-f`) options  
- 💡 Simple API for defining flags with default values and descriptions  
- 🧪 Includes built-in tests to demonstrate usage  
- 🛡️ Licensed under BSD‑3‑Clause  

---

## 📦 Installation

```bash
go get github.com/cainlara/flago
```

---

## 🧩 Usage

Import the module and define flags straightforwardly:

### Getting arguments as a map

```go
package main

import (
  "log"

  "github.com/cainlara/flago"
)

func main() {
  argsAsMap, err := flago.GetArgsMap()
  if err != nil {
    log.Fatal(err)
  }

  for argName, value := range argsAsMap {
    log.Printf("Arg Name: %s, Value: %s", argName, value)
  }
}
```
Then, run it with
```bash
go run main.go alice 1 bob 2 charlie 3
```
And the expected output should look like:
```bash
Arg Name: alice, Value: 1
Arg Name: bob, Value: 2
Arg Name: charlie , Value: 3
```

### Getting arguments as known structure

```go
package main

import (
  "log"

  "github.com/cainlara/flago"
)

type Config struct {
  Source string
  Limit  int
  Skip   bool
}

func main() {
  var customStruct Config

  err = flago.GetArgsStruct(&customStruct, false)
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("Output: %v", customStruct)
}
```
Then, run it with
```bash
go run main.go "source" "path/to/file" "limit" 100 "skip" true
```
And the expected output should look like:
```bash
Output: {path/to/file 100 false}
```

---

## 🧪 Testing

Unit tests are included. To run them:

```bash
go test ./... 
```

---

## 📄 Documentation

- Explore `main.go` for full API details and usage patterns  
- Function-level documentation is embedded in source ahead of upcoming enhancements  

---

## 🛠️ Roadmap & Contribution

### Planned Updates

- ✅ No further plans.

---

## 🧾 License

Distributed under the BSD‑3‑Clause License. See [LICENSE](./LICENSE) for details.

---

## 📌 About

`flago` provides a clean and efficient approach to command-line flag parsing in Go, helpful for small- to mid-scale tools and scripts.

**Latest release:** v0.1.0 (July 7, 2025)  
**Author:** Jose Andrés Lara Vecino (cainlara) ([github.com](https://github.com/cainlara))