# Understanding Adapter Design Pattern in Go

The Adapter design pattern is used when we want to connect two incompatible interfaces or classes so that they can work together. In this blog, we will see how adapter design patterns can be useful in writing modular code, understand its use cases, and see how we can implement it in Golang. 

## What is the Adapter design pattern?

The adapter design pattern acts as an intermediate layer between two incompatible interfaces and makes them work together by converting the interface of one class into another.

In Golang we can do this using interfaces and structs. Let’s see this with the help of an example. 

I am using a MacBook Pro (13-inch, 2020, Four Thunderbolt 3 ports). It has 4 USB type C ports. The laptop looks very stylish with its compact and sleek design. But I regularly face issues when I want to insert a USB type A pen-drive or mirror the screen on a monitor using an HDMI cable. To overcome this issue I use a separate dongle(adapter) for USB type A and one for HDMI port. Let’s see this with the help of code.

```
type laptop interface {
    insertUSBTypeC()
}

type client struct{}

func (c *client) insertUSBCIntoLaptop(l laptop) {
    l.insertUSBTypeC()
}
```

As you can see, we have a laptop interface that represents the target interface that our client code expects. The client structure helps insert USB type C into the laptop.

```
type typeCPendrive struct{}

func (t *typeCPendrive) insertUSBTypeC() {
    fmt.Println("insert USB Type C")
}

type typeAPendrive struct{}

func (t *typeAPendrive) insertUSBTypeA() {
    fmt.Println("insert USB Type A")
}
```

Now we have two structs one that accepts USB type C and one that accepts USB type A. While it is easy to insert typeCPendrive, the same is not true for typeAPendrive since it is not compatible with our laptop interface. This is where we are going to apply the adapter.

```
type typeAAdapter struct {
    pendrive *typeAPendrive
}

func (t *typeAAdapter) insertUSBTypeC() {
    t.pendrive.insertUSBTypeA()
}
```

Above we have created a new adapter type called typeAAdapter that accepts the typeAPendrive and implements the interface method insertUSBTypeC(). An adapter that accepts a USB type A connector and then translates its signals to USB type C format. We can use this inside the main function like below.

```
func main() {
    client := &client{}
    typeC := &typeCPendrive{}
    client.insertUSBCIntoLaptop(typeC)

    typeA := &typeAPendrive{}
    adapter := &typeAAdapter{
    pendrive: typeA,
    }
    client.insertUSBCIntoLaptop(adapter)
}
```

In this way, by using the adapter design pattern we can use a new type, USB type A with our existing system.

## Use Cases

It is especially useful when we have a legacy system or third-party library that we need to integrate with, but it doesn’t use the same interface as the rest of our code.

I recently used an Adapter design pattern in one of the projects I am working on. The project required integrating multiple payment service providers from various regions. Integrating multiple 3rd party providers directly into our code can make it difficult to understand, tightly coupled, and difficult to segregate. While to rest of the code should benefit with each new provider, but should not be aware of which provider are we exactly using. To overcome this challenge we can define an interface that consists of a set of methods that a new provider should implement for it to implement that adapter interface.

```
type Payments interface {
 InitTransaction()
 GetTransaction()
 RefundTransaction()
}
```

As you can see above we have defined an interface called Payments that has a set of methods to initialize a transaction, to get a transaction details, and to refund a transaction. In Go, we can create separate packages of each new Payment Service Provider to have a non-dependent unit of code with a structure Adapter that implements all these interface methods by integrating the API endpoints exposed by each Provider.

```
type Adapter struct {
    api API
}

func (a *Adapter) InitTransaction() {}

func (a *Adapter) GetTransaction() {}

func (a *Adapter) RefundTransaction() {}
```

In this way by using the Adapter design pattern our third-party providers were able to adapt to the client's requirements. We can make available just the required features from the third-party provider according to the needs of the application and hide all the extra features that it provides. 

## Benefits

* Helps in maintaining SRP (Single Responsibility Principle) and Interface Segregation principle by dividing and wrapping complex interfaces into simpler ones.
* Improves code reusability, as it helps write modular code by writing separate independent packages. 
Ensures flexibility and loose coupling by reducing direct dependencies between different parts of the code.
* Helps in grouping all the code in our system that deals with a particular dependency into packages, modules, or functions.
* Also helps in differentiating the business logic and dependencies, thus helps in managing different parts of code and enhancing code quality.
* It also allows mocking the dependencies to enable effective unit testing.
* Makes the code easier to test, extend, and manage.

## Conclusion

In this way, the Adapter design pattern can be very useful to integrate new components with minimal changes and helps to write a well-structured and maintainable codebase. While it adds a layer of abstraction, it should not be overused as it can lead to unnecessary complexity.
