# JavaScript note
* Eloquent JavaScript
* 3rd edition (2018)
* Written by Marijn Haverbeke.
* [Book Online](https://eloquentjavascript.net/)
* [Code sandbox and exercise solutions](https://eloquentjavascript.net/code/)
---
- [JavaScript note](#javascript-note)
  - [0. Introduction](#0-introduction)
    - [Overview of this book](#overview-of-this-book)
  - [1.Values, Types, and Operators](#1values-types-and-operators)
    - [Arithmetic](#arithmetic)
    - [Special numbers (NaN, Infinity and -Infinity)](#special-numbers-nan-infinity-and--infinity)
    - [String](#string)
    - [Unary operators](#unary-operators)
    - [Binary operators](#binary-operators)
    - [Logical operators](#logical-operators)
    - [Ternary operator (conditional operator)](#ternary-operator-conditional-operator)
    - [Empty values](#empty-values)
    - [Automatic type conversion (type coercion)](#automatic-type-conversion-type-coercion)
    - [To test whether something refers to the precise value (ex. false)](#to-test-whether-something-refers-to-the-precise-value-ex-false)
    - [Short-circuiting of logical operators](#short-circuiting-of-logical-operators)
    - [Summary](#summary)
---
## 0. Introduction
> JavaScript is ridiculously ***liberal*** in what it allows. It leaves space for a lot of techniques that are impossible in more rigid languages, and as you will see (for example in Chapter 10), it can be used to overcome some of JavaScript’s shortcomings.

> The fact that the language is evolving means that browsers have to constantly keep up, and if you’re using an older browser, it may not support every feature. The language designers are careful to not make any changes that could break existing programs, so new browsers can still run old programs. In this book, I’m using the 2017 version of JavaScript.
>
> Web browsers are not the only platforms on which JavaScript is used. Some databases, such as MongoDB and CouchDB, use JavaScript as their scripting and query language. Several platforms for desktop and server programming, most notably the Node.js project (the subject of Chapter 20), provide an environment for programming JavaScript outside of the browser.
>
### Overview of this book
This book contains roughly three parts. 
* 12 chapters discuss the JavaScript language. 
* 7 chapters are about web browsers and the way JavaScript is used to program them. 
* 2 chapters are devoted to Node.js, another environment to program JavaScript in.

Throughout the book, there are five project chapters. 
* a delivery robot (then error handling and bug fixing, regular expressions, modularity, and asynchronous programming)
* a programming language (concludes the first part)
* a platform game
* a pixel paint program 
* a dynamic website
---
## 1.Values, Types, and Operators
~~Master Yuan-Ma, The Book of Programming do not exist, lol~~
### Arithmetic
\+ \- \* / % (remainder operation)
### Special numbers (NaN, Infinity and -Infinity)
NaN stands for “not a number”, even though it is a value of the number type.
```
 0 / 0
\\ -> NaN
Infinity - Infinity
\\ -> NaN
```
### String
'a string' or "a string" or \`a string\`.

escaping : \    (ex. \n)
```
"\"\\n\"." 
\\ -> "\n".`
```
> Strings, too, have to be modeled as a series of bits to be able to exist inside the computer. The way JavaScript does this is based on the ***Unicode*** standard.
```
`half of 100 is ${100 / 2}`
\\ -> “half of 100 is 50”
```
### Unary operators
* ***typeof*** operator (numbers, strings, Booleans, and undefined values)
```
console.log(typeof 4.5)
// → number
console.log(typeof "x")
// → string
```
### Binary operators
```
console.log(3 > 2)
// → true
console.log("Aardvark" <= "Zoroaster")
// → true
console.log("Itchy" != "Scratchy")
// → true
console.log("Z" < "a")
// -> true
console.log("?" < "Z")
// -> true
```
> The way strings are ordered is roughly alphabetic but not really what you’d expect to see in a dictionary: uppercase letters are always “less” than lowercase ones, so `"Z" < "a"`, and nonalphabetic characters (!, -, and so on) are also included in the ordering. When comparing strings, JavaScript goes over the characters from left to right, comparing the Unicode codes one by one.
```
console.log(NaN == NaN)
// → false
```
> There is **only one** value in JavaScript that is not equal to itself, and that is `NaN` (“not a number”). `NaN` is supposed to denote the result of a nonsensical computation, and as such, it isn’t equal to the result of any other nonsensical computations.
### Logical operators
`and : &&; or : ||; not : !.`

`||` has the lowest precedence, then comes `&&`, then the comparison operators (`>, ==, and so on`), and then the rest. 
### Ternary operator (conditional operator)
```
console.log(true ? 1 : 2);
// → 1
console.log(false ? 1 : 2);
// → 2
```
### Empty values
> There are two special values, written ***null*** and ***undefined***, that are used to denote the absence of a meaningful value. They are themselves values, but they carry no information.

> Many operations in the language that don’t produce a meaningful value (you’ll see some later) yield undefined simply because they have to yield some value.

> The difference in meaning between undefined and null is an accident of JavaScript’s design, and it doesn’t matter most of the time.
### Automatic type conversion (type coercion)
```
console.log(8 * null)
// → 0
console.log("5" - 1)
// → 4
console.log(undefined + 1)
// → 51
console.log(NaN * 2)
// → NaN
console.log(false == 0)
// → true
```
> When an operator is applied to the “wrong” type of value, JavaScript will quietly convert that value to the type it needs, using a set of rules that often aren’t what you want or expect. This is called type coercion. 

> When something that doesn’t map to a number in an obvious way (such as "five" or undefined) is converted to a number, you get the value NaN. Further arithmetic operations on NaN keep producing NaN, so if you find yourself getting one of those in an unexpected place, look for accidental type conversions.
```
console.log(null == undefined);
// → true
console.log(null == 0);
// → false
console.log(undefined == NaN);
// → false
console.log(NaN == NaN);
// → false
```
>  You should get true when both values are the same, except in the case of NaN. But when the types differ, JavaScript uses a complicated and confusing set of rules to determine what to do. In most cases, it just tries to convert one of the values to the other value’s type. However, when null or undefined occurs on either side of the operator, it produces true only if both sides are one of null or undefined.

> That behavior is often useful. When you want to test whether a value has a real value instead of null or undefined, you can compare it to null with the `== (or !=)` operator.

### To test whether something refers to the precise value (ex. false)
> Expressions like `0 == false` and `"" == false` are also true because of automatic type conversion. When you do **not** want any type conversions to happen, there are two additional operators: `=== and !==`.

### Short-circuiting of logical operators
> The logical operators `&&` and `||` handle values of different types in a peculiar way. They will convert the value on their left side to Boolean type in order to decide what to do, but depending on the operator and the result of that conversion, they will return either the original left-hand value or the right-hand value.
```
console.log(null || "user")
// → user
console.log("Agnes" || "user")
// → Agnes
console.log(undefined || null)
// → null
console.log(null || undefined)
// → undefined
```
> The `||` operator, for example, will return the value to its left when that can be converted to true and will return the value on its right otherwise. This has the expected effect when the values are Boolean and does something analogous for values of other types.
```
a = 1
a_mod = a - 1
console.log(a_mod||a)
\\ -> 1
```
> We can use this functionality as a way to **fall back** on a **default value**. If you have a value that might be empty, you can put || after it with a replacement value. If the initial value can be converted to false, you’ll get the replacement instead. The rules for converting strings and numbers to Boolean values state that 0, NaN, and the empty string ("") count as false, while all the other values count as true.
```
console.log(null && "user")
// → null
console.log("Agnes" && "user")
// → user
console.log(0/0 && null)
// → NaN
console.log(undefined && 0/0)
// → undefined
```
> The `&&` operator works similarly but the other way around. When the value to its left is something that converts to false, it returns that value, and otherwise it returns the value on its right.

> Another important property of these two operators is that the part to their right is evaluated only when necessary. In the case of true || X, no matter what X is—even if it’s a piece of program that does something terrible—the result will be true, and X is never evaluated. The same goes for false && X, which is false and will ignore X. This is called **short-circuit evaluation**.

> The conditional operator (ex. `true ? 1 : 2`) works in a similar way. Of the second and third values, only the one that is selected is evaluated.

### Summary
> We looked at four types of JavaScript values in this chapter: numbers, strings, Booleans, and undefined values.

> Such values are created by typing in their name (true, null) or value (13, "abc"). You can combine and transform values with operators. We saw binary operators for arithmetic (+, -, *, /, and %), string concatenation (+), comparison (==, !=, ===, !==, <, >, <=, >=), and logic (&&, ||), as well as several unary operators (- to negate a number, ! to negate logically, and typeof to find a value’s type) and a ternary operator (?:) to pick one of two values based on a third value.



