# gg

A fast, type-safe HTML generator for Go.

## Features

- **Auto-escaping** - Strings are HTML-escaped automatically for security
- **Type-safe** - Compile-time checking of your HTML structure
- **Zero dependencies** - Pure Go standard library
- **Fast** - Minimal allocations, direct writer output
- **Composable** - Build complex layouts from simple components

## Installation

```bash
go get github.com/assaidy/gg
```

## Quick Start

```go
package main

import (
    "os"
    "github.com/assaidy/gg"
)

func main() {
    page := gg.Empty(
        gg.DoctypeHTML(),
        gg.Html(
            gg.Head(
                gg.Title("My Page"),
            ),
            gg.Body(
                gg.H1("Hello, World!"),
                gg.P("Auto-escaped: <script>alert('xss')</script>"),
            ),
        ),
    )
    
    if err := page.Render(os.Stdout); err != nil {
        panic(err)
    }
}
```

## Usage

### Basic Elements

```go
// Strings are auto-escaped
gg.Div("Hello", " ", "World")  // <div>Hello World</div>

gg.P("<script>alert('xss')</script>")
// <p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>

// Raw HTML (not escaped. use with caution)
gg.Div(gg.RawHTML("<svg>...</svg>")) // <svg>...</svg>

// Numbers and booleans are auto-converted
gg.P("Count: ", 42)           // <p>Count: 42</p>
gg.P("Active: ", true)        // <p>Active: true</p>
```

### Attributes

```go
gg.Div(gg.KV{"class": "container", "id": "main"}, "Content")
// <div class="container" id="main">Content</div>
```

### Conditional Rendering

```go
import "github.com/assaidy/gg/utils"

// Show element only if condition is true
ggu.If(isLoggedIn, gg.Div("Welcome back!"))

// Choose between two options
ggu.IfElse(isAdmin, gg.Div("Admin"), gg.Div("User"))
```

### Lists and Iteration

```go
items := []string{"Apple", "Banana"}

// Map over slice
gg.Ul(
    ggu.Map(items, func(item string) gg.Node {
        return gg.Li(item)
    }),
)

// Repeat N times
gg.Div(
    ggu.Repeat(3, func() gg.Node {
        return gg.P("Repeated")
    }),
)
```

### With Tailwind CSS

```go
gg.Div(gg.KV{"class": "bg-gray-100 min-h-screen p-8"},
    gg.Div(gg.KV{"class": "max-w-4xl mx-auto"},
        gg.H1(gg.KV{"class": "text-4xl font-bold text-gray-800"}, "Title"),
        gg.P(gg.KV{"class": "text-gray-600 mt-2"}, "Description"),
        gg.Button(
            gg.KV{"class": "px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"},
            "Click Me",
        ),
    ),
)
```

### With HTMX

```go
// HTMX button that loads content
gg.Button(
    gg.KV{
        "class":     "px-4 py-2 bg-blue-500 text-white rounded",
        "hx-get":    "/api/users",
        "hx-target": "#users-list",
        "hx-swap":   "outerHTML",
    },
    "Load Users",
)

// HTMX form
gg.Form(
    gg.KV{
        "hx-post":   "/api/submit",
        "hx-target": "#result",
        "class":     "space-y-4",
    },
    gg.Input(gg.KV{
        "type":  "text",
        "name":  "message",
        "class": "border rounded px-3 py-2",
    }),
    gg.Button(
        gg.KV{"type": "submit", "class": "bg-blue-500 text-white px-4 py-2 rounded"},
        "Submit",
    ),
)
```

### Complete Example

```go
package main

import (
    "os"
    "github.com/assaidy/gg"
    "github.com/assaidy/gg/utils"
)

func main() {
    users := []string{"Alice", "Bob", "Charlie"}
    isAdmin := true

    page := gg.Empty(
        gg.DoctypeHTML(),
        gg.Html(gg.KV{"lang": "en"},
            gg.Head(
                gg.Title("Dashboard"),
                gg.Script(gg.KV{"src": "https://cdn.tailwindcss.com"}),
            ),
            gg.Body(gg.KV{"class": "bg-gray-100 p-8"},
                gg.Div(gg.KV{"class": "max-w-2xl mx-auto"},
                    gg.H1(gg.KV{"class": "text-3xl font-bold mb-4"}, "Dashboard"),
                    
                    // Conditional admin panel
                    ggu.If(isAdmin,
                        gg.Div(gg.KV{"class": "bg-blue-50 p-4 rounded mb-4"},
                            gg.P(gg.KV{"class": "font-semibold"}, "Admin Panel"),
                        ),
                    ),
                    
                    // User count
                    gg.P("Total users: ", len(users)),
                    
                    // User list
                    gg.Ul(gg.KV{"class": "space-y-2 mt-4"},
                        ggu.Map(users, func(name string) gg.Node {
                            return gg.Li(
                                gg.KV{"class": "p-2 bg-white rounded shadow"},
                                name,
                            )
                        }),
                    ),
                ),
            ),
        ),
    )

    if err := page.Render(os.Stdout); err != nil {
        panic(err)
    }
}
```
