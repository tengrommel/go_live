# factory method
- Implementing a payments methods Factory
- Creating two methods of paying-cash and credit card
- Creating an interface with the Pay method

# What
- Delegates the creation of different type of payments
- The second-best known and used design pattern in the industry
- Abstracts the user from the knowledge of the struct he needs to achieve 
- Provides an interface that fits the user needs
- Eases the process of downgrading or upgrading of the implementation

# Acceptance Criteria
- To have a common method for every payment method called pay
- To be able to delegate the creation of payments methods to the factory
- To be able to add more payment methods to the library by just adding it to the factory method 

# Summarizing Factory Method
- Grouping families of objects
- Upgrading an implementation of used structs
- Tests must be written with care if you don't want to tie to certain implementations