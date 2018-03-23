# Builder

Taking the example of vehicle manufacturing to implement builder design pattern

# What

- Helps to construct complex objects without directly instantiating their struct
- Avoids writing the logic to create all the objects in the package

# Description

- Instance creation is simple or complex
- An object can be composed of many objects
- You can use the same technique to create many types of objects
- Builder pattern contributes to this approach

# Objectives of Builder Design Pattern

- Abstract complex creations
- Create an object step by step by filling its fields and creating embedded objects
- Reuse the object creation algorithm between many objects

# Example - Vehicle Manufacturing

- Create relationship between a director, a few builders, and the product they build
- Every kind of vehicle, choose vehicle type, assemble the structure, place the wheels, and place the seats
- Director is represented by the ManufacturingDirector type

# Requirements and Acceptance Criteria

- You must have a manufacturing type
- VehicleProduct with four wheels, five seats, and a structure defined as cat must be returned
- VehicleProduct with two wheels, two seats, and a structure defined as motorbike must be returned
- VehicleProduct built by any BuildProcess builder must be open to modifications

