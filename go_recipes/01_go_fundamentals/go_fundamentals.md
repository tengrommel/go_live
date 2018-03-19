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