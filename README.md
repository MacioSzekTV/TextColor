# TextColor 🎨 <a href="https://github.com/MacioSzekTV/TextColor"><img src="https://img.shields.io/static/v1?color=5865F2&logo=go&label=TextColor&style=flat-square&message=1.0.0"/></a>

A simple Go library for coloring text in the terminal.
It provides an easy way to wrap strings with ANSI escape codes for colored output.

---

## Table of Contents

* [About](#about)
* [Installation](#installation)
* [Usage](#usage)
* [Examples](#examples)
* [Available Colors](#available-colors)
* [License](#license)
* [Contact](#contact)

---

## About

TextColor is a lightweight Go package that allows you to add colors to your terminal output.
It’s useful for improving the readability of CLI tools, logs, or any console-based output.

The library is designed to be simple, intuitive, and easy to use.

---

## Installation

Install the package using `go get`:

```bash
go get github.com/MacioSzekTV/TextColor
```

Then import it in your Go code:

```go
import "github.com/MacioSzekTV/TextColor"
```

---

## Usage

```go
package main

import (
    "github.com/MacioSzekTV/TextColor"
)

func main() {
    TextColor.GREEN("YOUR TEXT")
}
```

---

## Examples

```go
TextColor.RED("Error: something went wrong")
TextColor.GREEN("Success!")
TextColor.YELLOW("Warning!")
TextColor.MAGENTA("Debugging")
```
![TextColor Example](https://i.imgur.com/kwE4Gtp.png)

---

## Available Colors

- <span style="color:black">BLACK</span>  
- <span style="color:red">RED</span>  
- <span style="color:green">GREEN</span>  
- <span style="color:yellow">YELLOW</span>  
- <span style="color:purple">PURPLE</span>  
- <span style="color:magenta">MAGENTA</span>  
- <span style="color:teal">TEAL</span>  
- <span style="color:white">WHITE</span>  


---

## License

This project is licensed under the **[GPL-3.0](./LICENSE)** license.

---

## Contact

* Author: **Maciej (MacioSzekTV)**
* GitHub: [MacioSzekTV/TextColor](https://github.com/MacioSzekTV/TextColor)
* For questions or suggestions, open an **issue** on GitHub
