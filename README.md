# Golang DDD (Domain Driven Design) Architecture Implementation

Domain - Niche of Business Logic
Core Topic Domain - Tavern, Sub Topic Doman - Customers.

- Entity: Unique Identifer, Mutable ()
- Value Object: No Identifer, Immutable ()
- Agregates: Unique Identifer, by Root Entity, Multiple Entities/Value objects combined (Combination of entities and value objects, when you combine them you get aggregate)
    - e.g. Instance of Customer (Entity): Person - Root Entity, Dante, Products (Entity) and Transactions (Entity).
    - It is not up to the aggregate to decide how the data is supposed to be formatted (JSON).

- Domain
- Entity
- Value Object
- Aggregates

---

Actuall Business Logic.

Factory Pattern - is design pattern that is used to encapsulate complex logic inside function for creating The Wanted instances without the caller knowing anything about the actual implementation details (With Interfaces).

DDD suggests factories for creating complex Aggregates repositories or Service.
