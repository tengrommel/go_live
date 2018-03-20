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

## Working with Dynamic Arrays Using Slices

Declaring nil Slice
>Declaring a slice is similar to declaring an array, but when you declare slices, you don't need to specify the size because
it is a dynamic array.

**Initializing Slices Using make Function**
> A slice must be initialized before assigning values.

In the preceding declaration, the slice x was declared, but it was not initialized, 
so if you try to assign values to it this will cause a runtime error.

>Go's built-in make function is used to initialize slices.

## Enlarging Slices with copy and append Functions

1. Because slices are dynamic arrays, you can enlarge them whenever you want. 
2. When you want to increase the capacity of a slice, one approach is to create a new,
 bigger slice and copy the elements of the original slice into newly created one.
3. Go’s built-in copy function is used to copy data from one slice to another. 


## Iterating Over Slices
> An idiomatic approach for iterating over the elements of a slice is to use a range construct.

## Persisting Key/Value Pairs Using Map
> Go's map type allows you to store a collection of key/value pairs into a structure similar to hash table.

Go's map type is a data structure that provides an implementation of a hash table.

## Working with Maps
>Maps provide fast lookups on the data elements in the data structure.