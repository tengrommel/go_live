# go fundamentals

## Writing Functions in Go

- Problem
> How do you manage Go code in functions?

- Solution

1 .The keyword func is used to declare functions. 

2 A function is declared with a name, a list of parameters, an optional list types, 
and a body to write the logic for the function.

- How It Works

1. A function in Go is a reusable piece of code that organizes a sequence of code statements as a unit,
which can be called from within the package, and also from other packages if the functions are exported to other packages.

2. Because functions are reusable piece of code, you can call this form multiple times.

3. When you write shared library packages, the functions with names that start with an uppercase letter will be exported to 
other packages.

4. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function
within the same package.

- Declaring Functions

    
    func name(list of parameters) (list of return types){
        function body
    }
    
- Naming Return Values
> When you write functions, you can name the return values by defining variables at the top of the function.

- Returning Multiple Values
> Go is a language that provides lot of pragmatism in its language design.

In Go, you can return multiple values from a function, which is a helpful feature in many practical scenarios.

## Variadic Functions
> A variadic function is a function that accepts a variable number of arguments.

## Function Values, Anonymous Functions, and Closures
>Go's pragmatism gives developers productivity like a dynamically typed language although Go is a statically typed language.

