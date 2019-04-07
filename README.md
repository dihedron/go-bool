# go-bool: a simple implementation of boolean logic in Golang

This library provides support for evaluation of boolean constructs using a simple, flowing syntax such as:

```golang
expr := And(
            Or(
                Expression1(),
                Expression2(),
                Not(
                    ExpressionN(),
                ),
            ),
        )
result, err := expr.Evaluate(nil)
```

Apart from the core boolean operators (`And`, `Not`, `Or`, `Xor`) and the boolean value (`True` and `False`), you can create your own `Operand`s by defining a type that implements the `Evaluate` method:

```golang
type MyOperand struct {}

func (o MyOperand)Evaluate(ctx interface{})(bool, error) {
    // do your processing here
    return result, nil
}
```

When an error is returned by any nested `Operand` or `Operator`, it must bubble to the top, and the boolean value should be ignored.

`ctx` is an opaque context; it is up to the caller to provide a meaningful object if necessary; default `Operand`s and `Operator`s ignore it and simply propagate it as-is to nested expressions.

## TODO

- [ ] Support for parsing from and serialisation to a simple DSL of boolean expressions
- [ ] Support for parsing from and serialisation to JSON of boolean expressions
- [ ] Support for parsing from and serialisation to YAML of boolean expressions
- [ ] Support for parsing from and serialisation to XML of boolean expressions