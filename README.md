# spectracld-backend
## Hexagonal architecture 

* This application includes the Ports and Adapters, its a design pattern that allows to design and implement maintainable and scalable software systems.
* In this architecture, the business logic is placed in the core of the system, surrounded by ports (interfaces) and adapters (implementations of those interfaces).
* This approach enables easy integration and testing of different parts of the systems.


* **Application Layer**: This layer contains the **main** function and the entry point for the application. It communicates with the application's users or other systems by calling methods on the use case interfaces.
* **Use case Layer**: This layer contains the application's business logic, which is encapsulated in the use case classes. Each use case is a seprate module that represents a single user action. The use case layer is completely independent of the other layers, including the persistence layer and the user interface layer
*  **Domain Layer**: This layer contains the business entities, value objects, and other domain-specific classes. It provides the underlyuing abstractions that the use case layer builds upon.
*  **Infrastructure Layer**: This layer contains the adapters for the ports defined in the domain layer. Adapters translate the domain-specific interfaces into the format reuired by the underlying infrastructure or third-party services.



