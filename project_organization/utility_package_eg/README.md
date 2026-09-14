## This

```go
    package stringset

    func New(...string) map[string]struct{} { ... }
    func Sort(map[string]struct{}) []string { ... }

    //Using this package in another file
    set := stringset.New("c", "a", "b")

    fmt.Println(stringset.Sort(set))
```
## Or this

```go
package stringset

type Set map[string]struct{}

func New(...string) Set { ... }
func (s Set) Sort() []string { ... }

//This change makes the client even simpler. There would only be one reference to the stringset package:
set := stringset.New("c", "a", "b")

fmt.Println(set.Sort())
```

