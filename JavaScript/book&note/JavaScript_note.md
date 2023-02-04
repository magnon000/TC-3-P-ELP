# JavaScript note
* Eloquent JavaScript
* 3rd edition (2018)
* Written by Marijn Haverbeke.
* [Book Online](https://eloquentjavascript.net/)
* [Code sandbox and exercise solutions](https://eloquentjavascript.net/code/)
---
- [JavaScript note](#javascript-note)
  - [0.Introduction](#0introduction)
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
  - [2.Program Structure](#2program-structure)
    - [Expressions and statements](#expressions-and-statements)
    - [Bindings](#bindings)
    - [Binding names](#binding-names)
    - [The environment](#the-environment)
    - [Functions](#functions)
    - [The console.log function](#the-consolelog-function)
    - [Return values](#return-values)
    - [Control flow](#control-flow)
    - [Conditional execution](#conditional-execution)
    - [while and do loops](#while-and-do-loops)
    - [Indenting Code](#indenting-code)
    - [for loops](#for-loops)
    - [Breaking Out of a Loop](#breaking-out-of-a-loop)
    - [Updating bindings succinctly](#updating-bindings-succinctly)
    - [Dispatching on a value with switch](#dispatching-on-a-value-with-switch)
    - [Capitalization](#capitalization)
    - [Comments](#comments)
    - [Summary](#summary-1)
    - [3 exercices](#3-exercices)
  - [3.Functions](#3functions)
    - [Defining a function](#defining-a-function)
    - [Bindings and scopes](#bindings-and-scopes)
    - [Nested scope](#nested-scope)
    - [Functions as values](#functions-as-values)
    - [Declaration notation](#declaration-notation)
    - [Arrow functions](#arrow-functions)
    - [The call stack](#the-call-stack)
    - [Optional Arguments](#optional-arguments)
    - [Closure](#closure)
    - [Recursion](#recursion)
    - [Growing functions](#growing-functions)
    - [Functions and side effects](#functions-and-side-effects)
    - [Summary](#summary-2)
    - [3 exercices](#3-exercices-1)
  - [4.Data Structures: Objects and Arrays](#4data-structures-objects-and-arrays)
    - [Properties](#properties)
    - [Methods](#methods)
    - [Objects](#objects)
    - [Mutability](#mutability)
    - [Array loops](#array-loops)
    - [Further arrayology](#further-arrayology)
    - [Strings and their properties](#strings-and-their-properties)
    - [Rest parameters](#rest-parameters)
    - [The Math object](#the-math-object)
    - [JSON](#json)
    - [Summary](#summary-3)
    - [4 exercices](#4-exercices)
  - [5.Higher-Order Functions](#5higher-order-functions)
    - [Abstraction](#abstraction)
    - [Abstracting repetition](#abstracting-repetition)
    - [Higher-order functions](#higher-order-functions)
    - [Script data set](#script-data-set)
---
## 0.Introduction
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
`NaN` stands for “not a number”, even though it is a value of the number type.
```
 0 / 0
\\ -> NaN
Infinity - Infinity
\\ -> NaN
```
### String
`'a string'` or `"a string"` or \`a string\`.

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
// -> number
console.log(typeof "x")
// -> string
```
### Binary operators
```
console.log(3 > 2)
// -> true
console.log("Aardvark" <= "Zoroaster")
// -> true
console.log("Itchy" != "Scratchy")
// -> true
console.log("Z" < "a")
// -> true
console.log("?" < "Z")
// -> true
```
> The way strings are ordered is roughly alphabetic but not really what you’d expect to see in a dictionary: uppercase letters are always “less” than lowercase ones, so `"Z" < "a"`, and nonalphabetic characters (!, -, and so on) are also included in the ordering. When comparing strings, JavaScript goes over the characters from left to right, comparing the Unicode codes one by one.
```
console.log(NaN == NaN)
// -> false
console.log(NaN === NaN);
// -> false
```
> There is **only one** value in JavaScript that is not equal to itself, and that is `NaN` (“not a number”). `NaN` is supposed to denote the result of a nonsensical computation, and as such, it isn’t equal to the result of any other nonsensical computations.
### Logical operators
`and: &&; or: ||; not: !.`

`||` has the lowest precedence, then comes `&&`, then the comparison operators (`>, ==, and so on`), and then the rest. 
### Ternary operator (conditional operator)
```
console.log(true ? 1 : 2);
// -> 1
console.log(false ? 1 : 2);
// -> 2
```
### Empty values
> There are two special values, written ***null*** and ***undefined***, that are used to denote the absence of a meaningful value. They are themselves values, but they carry no information.

> Many operations in the language that don’t produce a meaningful value (you’ll see some later) yield undefined simply because they have to yield some value.

> The difference in meaning between undefined and null is an accident of JavaScript’s design, and it doesn’t matter most of the time.
### Automatic type conversion (type coercion)
```
console.log(8 * null)
// -> 0
console.log("5" - 1)
// -> 4
console.log(undefined + 1)
// -> 51
console.log(NaN * 2)
// -> NaN
console.log(false == 0)
// -> true
```
> When an operator is applied to the “wrong” type of value, JavaScript will quietly convert that value to the type it needs, using a set of rules that often aren’t what you want or expect. This is called type coercion. 

> When something that doesn’t map to a number in an obvious way (such as "five" or undefined) is converted to a number, you get the value NaN. Further arithmetic operations on NaN keep producing NaN, so if you find yourself getting one of those in an unexpected place, look for accidental type conversions.
```
console.log(null == undefined);
// -> true
console.log(null == 0);
// -> false
console.log(undefined == NaN);
// -> false
console.log(NaN == NaN);
// -> false
```
>  You should get true when both values are the same, except in the case of NaN. But when the types differ, JavaScript uses a complicated and confusing set of rules to determine what to do. In most cases, it just tries to convert one of the values to the other value’s type. However, when null or undefined occurs on either side of the operator, it produces true only if both sides are one of null or undefined.

> That behavior is often useful. When you want to test whether a value has a real value instead of null or undefined, you can compare it to null with the `== (or !=)` operator.

### To test whether something refers to the precise value (ex. false)
> Expressions like `0 == false` and `"" == false` are also true because of automatic type conversion. When you do **not** want any type conversions to happen, there are two additional operators: `=== and !==`.

### Short-circuiting of logical operators
> The logical operators `&&` and `||` handle values of different types in a peculiar way. They will convert the value on their left side to Boolean type in order to decide what to do, but depending on the operator and the result of that conversion, they will return either the original left-hand value or the right-hand value.
```
console.log(null || "user")
// -> user
console.log("Agnes" || "user")
// -> Agnes
console.log(undefined || null)
// -> null
console.log(null || undefined)
// -> undefined
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
// -> null
console.log("Agnes" && "user")
// -> user
console.log(0/0 && null)
// -> NaN
console.log(undefined && 0/0)
// -> undefined
```
> The `&&` operator works similarly but the other way around. When the value to its left is something that converts to false, it returns that value, and otherwise it returns the value on its right.

> Another important property of these two operators is that the part to their right is evaluated only when necessary. In the case of true || X, no matter what X is—even if it’s a piece of program that does something terrible—the result will be true, and X is never evaluated. The same goes for false && X, which is false and will ignore X. This is called **short-circuit evaluation**.

> The conditional operator (ex. `true ? 1 : 2`) works in a similar way. Of the second and third values, only the one that is selected is evaluated.

### Summary
> We looked at four types of JavaScript values in this chapter: numbers, strings, Booleans, and undefined values.

> Such values are created by typing in their name (true, null) or value (13, "abc"). You can combine and transform values with operators. We saw binary operators for arithmetic (+, -, *, /, and %), string concatenation (+), comparison (==, !=, ===, !==, <, >, <=, >=), and logic (&&, ||), as well as several unary operators (- to negate a number, ! to negate logically, and typeof to find a value’s type) and a ternary operator (?:) to pick one of two values based on a third value.

## 2.Program Structure
### Expressions and statements
* A fragment of code that produces a value is called an `expression`. 
* If an expression corresponds to a sentence fragment, a JavaScript `statement` corresponds to a full sentence.
* It could display something on the screen—that counts as changing the world—or it could change the internal state of the machine in a way that will affect the statements that come after it. These changes are called `side effects`.
### Bindings
To catch and hold values, JavaScript provides a thing called a `binding`, or `variable`: `let caught = 5 * 5;`
* `let`: keyword
* `caught`: name
* `5 * 5`: value

After a binding has been defined, its name can be used as an expression.
> When a binding points at a value, that does not mean it is tied to that value forever. The `=` operator can be used at any time on existing bindings to disconnect them from their current value and have them point to a new one.
> 
> The words `var` and `const` can also be used to create bindings, in a way similar to `let`.
```
var name = "Ayda";
const greeting = "Hello ";
console.log(greeting + name);
// -> Hello Ayda
```
* `var` (short for “variable”), is the way bindings were declared in pre-2015 JavaScript. 
* `var` differs from `let` in the next chapter (3). 
* The word `const` stands for constant. It defines a constant binding, which points at the same value for as long as it lives.
### Binding names
* name must not start with a digit
* may include dollar signs ($) or underscores (_) but no other punctuation or special characters
* keywords (such as `let`) may not be used as binding names
* also a number of words that are “reserved for use” in future versions of JavaScript
```
break case catch class const continue debugger default
delete do else enum export extends false finally for
function if implements import interface in instanceof let
new package private protected public return static super
switch this throw true try typeof var void while with yield
```
### The environment
The collection of bindings and their values that exist at a given time is called the `environment`.

> When a program starts up, this environment is `not empty`. It always contains bindings that are part of the language standard, and most of the time, it also has bindings that provide ways to interact with the surrounding system.(For example, in a browser, there are functions to interact with the currently loaded website and to read mouse and keyboard input.)
### Functions
A function is a piece of program wrapped in a value. Such values can be applied in order to run the wrapped program.

Executing a function is called `invoking`, `calling`, or `applying` it: `fct()`
### The console.log function
> Most JavaScript systems (including all modern web browsers and Node.js) provide a console.log function that writes out its arguments to some text output device. 
>
> In browsers, the output lands in the JavaScript console. This part of the browser interface is hidden by default, but most browsers open it when you press F12 or, on a Mac, command-option-I.
>
> Though binding names cannot contain period characters, console.log does have one. This is because console.log isn’t a simple binding. It is actually an expression that retrieves the log property from the value held by the console binding. We’ll find out exactly what this means in Chapter 4.
### Return values
> Showing a dialog box or writing text to the screen is a side effect. A lot of functions are useful because of the side effects they produce. 

When a function produces a value, it is said to `return` that value.
### Control flow
program contains more than one statement -> executed from top to bottom
### Conditional execution
* `if` keyword in JavaScript
* The statement after the if is wrapped in braces ({ and })
* The braces can be used to group any number of statements into a single statement, called a `block`. 
* Can use the `else` keyword, together with if, to create two separate, alternative execution paths.
```
let theNumber = Number(prompt("Pick a number"));
if (!Number.isNaN(theNumber)) {
  console.log("Your number is the square root of " +
              theNumber * theNumber);
} else {
  console.log("Hey. Why didn't you give me a number?");
}
```
```
let num = Number(prompt("Pick a number"));

if (num < 10) {
  console.log("Small");
} else if (num < 100) {
  console.log("Medium");
} else {
  console.log("Large");
}
```
### while and do loops
> What we need is a way to run a piece of code multiple times. This form of control flow is called a `loop`.

> A statement starting with the keyword `while` creates a loop. The word `while` is followed by an expression in **parentheses** and then a statement, much like `if`.
```
let number = 0;
while (number <= 12) {
  console.log(number);
  number = number + 2;
}
// -> 0
// -> 2
//   … etcetera
```
> A `do` loop is a control structure similar to a while loop. It differs only on one point: a do loop always executes its body at least once, and it starts testing whether it should stop only after that first execution. 
```
let yourName;
do {
  yourName = prompt("Who are you?");
} while (!yourName);
console.log(yourName);
```
### Indenting Code
> The role of this indentation inside blocks is to make the structure of the code stand out.
>
> The important thing is that each new block adds the same amount of space.
### for loops
```
for (let number = 0; number <= 12; number = number + 2) {
  console.log(number);
}
// -> 0
// -> 2
//   … etcetera
```
> The parentheses after a for keyword must contain `two semicolons`. The part before the first semicolon `initializes` the loop, usually by defining a binding. The second part is the expression that `checks` whether the loop must continue. The final part `updates` the state of the loop after every iteration. In most cases, this is shorter and clearer than a while construct.
### Breaking Out of a Loop
> There is a special statement called `break` that has the effect of immediately jumping out of the enclosing loop.
```
for (let current = 20; ; current = current + 1) {
  if (current % 7 == 0) {
    console.log(current);
    break;
  }
}
// -> 21
```
no `break` -> `infinite loop`
> The `continue` keyword is similar to break, in that it influences the progress of a loop. When continue is encountered in a loop body, control jumps out of the body and continues with the loop’s next iteration.
### Updating bindings succinctly
> For `counter += 1` and `counter -= 1`, there are even shorter equivalents: `counter++` and `counter--`.
### Dispatching on a value with switch
> There is a construct called `switch` that is intended to express such a “dispatch” in a more direct way. Unfortunately, the syntax JavaScript uses for this (which it inherited from the C/Java line of programming languages) is somewhat awkward—a chain of if statements may look better.
```
switch (prompt("What is the weather like?")) {
  case "rainy":
    console.log("Remember to bring an umbrella.");
    break;
  case "sunny":
    console.log("Dress lightly.");
  case "cloudy":
    console.log("Go outside.");
    break;
  default:
    console.log("Unknown weather type!");
    break;
}
```
But be careful—it is easy to forget such a break
### Capitalization
> he standard JavaScript functions, and most JavaScript programmers, follow the bottom style—they capitalize every word except the first. 

>In a few cases, such as the `Number` function, the first letter of a binding is also capitalized. This was done to mark this function as a `constructor`.
### Comments
> To write a single-line comment, you can use two slash characters (`//`) and then the comment text after it.
>
> A section of text between `/*` and `*/` will be ignored in its entirety, regardless of whether it contains line breaks. 
### Summary
> You now know that a program is built out of `statements`, which themselves sometimes contain more statements. Statements tend to contain `expressions`, which themselves can be built out of smaller expressions.

> Putting statements after one another gives you a program that is executed from `top to bottom`. You can introduce disturbances in the flow of control by using conditional (`if, else, and switch`) and looping (`while, do, and for`) statements.

> `Bindings` can be used to file pieces of data under a name, and they are useful for tracking state in your program. The `environment` is the set of bindings that are defined. JavaScript systems always put a number of useful standard bindings into your environment.

> `Functions` are special values that encapsulate a piece of program. You can invoke them by writing `functionName(argument1, argument2)`. Such a function call is an expression and may produce a value.
### [3 exercices](https://eloquentjavascript.net/code/#2)
## 3.Functions
### Defining a function
```
const square = function(x) {
  return x * x;
};

console.log(square(12));
// -> 144
```
> A function is created with an expression that starts with the keyword `function`. Functions have a set of `parameters` (in this case, only x) and a `body`, which contains the statements that are to be executed when the function is called. The function body of a function created this way must always be `wrapped in braces`, even when it consists of only a single statement.
>
> A return statement determines the value the function `returns`. When control comes across such a statement, it immediately **jumps out** of the current function and gives the returned value to the code that called the function. 
> 
> A `return` keyword **without an expression after it** will cause the function to return `undefined`. 
> 
> Functions that **don’t have a return** statement at all similarly return `undefined`.
> 
> `Parameters` to a function behave like regular bindings, but their initial values are given by the `caller` of the function, not the code in the function itself.
### Bindings and scopes
> Each binding has a `scope`, which is the part of the program in which the binding is visible. For bindings defined outside of any function or block are called `global`.
```
let x = 10;
if (true) {
  let y = 20;
  var z = 30;
  console.log(x + y + z);
  // -> 60
}
// y is not visible here
console.log(x + z);
// -> 40
```
> Bindings declared with `let` and `const` are in fact local to the block that they are declared in, so if you create one of those inside of a loop, the code before and after the loop cannot “see” it. In pre-2015 JavaScript, only functions created new scopes, so old-style bindings, created with the var keyword, are visible throughout the whole function that they appear in—or throughout the global scope, if they are not in a function.
```
const halve = function(n) {
  return n / 2;
};

let n = 10;
console.log(halve(100));
// -> 50
console.log(n);
// -> 10
```
> The exception is when multiple bindings have the same name—in that case, code can see only the innermost one.
### Nested scope
> JavaScript distinguishes not just `global` and `local` bindings. Blocks and functions can be created inside other blocks and functions, producing multiple degrees of locality.
```
const hummus = function(factor) {
  const ingredient = function(amount, unit, name) {
    let ingredientAmount = amount * factor;
    if (ingredientAmount > 1) {
      unit += "s";
    }
    console.log(`${ingredientAmount} ${unit} ${name}`);
  };
  ingredient(1, "can", "chickpeas");
  ingredient(0.25, "cup", "tahini");
  ingredient(0.25, "cup", "lemon juice");
  ingredient(1, "clove", "garlic");
  ingredient(2, "tablespoon", "olive oil");
  ingredient(0.5, "teaspoon", "cumin");
};
```
> The code inside the ingredient function can see the factor binding from the outer function. But its local bindings, such as unit or ingredientAmount, are not visible in the outer function.

> Each local scope can also see all the local scopes that contain it, and all scopes can see the global scope. This approach to binding visibility is called `lexical scoping`.
### Functions as values
> A function binding usually simply acts as a name for a specific piece of the program. Such a binding is defined once and never changed. This makes it easy to confuse the function and its name.

> But the two are different. A function value can do all the things that other values can do—you can use it in arbitrary expressions, not just call it. It is possible to store a function value in a new binding, pass it as an argument to a function, and so on. Similarly, a binding that holds a function is still just a regular binding and can, if not constant, be assigned a new value, like so:
```
let launchMissiles = function() {
  missileSystem.launch("now");
};
if (safeMode) {
  launchMissiles = function() {/* do nothing */};
}
```
### Declaration notation
> There is a slightly shorter way to create a function binding. When the function keyword is used at the start of a statement, it works differently.
```
function square(x) {
  return x * x;
}
```
> This is a function `declaration`. The statement defines the binding square and points it at the given function. It is slightly easier to write and doesn’t require a semicolon after the function.

* Function declarations are not part of the regular top-to-bottom flow of control. They are conceptually moved to the top of their scope and can be used by all the code in that scope.
### Arrow functions
Instead of the function keyword, it uses an arrow (`=>`) 
```
const power = (base, exponent) => {
  let result = 1;
  for (let count = 0; count < exponent; count++) {
    result *= base;
  }
  return result;
};
```
> “this input (the parameters) produces(`=>`) this result (the body)”
```
const square1 = (x) => { return x * x; };
const square2 = x => x * x; 
// one parameter name & single expression
const horn = () => {
  console.log("Toot");
};
// no parameter
```
> Arrow functions were added in 2015, mostly to make it possible to write small function expressions in a less verbose way. We’ll be using them a lot in Chapter 5.
### The call stack
> Because a function has to jump back to the place that called it when it returns, the computer must remember the context from which the call happened. The place where the computer stores this context is the `call stack`. 

> Every time a function is called, the current context is stored on top of this stack. When a function returns, it removes the top context from the stack and uses that context to continue execution. (`LIFO`)
### Optional Arguments
```
function square(x) { return x * x; }
console.log(square(4, true, "hedgehog"));
// → 16
```
> JavaScript is extremely broad-minded about the number of arguments you pass to a function. If you pass too many, the extra ones are `ignored`. If you pass too few, the missing parameters get assigned the value `undefined`.

> The upside is that this behavior can be used to allow a function to be called with different numbers of arguments.
```
function minus(a, b) {
  if (b === undefined) return -a;
  else return a - b;
}

console.log(minus(10));
// → -10
console.log(minus(10, 5));
// → 5
```
> If you write an = operator after a parameter, followed by an expression, the value of that expression will replace the argument when it is not given.
```
function power(base, exponent = 2) {
```
### Closure
```
tion wrapValue(n) {
  let local = n;
  return () => local;
}

let wrap1 = wrapValue(1);
let wrap2 = wrapValue(2);
console.log(wrap1());
// → 1
console.log(wrap2());
// → 2
```
> This feature—being able to reference a specific instance of a local binding in an enclosing scope—is called `closure`. A function that references bindings from local scopes around it is called a `closure`.

> A good mental model is to think of function values as containing both the code in their body and the environment in which they are created.
### Recursion
> A function that calls itself is called `recursive`. In typical JavaScript implementations, it’s about three times `slower` than the looping version.
### Growing functions
two more or less natural ways for functions:
* writing similar code multiple times.
* need some functionality that you haven’t written yet and that sounds like it deserves its own function. 
### Functions and side effects
> Functions can be roughly divided into those that are called for their side effects and those that are called for their return value. 

> A `pure` function is a specific kind of value-producing function that not only has no side effects but also doesn’t rely on side effects from other code—for example, it doesn’t read global bindings whose value might change. 
### Summary
```
// Define f to hold a function value
const f = function(a) {
  console.log(a + 2);
};

// Declare g to be a function
function g(a, b) {
  return a * b * 3.5;
}

// A less verbose function value
let h = a => a % 3;
```
> Parameters and bindings declared in a given scope are local and not visible from the outside. Bindings declared with var behave differently—they end up in the nearest function scope or the global scope.

> Separating the tasks your program performs into different functions is helpful.
### [3 exercices](https://eloquentjavascript.net/code/#3)
## 4.Data Structures: Objects and Arrays
> Objects allow us to group values—including other objects—to build more complex structures.

> JavaScript provides a data type specifically for storing sequences of values. It is called an array and is written as a list of values between square brackets, separated by commas.
```
let listOfNumbers = [2, 3, 5, 7, 11];
```
* The first index of an array is zero
* aList.length / aList["length"]: access the length property of the value in aList
### Properties
> Almost all JavaScript values have properties. The exceptions are `null` and `undefined`.
```
null.length;
// → TypeError: null has no properties
```
> The two main ways to access properties in JavaScript are with a dot and with square brackets. Both `value.x` and `value[x]` access a property on value—but not necessarily the same property. The difference is in how x is interpreted. 

* the word after the dot is the literal name of the property. `array.length`
* the expression between the brackets is evaluated to get the property name. `array.length`
### Methods
```
let doh = "Doh";
console.log(typeof doh.toUpperCase);
// → function
console.log(doh.toUpperCase());
// → DOH
```
```
let sequence = [1, 2, 3];
sequence.push(4);
console.log(sequence);
// → [1, 2, 3, 4]
console.log(sequence.pop());
// → 4
console.log(sequence);
// → [1, 2, 3]
```
> A `stack`, in programming, is a data structure that allows you to push values into it and pop them out again in the opposite order so that the thing that was added last is removed first. (`LIFO`) These are common in programming—you might remember the function call `stack` from the previous chapter, which is an instance of the same idea.
### Objects
> Values of the type object are arbitrary collections of properties. One way to create an object is by using braces as an expression.
```
let day1 = {
  squirrel: false,
  events: ["work", "touched tree", "pizza", "running"]
};
console.log(day1.squirrel);
// → false
console.log(day1.wolf);
// → undefined
day1.wolf = false;
console.log(day1.wolf);
// → false
```
> Reading a property that doesn’t exist will give you the value `undefined`.

> When an object is written over multiple lines, indenting it like in the example helps with readability.

> This means that braces have two meanings in JavaScript. At the start of a statement, they start a block of statements. In any other position, they describe an object. Fortunately, it is rarely useful to start a statement with an object in braces, so the ambiguity between these two is not much of a problem.

* assign a value to a property expression: `=`
* unary operator, when applied to an object property, remove the named property from the object: `delete`
* applied to a string and an object, whether that object has a property with that name: `in`
* what properties an object has: `Object.keys`
```
console.log(Object.keys({x: 0, y: 0, z: 2}));
// → ["x", "y", "z"]
```
> The difference between setting a property to undefined and actually deleting it is that, in the first case, the object still has the property (it just doesn’t have a very interesting value), whereas in the second case the property is no longer present and in will return false.
```
let objectA = {a: 1, b: 2};
Object.assign(objectA, {b: 3, c: 4});
console.log(objectA);
// → {a: 1, b: 3, c: 4}
```
> There’s an `Object.assign` function that copies all properties from one object into another.
### Mutability
```
let object1 = {value: 10};
let object2 = object1;
let object3 = {value: 10};
console.log(object1 == object2);
// → true
console.log(object1 == object3);
// → false
object1.value = 15;
console.log(object2.value);
// → 15
console.log(object3.value);
// → 10
```
> The object1 and object2 bindings grasp the same object, which is why changing object1 also changes the value of object2. They are said to have the same `identity`. The binding object3 points to a different object, which initially contains the same properties as object1 but lives a separate life.

> Bindings can also be changeable or constant, but this is separate from the way their values behave. you can use a let binding to keep track of a changing number by changing the value the binding points at. Similarly, though a const binding to an object can itself not be changed and will continue to point at the same object, the contents of that object might change.
```
const score = {visitors: 0, home: 0};
// This is okay
score.visitors = 1;
// This isn't allowed: TypeError: invalid assignment to const 'score'
score = {visitors: 1, home: 1};
```
> When you compare objects with JavaScript’s == operator, it compares by `identity`: it will produce true only if both objects are precisely the same value. (Comparing different objects will return false)

> Arrays have an `includes` method that checks whether a given value exists in the array. The function uses that to determine whether the event name it is interested in is part of the event list for a given day.
### Array loops
> There is a simpler way to write such loops in modern JavaScript.
```
for (let entry of JOURNAL) {
  console.log(`${entry.events.length} events.`);
}
```
> When a `for` loop looks like this, with the word `of` after a variable definition, it will loop over the elements of the value given after of. This works not only for arrays but also for strings and some other data structures. We’ll discuss how it works in Chapter 6.
### Further arrayology
> We saw `push` and `pop`, which add and remove elements at the end of an array, earlier in this chapter. The corresponding methods for adding and removing things at the start of an array are called `unshift` and `shift`. 

> To search for a specific value, arrays provide an `indexOf` method. Both indexOf and lastIndexOf take an optional second argument that indicates where to start searching.
```
console.log([1, 2, 3, 2, 1].indexOf(2));
// → 1
console.log([1, 2, 3, 2, 1].lastIndexOf(2));
// → 3
```
> Another fundamental array method is `slice`, which takes start and end indices and returns an array that has only the elements between them. The start index is inclusive, the end index exclusive.
```
console.log([0, 1, 2, 3, 4].slice(2, 4));
// → [2, 3]
console.log([0, 1, 2, 3, 4].slice(2));
// → [2, 3, 4]
```
> The `concat` method can be used to glue arrays together to create a new array, similar to what the + operator does for strings.
### Strings and their properties
* `length` and `toUpperCase` from string values
* immutable 
* do have built-in properties: `slice` and `indexOf`
* `trim` method removes whitespace
* `padStart`
* `split` and `join`
* a string can be repeated with the `repeat` method, which creates a new string
```
console.log("one two three".indexOf("ee"));
// → 11
console.log("  okay \n ".trim());
// → okay
console.log(String(6).padStart(3, "0"));
// → 006
let sentence = "Secretarybirds specialize in stomping";
let words = sentence.split(" ");
console.log(words);
// → ["Secretarybirds", "specialize", "in", "stomping"]
console.log(words.join(". "));
// → Secretarybirds. specialize. in. stomping
console.log("LA".repeat(3));
// → LALALA
```
### Rest parameters
> It can be useful for a function to accept any number of arguments. To write such a function, you put three dots `...` before the function’s last parameter. 
```
function max(...numbers) {
  let result = -Infinity;
  for (let number of numbers) {
    if (number > result) result = number;
  }
  return result;
}
console.log(max(4, 1, 9, -2));
// → 9
```
> can use a similar three-dot notation to call a function with an array of arguments.
```
let numbers = [5, 1, 7];
console.log(max(...numbers));
// → 7
let words = ["never", "fully"];
console.log(["will", ...words, "understand"]);
// → ["will", "never", "fully", "understand"]
```
### The Math object
> The Math object is used as a container to group a bunch of related functionality. There is only one Math object, and it is almost never useful as a value. Rather, it provides a `namespace` so that all these functions and values do not have to be global bindings.

> Having too many global bindings “pollutes” the namespace. The more names have been taken, the more likely you are to accidentally overwrite the value of some existing binding.

> Many languages will stop you, or at least warn you, when you are defining a binding with a name that is already taken. JavaScript does this for bindings you declared with `let` or `const` but—perversely—not for standard bindings nor for bindings declared with `var` or `function`.

* Math.random() -> [0,1)
* Math.floor (rounds down to the nearest whole number)
```
console.log(Math.floor(Math.random() * 10));
// → 2
```
### JSON
> What we can do is serialize the data. That means it is converted into a flat description. `JavaScript Object Notation`

> JSON looks similar to JavaScript’s way of writing arrays and objects, with a few restrictions. All property names have to be `surrounded by double quotes`, and only simple data expressions are allowed—no function calls, bindings, or anything that involves actual computation. Comments are not allowed in JSON.

> JavaScript gives us the functions `JSON.stringify` and `JSON.parse` to convert data to and from this format. The first takes a JavaScript value and returns a JSON-encoded string. The second takes such a string and converts it to the value it encodes.
### Summary
> `Objects` and `arrays` (which are a specific kind of object) provide ways to group several values into a single value.

> Most values in JavaScript have properties, the exceptions being `null` and `undefined`. Properties are accessed using `value.prop` or `value["prop"]`. Objects tend to use names for their properties and store more or less a fixed set of them. `Arrays`, on the other hand, usually contain varying amounts of conceptually identical values and use numbers (starting from 0) as the names of their properties.

> There are some named properties in arrays, such as `length` and a number of methods. Methods are functions that live in properties and (usually) act on the value they are a property of.

> You can iterate over arrays using a special kind of for loop—for (`let element of array`).
### [4 exercices](https://eloquentjavascript.net/code/#4)
## 5.Higher-Order Functions
> expressing simpler concepts than the program as a whole is easier to get right.
### Abstraction
> `Abstractions` hide details and give us the ability to talk about problems at a higher (or more abstract) level.
### Abstracting repetition
> Often, it is easier to create a function value on the spot instead.
```
let labels = [];
repeat(5, i => {
  labels.push(`Unit ${i + 1}`);
});
console.log(labels);
// → ["Unit 1", "Unit 2", "Unit 3", "Unit 4", "Unit 5"]
```
### Higher-order functions
> Functions that operate on other functions, either by taking them as arguments or by returning them, are called `higher-order functions`.

> Higher-order functions allow us to abstract over actions, not just values. They come in several forms. For example, we can have functions that create new functions.
```
function greaterThan(n) {
  return m => m > n;
}
let greaterThan10 = greaterThan(10);
console.log(greaterThan10(11));
// → true
```
> There is a built-in array method, `forEach`, that provides something like a for/of loop as a higher-order function.
```
["A", "B"].forEach(l => console.log(l));
// → A
// → B
```
### Script data set


