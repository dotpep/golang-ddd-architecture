# Golang DDD (Domain Driven Design) Architecture Implementation

Domain - Niche of Business Logic
Core Topic Domain - Tavern, Sub Topic Doman - Customers.

- Core Domain
- Entities
- Value Objects
- Aggregates
- Factory Pattern
- Repositories
- Services
- Configuration Pattern

---

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

DDD suggests factories for creating complex Aggregates repositories or Service.

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

**Services:**

Services - will tie together all loosely coupled repositories into business logic that fulfills the needs of domain.

Order service - is responsible for shaming together the repositories that performs an order, like getting customer, product repository for instance and then making order, billing service.

Whenever you have Services the factories tends to get larger and more complex
because you accept multiple repositories as input for instance, for this we will use Service Configuration Generator Pattern.

---

**Configuration Pattern:**

Configuration Pattern - .

Order Configuration - applies a customer repository to the service.
