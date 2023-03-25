Here is an example of how to structure package names for a banking system using the Hexagonal Architecture pattern:

domain: account, transaction, customer
application: accountservice, transactionservice, customerservice
infrastructure: accountrepository, transactionrepository, customerrepository
interface: accountport, transactionapi, customerapi


doubts
1. naming handler in adapter is according to ports adap , but actual go given name is controller


Hexagonal Architecture Layers

Creating dependency rules can decouple the layers. The hexagonal architecture is also called the Ports and Adapters. The outer layers may only depend on the inner layers, and the inner layers should not rely on the outer ones. Each layer is defined as follows.

Infrastructure Layer ( infra in our application )
The Infrastructure Layer contains code interfaces with application infrastructure – controllers, UI, persistence, and gateways to external systems. A web framework or persistence library will provide many of the objects in this layer. Concretions of domain repositories are placed in this layer, while the virtual interfaces are defined in the domain layer.

Application Layer (components in our application)
The Application Layer provides an API for all functionality provided by the application. It accepts commands from the client (web, API, or CLI) and translates them into values understood by the domain layer. For example, a RegisterUser service would accept a Data Transfer Object containing a new user's credentials and delegate responsibility for creating a user to the domain layer.

Domain Layer ( domain )
The Domain Layer contains any core domain logic. It deals entirely with domain concepts and lacks knowledge of the outer layers.

