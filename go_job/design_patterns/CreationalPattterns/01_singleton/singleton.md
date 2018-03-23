# Singleton Design Pattern

- Exploring the example of a unique counter
- Implementing singleton patterns


# What is Singleton Design Pattern?

- Provides a single instance of an object
- Guarantees that there are no duplicates
- Easy to remember

# Description

- When you want to use the same connection to a database to make every query.
- When you open SSH connection to a server to do a few tasks
- If you need to limit the access to some variable of space
- If you need to limit the number of calls to some places

# Objectives

- We need a single, shared value, of some particular type
- We need to restrict object creation of some type to a single unit along the entire program

# Example - A Unique Counter

## Requirements and Acceptance Criteria

- When no counter has been created before, a new one is created with the value 0
- If a counter has already been created, return this instance that holds actual count
- If we call the method AddOne, the count must be incremented by 1

## Writing Unit Tests
- In Go, there's nothing like static members
- You have package scope to deliver a similar result
- To set up our project, create a new folder with in $GOPATH/src directory
- General rule is to create a subfolder with the VCS provider, the username, and the name of the project

