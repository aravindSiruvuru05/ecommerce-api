**Running the debugger inside docker container:**

1.  Install the Go extension for VSCode. You can do this by searching for "Go" in the extensions tab on the left side of the VSCode window and clicking "Install
2.  Create a launch.json file by clicking on the debug tab on the left side of the VSCode window and selecting  “create a launch.json file”. 
3.  Edit the launch.json file to include the below configuration (JSON below).
4.  Build debugger  `./dev-env run build debugger`
5.  Run application in DEBUG MODE  `./dev-env up debug`

```javascript
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Beego in Docker",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "port": 2345,
            "host": "localhost",
            "cwd": "${workspaceFolder}",
            "showLog": true
        }
    ]
}   
```

**Here is an example of how to structure package names for a banking system using the Hexagonal Architecture pattern:**

domain: account, transaction, customer

application: accountservice, transactionservice, customerservice

infrastructure: accountrepository, transactionrepository, customerrepository

interface: accountport, transactionapi, customerapi

   
 

doubts

1\. naming handler in adapter is according to ports adap , but actual go given name is controller

   
 

Hexagonal Architecture Layers

Creating dependency rules can decouple the layers. The hexagonal architecture is also called the Ports and Adapters. The outer layers may only depend on the inner layers, and the inner layers should not rely on the outer ones. Each layer is defined as follows.

Infrastructure Layer ( infra in our application )

The Infrastructure Layer contains code interfaces with application infrastructure – controllers, UI, persistence, and gateways to external systems. A web framework or persistence library will provide many of the objects in this layer. Concretions of domain repositories are placed in this layer, while the virtual interfaces are defined in the domain layer.

Application Layer (components in our application)

The Application Layer provides an API for all functionality provided by the application. It accepts commands from the client (web, API, or CLI) and translates them into values understood by the domain layer. For example, a RegisterUser service would accept a Data Transfer Object containing a new user's credentials and delegate responsibility for creating a user to the domain layer.

Domain Layer ( domain )

The Domain Layer contains any core domain logic. It deals entirely with domain concepts and lacks knowledge of the outer layers.