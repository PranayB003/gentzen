# Gentzen Proof Theory
Sequent calculus (Gentzen system) implementation in GO to check the validity of
a propositional formula, done as part of an assignment for Mathematical Foundations of Computer Science (CS6L015).

## Installation
To compile and execute this program, you need to have [Go](https://go.dev/) installed. You can do this by following the [installation instructions](https://go.dev/dl/) on the official website, or by using a package manager like [Homebrew](https://brew.sh/):
```bash
brew install --formula go
```
Next, clone the repository:
```bash
git clone https://github.com/PranayB003/gentzen.git ./
cd gentzen
```

## Usage
Compile and run the program as follows:
```bash
go build -o gzn ./
./gzn
```
Or alternatively, 
```bash
go run .
```

### Syntax
The following connectives are supported:
Symbol|Logical Connective
---|---:
!|Not ($\neg$)
&|And ($\land$)
\||Or ($\lor$)
->|Implication ($\implies$)
<->|Double Implication ($\iff$)
TRUE|True ($\top$)
FALSE|False ($\perp$)

In addition, parenthesis `(`, `)` can be used to specify precedence. 
> [!NOTE]
> All propositional symbols, connectives, and parenthesis symbols **must be separated by spaces** for the expression to be recognised properly by the program.

For more information, see [Examples](#examples) below.

### Options
```
--help          : Display available CLI options and other usage information
--debug-parse	: Display tokens and parsed expression
--validity-only	: Omit proof tree and sequents from output
```

### Examples
- Using the `--validity-only` option
```
$ ./gzn --validity-only
>> ( a -> c ) -> ( ( b -> c ) -> ( ( a | b ) -> c ) )
INVALID
>>
>> a -> a
VALID
>>
```
- Using the `--debug-parse` option
```
$ ./gzn --debug-parse
>> ( p & ! q ) -> ( p <-> q )
tokens:  [( p & ! q ) -> ( p <-> q )]
expression:  (((p) & (! (q))) -> (((p) -> (q)) & ((q) -> (p))))

Proof Tree:
 => (((p) & (! (q))) -> (((p) -> (q)) & ((q) -> (p))))
((p) & (! (q))) => (((p) -> (q)) & ((q) -> (p)))
((p) & (! (q))) => ((p) -> (q))     ((p) & (! (q))) => ((q) -> (p))
((p) & (! (q))), (p) => (q)     ((p) & (! (q))), (q) => (p)
(p), (p), (! (q)) => (q)     (q), (p), (! (q)) => (p)
(p), (p) => (q), (q)     (q), (p) => (p), (q)

Final Sequents:
(p), (p) => (q), (q)     (q), (p) => (p), (q)

INVALID
```
- Without any other options
```
$ ./gzn
>> ! p | q -> p -> q
Proof Tree:
 => (((! (p)) | (q)) -> ((p) -> (q)))
((! (p)) | (q)) => ((p) -> (q))
((! (p)) | (q)), (p) => (q)
(p), (! (p)) => (q)     (p), (q) => (q)
(p) => (q), (p)

Final Sequents:
(p), (q) => (q)     (p) => (q), (p)

VALID
```
