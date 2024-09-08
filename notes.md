///////////////////////////////////////////////////////////////////////
// gRPC: Up and Running book //////////////////////////////////////////
///////////////////////////////////////////////////////////////////////



///////////////////////////////////////////////////////////////////////
// Chapter 01 /////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Microservices architecture is about building a software application
as a collection of independent, autonomous (developed, deployed,
and scaled independently), business capability–oriented, and
loosely coupled services.

2. Inter-process communications are usually implemented using message passing with a
synchronous request-response style or asynchronous event-driven styles. In the syn‐
chronous communication style, the client process sends a request message to the
server process over the network and waits for a response message.

3. In asynchronous event-driven messaging, processes communicate with asynchronous message
passing by using an intermediary known as an event broker. Depending on your business use
case, you can select the communication pattern that you want to implement.

4. gRPC is an interprocess communication technology that allows you to connect, invoke,
operate, and debug distributed heterogeneous applications as easily as making a local function call.

5. When you develop a gRPC application the first thing that you do is define a service
interface.
    - The service interface definition contains information on how your service
    can be consumed by consumers.
    - what methods you allow the consumers to call remotely.
    - what method parameters and message formats to use when invoking those
    methods.
    - The language that we specify in the service definition is known
    as an interface definition language (IDL).

6. Using that service definition:
    - you can generate the server-side code known as a server
    skeleton, which simplifies the server-side logic by providing low-level communication
    abstractions.
    - you can generate the client-side code, known as a client stub,
    which simplifies the client-side communication with abstractions to hide low-level
    communication for different programming languages.

7. gRPC uses protocol buffers as the IDL to define the service interface. Protocol buffers
are a language-agnostic, platform-neutral, extensible mechanism to serializing struc‐
tured data.

8. On the server side, the server implements that service definition and runs a gRPC
server to handle client calls. Therefore, on the server side, to make the ProductInfo
service do its job you need to do the following:
    1. Implement the service logic of the generated service skeleton by overriding the
    service base class.
    2. Run a gRPC server to listen for requests from clients and return the service
    responses.

9. Marshaling is the process of packing parameters and a remote function into a message
packet that is sent over the network,
while unmarshaling unpacks the message packet into the respective method invocation.

10. RPC was a popular inter-process communication technique for building client-
service applications. With RPC a client can remotely invoke a function of a method
just like calling a local method.

11. SOAP is the standard communication technique in a service-oriented architecture (SOA) to exchange XML-based structured data between services (usually called web services in the context of SOA) and communicates over any underlying communication protocol such as HTTP (most commonly used).

12. Representational State Transfer (REST) is an architectural style that originated from
Roy Fielding’s PhD dissertation. Fielding is one of the principal authors of the HTTP
specification and the originator of the REST architectural style. REST is the founda‐
tion of the resource-oriented architecture (ROA), where you model distributed appli‐
cations as a collection of resources and the clients that access those resources can
change the state (create, read, update, or delete) of those resources.

13. Disadvantages of gRPC
    - It may not be suitable for external-facing services
    - Drastic service definition changes are a complicated development process
    - The ecosystem is relatively small

14. Summary
    Modern software applications or services rarely live in isolation and the inter-process
    communication techniques that connect them are one of the most important aspects
    of modern distributed software applications. gRPC is a scalable, loosely coupled, and
    type-safe solution that allows for more efficient inter-process communication than
    conventional REST/HTTP-based communication. It allows you to connect, invoke,
    operate, and debug distributed heterogeneous applications as easy as making a local
    method call via network transport protocols such as HTTP/2.

///////////////////////////////////////////////////////////////////////
// Chapter 02 /////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Message: is the data structure that is exchanged between client and service.

2. According to the protocol buffer rule, we can only have one
   input parameter in a remote method and it can return only one value.

3. Context object contains metadata such as the identity of the end user,
   authorization tokens, and the request’s deadline and it will exist during the lifetime of the request.

///////////////////////////////////////////////////////////////////////
// Chapter 03 /////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Simple RPC (Unary RPC)
   when a client invokes a remote function of a server, the client sends a single request
   to the server and gets a single response that is sent along with status details and trailing metadata.

2. Server-Streaming RPC
   In server-side streaming RPC, the server sends back a sequence of responses after getting the client’s request message.
   This sequence of multiple responses is known as a “stream.” After sending all the
   server responses, the server marks the end of the stream by sending the server’s status
   details as trailing metadata to the client.

3. Client-Streaming RPC
   In client-streaming RPC, the client sends multiple messages to the server instead of a
   single request. The server sends back a single response to the client.

4. Bidirectional-Streaming RPC
    In bidirectional-streaming RPC, the client is sending a request to the server as a
    stream of messages. The server also responds with a stream of messages. The call has
    to be initiated from the client side, but after that, the communication is completely
    based on the application logic of the gRPC client and the server.
- Each order ID is sent to the server as a separate gRPC message.
- The service processes each order for the specified order ID and organizes them
  into combined shipments based on the delivery location of the order.
- The key idea behind this business use case is that once the RPC method is invoked
  either the client or service can send messages at any arbitrary time. (This also
  includes the end of stream markings from either of the parties.)

5. Using gRPC for Microservices Communication
    In most of the real-world use cases, these external-facing services are exposed
    through an API gateway. That is the place where you apply various nonfunctional
    capabilities such as security, throttling, versioning, and so on. Most such APIs lever‐
    age protocols such as REST or GraphQL. Although it’s not very common, you may
    also expose gRPC as an external-facing service, as long as the API gateway supports
    exposing gRPC interfaces.
6. By using an API gateway with your gRPC APIs, you are able to deploy this functionality
    outside of your core gRPC services. One of the other important aspects of this architec‐
    ture is that we can leverage multiple programming languages but share the same
    service contract between then (i.e., code generation from the same gRPC service
    This allows us to pick the appropriate implementation technology based
    on the business capability of the service.

///////////////////////////////////////////////////////////////////////
// Chapter 04 /////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////
