# Golang DDD (Domain Driven Design) Architecture Implementation

This is a **Tavern** Business - Domain Driven Design Architecture Implementation in Golang.

---

Domain - Niche of Business Logic
Core Topic Domain - Tavern, Sub-Topic Doman - Customers.

**Layers:**

- Core Domain
- Entities
- Value Objects
- Aggregates
    - Factory Pattern
- Repositories
- Services
    - Configuration Pattern

## Explanation

**Entity:**

- Entity: Unique Identifer, Mutable.
Entity - .

---

**Value Object:**

- Value Object: No Identifer, Immutable.
Value Object - .

---

**Agregates:**

- Agregates: Unique Identifer, by Root Entity, Multiple Entities/Value objects combined.
Agregates - combination of entities and value objects, when you combine them you get aggregate

e.g. Instance of Customer (Entity): Person - Root Entity, Dante, Products (Entity) and Transactions (Entity).
It is not up to the aggregate to decide how the data is supposed to be formatted (JSON).

---

Actuall Business Logic.

**Factory Pattern:**

Factory Pattern - is design pattern that is used to encapsulate complex logic inside function for creating The Wanted instances without the caller knowing anything about the actual implementation details (with Interfaces).

DDD suggests factories for creating complex Aggregates, Repositories or Service.

---

- Aggregates stored by Repositories.

Aggregates is combination of Entities and Value Objects (Factory Pattern is used for hide complexity of aggregates logic and also creating instances without the caller knowing anything about actual implementation details),
but when we store them or manage them we are using a Repository.

**Repository:**

Repository - used to store and manage Aggregates.
Repository Pattern - relies on hiding the implementation details behind a interface,
and this is allows us to build very modular and changeable software.
We can have In-Memory Repository which stores customers in memory whenever we do Unit tests,
but then also we can have a MySQL Repository, and whenever the managers come and say hey we're changing from MySQL to MongoDB for instance we can build a new Repository for MongoDB and fulfill the same Interface as MySQL Repository and then just swap it, and everything should work as expected.

---

- How do we need to fetch our Customers from memory repository/?
- we need to add a way to retrieve data from the aggregate, like customer.person.ID

The data is not accessible from outside of the aggregate,
nothing outside of the aggregate can modify data,
this is done by exposing functions that allows others to do it.
If we want to modify name we expose a function which allows you to do that.
You can not go directly and modify it.

---

In this case of Domain/Product/Memory/memory.go Repository,
we will never return our error
but it's up to the interface to determine if it's possible
it's not up to the implementation it's up to the interface
so repository decides if we need to return an error or not

---

**Services:**

Services - will tie together all loosely coupled repositories into business logic that fulfills the needs of domain.

Order service - is responsible for shaming together the repositories that performs an order, like getting customer, product repository for instance and then making order, billing service.

Whenever you have Services the factories tends to get larger and more complex
because you accept multiple repositories as input for instance, for this we will use Service Configuration Generator Pattern.

---

Service can hold multiple repositories, but a service is also allowed to hold other services.
If our service is needs other service, we could simply hold a sub-service inside a service.

---

**Configuration Pattern:**

Configuration Pattern - .

Order Configuration - applies a customer repository to the service.

---

**Summary:**

Entities and Value objects are instances,
Entities are mutable (changeable) and Value objects are Not.

Aggregates hold multiple Entities or Value objects,
but they are related to one root Entity.

Repository manages Aggregates.

Services combines and ties together Repositories and other Services.

## TO DO

- [ ] Refactor (names, dirs and etc.)
- [ ] fix error handling (like messages in repo interface and also names of variables)
    - e.g. `product.ErrorFailedToAddProduct`-`"failed to add the product"` to `product.ErrorProductAlreadyExists`-`"product already exists"`
- [ ] write unit tests for product memory repository, product aggregate, order service
    - Customer Memory Repo:
        - TestMemory_GetCustomer, aggregate.NewCustomer
        - TestMemory_AddCustomer
    - Product Memory Repo:
        - TestMemoryProductRepository_Get
        - TestMemoryProductRepository_Add
        - Update
        - Delete
    - Customer Aggregate:
        - TestCustomer_NewCustomer
    - Product Aggregate:
        - TestProduct_NewProduct
    - Order Service:
        - TestOrder_NewOrderService, with customer and products aggregate
- [ ] correct naming of unit tests and other stuffs (Naming Convention)
- [ ] unit testing: make package memory to memory_test and use memory.Something
    - change `t.Fatal(err)` to `t.Error(err)`
- [ ] setup it in your vscode and test it - [Setting up Delve and Air to debug Golang with VS Code](https://dev.to/nerdherd/setting-up-golang-on-vs-code-with-debugging-1kbe)

## Links and Sources

**About DDD:**

- Eric Evans is inventor of DDD Architecture
- DDD Architecture in microservices

---

**Used Patterns:**

- factory pattern
- configuration pattern

Questions:

- difference between configuration pattern and dependency injection pattern?
- is dependency injection pattern used in DDD?

---

**Golang DDD Implementation:**

ProgrammingPercy:

- https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/
- https://programmingpercy.tech/blog/how-to-structure-ddd-in-go/

---

**Golang Specific Topic:**

**My Research's:**
