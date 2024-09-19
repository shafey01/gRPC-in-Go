///////////////////////////////////////////////////////////////////////
// gRPC: Up and Running book //////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////
// Chapter 01 /////////////////////////////////////////////////////////
// Introduction to gRPC ///////////////////////////////////////////////
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
   interface. - The service interface definition contains information on how your service
   can be consumed by consumers. - what methods you allow the consumers to call remotely. - what method parameters and message formats to use when invoking those
   methods. - The language that we specify in the service definition is known
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
   service do its job you need to do the following: 1. Implement the service logic of the generated service skeleton by overriding the
   service base class. 2. Run a gRPC server to listen for requests from clients and return the service
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
// Getting Started with gRPC //////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Message: is the data structure that is exchanged between client and service.

2. According to the protocol buffer rule, we can only have one
   input parameter in a remote method and it can return only one value.

3. Context object contains metadata such as the identity of the end user,
   authorization tokens, and the request’s deadline and it will exist during the lifetime of the request.

///////////////////////////////////////////////////////////////////////
// Chapter 03 /////////////////////////////////////////////////////////
// gRPC Communication Patterns ////////////////////////////////////////
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
//gRPC: Under the Hood ////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. In gRPC, all requests are HTTP POST requests with content-type prefixed with
   application/grpc. The remote function (/ProductInfo/getProduct) that it
   invokes is sent as a separate HTTP header.

2. When the message is received at the server, the server examines the message
   headers to see which service function needs to be called and hands over the mes‐
   sage to the service stub.

3. Message = Tag + Value
   Tag = (Field Index << 3) | wire type // wire type like: 0 for int32, int64, uint32, uint64, sint32, sint64, bool, enum
   // 2 for string, bytes, embedded messages, packed repeated fields

4. Encoding Techniques // https://protobuf.dev/programming-guides/encoding/
   string values are encoded using UTF-8 character encoding,
   whereas int32 values are encoded using varints.

- Varints
  Varints (variable length integers) are a method of serializing integers using one or more bytes.
  int32, int64, uint32, uint64, sint32, sint64, bool, enum.
- For negative integer values, it is recommended to use signed integer types like sint32
  and sint64 because if we use a regular type such as int32 or int64, negative values are
  converted to binary using varints encoding.
- Nonvarint numbers
  Nonvarint types are just the opposite of the varint type. They allocate a fixed number
  of bytes irrespective of the actual value. Protocol buffers use two wire types that cate‐gorize
  as nonvarint numbers. One is for the 64-bit data types like fixed64, sfixed64,
  and double. The other is for 32-bit data types like fixed32, sfixed32, and float.
- String type
  In protocol buffers, the string type belongs to the length-delimited wire type, which
  means that the value is a varint-encoded length followed by the specified number of
  bytes of data. String values are encoded using UTF-8 character encoding.

5. gRPC uses a message-framing technique called length-prefix framing.
   Length-prefix is a message-framing approach that writes the size of each message
   before writing the message itself. As you can see in Figure 4-4, before the encoded
   binary message there are 4 bytes allocated to specify the size of the message. In gRPC
   communication, 4 additional bytes are allocated for each message to set its size. The
   size of the message is a finite number, and allocating 4 bytes to represent the message
   size means gRPC communication can handle all messages up to 4 GB in size.

6. In addition to the message size, the frame also has a 1-byte unsigned integer to indi‐
   cate whether the data is compressed or not. A Compressed-Flag value of 1 indicates
   that the binary data is compressed using the mechanism declared in the Message-Encoding header,
   which is one of the headers declared in HTTP transport. The value 0 indicates
   that no encoding of message bytes has occurred.

7. The recipient is the client on the recipient side, once a message is received, it first
   needs to read the first byte to check whether the message is compressed or not. Then
   the recipient reads the next four bytes to get the size of the encoded binary message.
   Once the size is known, the exact length of bytes can be read from the stream. For
   unary/simple messages, we will have only one length-prefixed message, and for
   streaming messages, we will have multiple length-prefixed messages to process.

8. HTTP/2 is the second major version of the internet protocol HTTP. It was introduced
   to overcome some of the issues encountered with security, speed, etc. in the previous
   version (HTTP/1.1). HTTP/2 supports all of the core features of HTTP/1.1 but in a
   more efficient way. So applications written in HTTP/2 are faster, simpler, and more
   robust.

9. Request Message
   The request message is the one that initiates the remote call. In gRPC, the request
   message is always triggered by the client application and it consists of three main
   components: request headers, the length-prefixed message, and the end of stream flag.

10. Request header
    HEADERS (flags = END_HEADERS)
    :method = POST 1
    :scheme = http 2
    :path = /ProductInfo/getProduct 3
    :authority = abnc.com 4
    te = trailers 5
    grpc-timeout = 1S 6
    content-type = application/grpc 7
    grpc-encoding = gzip 8
    authorization = Bearer xxxxxx 9
11. Defines the (HTTP method. For gRPC, the :method header is always POST.
12. Defines the 8HTTP scheme. If TLS (Transport Level Security) is enabled, the scheme is set to “https,” otherwise it is “http.”
13. Defines the dendpoint path. For gRPC, this value is constructed as “/” {service
    name} “/” {9method name}.
14. Defines the virtual hostname of the target URI.
15. Defines detection of incompatible proxies. For gRPC, the value must be “trailers.”
16. Defines call timeout. If not specified, the server should assume an infinite
    timeout.
17. Defines the content-type. For gRPC, the content-type should begin with applica
    tion/grpc. If not, gRPC servers will respond with an HTTP status of 415
    (Unsupported Media Type).
18. Defines the message compression type. Possible values are identity, gzip,
    deflate, snappy, and {custom}.
19. This is optional metadata. authorization metadata is used to access the secure
    endpoint.

20. Headers passed in gRPC communication are categorized into
    two types: call-definition headers and custom metadata.

21. Call-definition headers are predefined headers supported by
    HTTP/2. Those headers should be sent before custom meta‐
    data.
22. Custom metadata is an arbitrary set of key-value pairs defined
    by the application layer. When you are defining custom meta‐
    data, you need to make sure not to use a header name starting
    with grpc-. This is listed as a reserved name in the gRPC core.

23. Response Message
    The response message is generated by the server in response to the client’s request.
    Similar to the request message, in most cases the response message also consists of
    three main components: response headers, length-prefixed messages, and trailers.

24. Response header
    HEADERS (flags = END_HEADERS)
    :status = 200 1
    grpc-encoding = gzip 2
    content-type = application/grpc 3
25. Indicates the status of the HTTP request.
26. Defines the message compression type. Possible values include identity, gzip, deflate, snappy, and {custom}.
27. Defines the content-type. For gRPC, the content-type should begin with application/grpc.

28. the END_STREAM flag isn’t sent with data frames. It is sent as a separate
    header called a trailer:
    DATA
    <Length-Prefixed Message>
    In the end, trailers are sent to notify the client that we finished sending the response
    message. Trailers also carry the status code and status message of the request:
    HEADERS (flags = END_STREAM, END_HEADERS)
    grpc-status = 0 # OK 1
    grpc-message = xxxxxx 2
29. Defines the gRPC status code. gRPC uses a set of well-defined status codes. You
    can find the definition of status codes in the official gRPC documentation. //https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
30. Defines the description of the error. This is optional. This is only set when there
    is an error in processing the request.

31. In certain scenarios, there can be an immediate failure in the request call. In those
    cases, the server needs to send a response back without the data frames. So the server
    sends only trailers as a response. Those trailers are also delivered as an HTTP/2
    header frame and also contain the END_STREAM flag. Additionally, the following headers are included in trailers:
    • HTTP-Status → :status
    • Content-Type → content-type
    • Status → grpc-status
    • Status-Message → grpc-message

32. Here “half-close the connection” means the client closes the connection
    on its side so the client is no longer able to send messages to the server but still
    can listen to the incoming messages from the server.

33. Simple RPC

34. Server-streaming RPC

35. Client-streaming RPC

36. Bidirectional-streaming RPC

37. Summary
    gRPC builds on top of two fast and efficient protocols called protocol buffers and
    HTTP/2. Protocol buffers are a data serialization protocol that is a language-agnostic,
    platform-neutral, and extensible mechanism for serializing structured data. Once
    serialized, this protocol produces a binary payload that is smaller in size than a normal JSON
    payload and is strongly typed. This serialized binary payload then travels
    over the binary transport protocol called HTTP/2.
    HTTP/2 is the next major version of the internet protocol HTTP. HTTP/2 is fully
    multiplexed, which means that HTTP/2 can send multiple requests for data in parallel
    over a single TCP connection. This makes applications written in HTTP/2 faster, simpler, and more robust than others

///////////////////////////////////////////////////////////////////////
// Chapter 05 /////////////////////////////////////////////////////////
// gRPC: Beyond the Basics ////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Interceptors
   In gRPC you can intercept that RPC’s execution to meet certain requirements such as
   logging, authentication, metrics, etc., using an extension mechanism called an interceptor.

- gRPC interceptors can be categorized into two types based on the type of RPC calls
  they intercept. For unary RPC you can use unary interceptors, while for streaming
  RPC you can use streaming interceptors. These interceptors can be used on the gRPC
  server side or on the gRPC client side

2. Server-Side Interceptors
   When a client invokes a remote method of a gRPC service, you can execute a common logic
   prior to the execution of the remote methods by using a server-side interceptor.

- Unary interceptor can usually be divided into
  three parts: preprocessing, invoking the RPC method, and postprocessing.
  - As the name implies, the preprocessor phase is executed prior to invoking the remote
    method intended in the RPC call. In the preprocessor phase, users can get info about
    the current RPC call by examining the args passed in, such as RPC context, RPC
    request, and server information. Thus, during the preprocessor phase you can even
    modify the RPC call.
  - Then, in the invoker phase, you have to call the gRPC UnaryHandler to invoke the
    RPC method. Once you invoke the RPC, the postprocessor phase is executed. This
    means that the response for the RPC call goes through the postprocessor phase. In
    the phase, you can deal with the returned reply and error when required. Once the
    postprocessor phase is completed, you need to return the message and the error as
    the return parameters of your interceptor function. If no postprocessing is required,
    you can simply return the handler call (handler(ctx, req)).
- Stream interceptor
  - The server-side streaming interceptor intercepts any streaming RPC calls that the
    gRPC server deals with. The stream interceptor includes a preprocessing phase and a
    stream operation interception phase.
  - Similar to a unary interceptor, in the preprocessor phase, you can intercept a streaming
    RPC call before it goes to the service implementation. After the preprocessor
    phase, you can then invoke the StreamHandler to complete the execution of RPC
  - invocation of the remote method. After the preprocessor phase, you can intercept the
    streaming RPC message by using an interface known as a wrapper stream that implements
    the grpc.ServerStream interface. You can pass this wrapper structure when
    you invoke grpc.StreamHandler with handler(srv, newWrappedStream(ss)).
  - The wrapper of grpc.ServerStream intercepts the streaming messages sent or received by
    the gRPC service. It implements the SendMsg and RecvMsg functions, which will be
    invoked when the service receives or sends an RPC streaming message.

3. Client-Side Interceptors
   When a client invokes an RPC call to invoke a remote method of a gRPC service, you
   can intercept those RPC calls on the client side. with client-side interceptors,
   you can intercept unary RPC calls as well as streaming RPC calls.

- This is particularly useful when you need to implement certain reusable features, such
  as securely calling a gRPC service outside the client application code.
- Unary interceptor
  - A client-side unary RPC interceptor is used for intercepting the unary RPC client
    side. UnaryClientInterceptor is the type for a client-side unary interceptor that has
    a function signature as follows:
    func(ctx context.Context, method string, req, reply interface{},
    cc \*ClientConn, invoker UnaryInvoker, opts ...CallOption) error
  - In the preprocessor phase, you can intercept the RPC
    calls before invoking the remote method.
  - Registering the interceptor function is done inside the grpc.Dial operation using
    grpc.WithUnaryInterceptor.
- Stream interceptor
  The client-side streaming interceptor intercepts any streaming RPC calls that the
  gRPC client deals with. The implementation of the client-side stream interceptor is
  quite similar to that of the server side. StreamClientInterceptor is the type for a
  client-side stream interceptor.
  a function type with this signature:
  func(ctx context.Context, desc *StreamDesc, cc *ClientConn,
  method string, streamer Streamer,
  opts ...CallOption) (ClientStream, error)

4. Deadlines

   - Timeouts allow you to specify how long a client application can wait for an RPC to
     complete before it terminates with an error. A timeout is usually specified as a duration
     and locally applied at each client side.
   - For example, a single request may consist of multiple downstream RPCs
     that chain together multiple services. So we can apply timeouts,
     relative to each RPC, at each service invocation. Therefore, timeouts cannot
     be directly applied for the entire life cycle of the request. That’s where we need to use deadlines.
   - Deadline is expressed in absolute time from the beginning of a request (even if the
     API presents them as a duration offset) and applied across multiple service invoca‐
     tions. The application that initiates the request sets the deadline and the entire
     request chain needs to respond by the deadline.
   - deadline = current time + offset
   - A client application can set a deadline when it initiates a connection with a gRPC service.
     Once the RPC call is made, the client application waits for the duration specified
     by the deadline; if the response for the RPC call is not received within that time, the
     RPC call is terminated with a DEADLINE_EXCEEDED error.

5. Cancellation

   - When either the client or server application wants to terminate the RPC this can be done
     by canceling the RPC. Once the RPC is canceled, no further RPC-related messaging can be
     done and the fact that one party has canceled the RPC is propagated to the other side.
   - When one party cancels the RPC, the other party can determine it by checking the
     context. In this example, the server application can check whether the current context
     is canceled by using stream.Context().Err() == context.Canceled.

6. Error Handling

   - When an error occurs, gRPC returns one of its error-status codes with an optional
     error message that provides more details of the error condition. The status object is
     composed of an integer code and a string message that are common to all gRPC
     implementations for different languages.
   - gRPC uses a set of well-defined gRPC-specific status codes. This includes status codes
     such as the following:
   - OK: Successful status; not an error.
   - CANCELLED: The operation was canceled, typically by the caller.
   - DEADLINE_EXCEEDED: The deadline expired before the operation could complete.
   - INVALID_ARGUMENT: The client specified an invalid argument.

   - It’s always good practice to use the appropriate gRPC error codes and a richer error
     model whenever possible for your gRPC applications. gRPC error status and details
     are normally sent via the trailer headers at the transport protocol level.

7. Multiplexing
   gRPC allows you to run multiple gRPC services on the same gRPC server,
   as well as use the same gRPC client connection for multiple gRPC client stubs.
   This capability is known as multiplexing.

   - Running multiple services or using the same connection between multiple stubs is a
     design choice that is independent of gRPC concepts. In most everyday use cases such
     as microservices, it is quite common to not share the same gRPC server instance
     between two services.
   - One powerful use for gRPC multiplexing in a microservice architecture is
     to host multiple major versions of the same service in one
     server process. This allows a service to accommodate legacy clients
     after a breaking API change. Once the old version of the service
     contract is no longer in use, it can be removed from the server.

8. Metadata
   in certain conditions, you may want to share information about the RPC calls that are
   not related to the business context of the RPC, so they shouldn’t be part of the RPC
   arguments. In such cases, you can use gRPC metadata
   Metadata is structured in the form of a list of key(string)/value pairs.

- One of the most common usages of metadata is to exchange security headers between
  gRPC applications. Similarly, you can use it to exchange any such information
  between gRPC applications. Often gRPC metadata APIs are heavily used inside the
  interceptors that we develop.
- Reading metadata from either the client or server side can be done using the incoming
  context of the RPC call with metadata.FromIncomingContext(ctx), which returns the metadata map in Go

9. Name Resolver
   A name resolver takes a service name and returns a list of IPs of the backends. The
   resolver used in Example 5-15 resolves lb.example.grpc.io to localhost:50051
   and localhost:50052.

- Thus, based on this name resolver implementation, you can implement resolvers for
  any service registry of your choice such as Consul, etcd, and Zookeeper. The gRPC
  load-balancing requirements may be quite dependent on the deployment patterns
  that you use or on the use cases.

10. Load Balancing
    two main load-balancing mechanisms are commonly used in
    gRPC: a load-balancer (LB) proxy and client-side load balancing.

    - Load-Balancer Proxy
      In proxy load balancing the client issues RPCs to the LB proxy. Then the
      LB proxy distributes the RPC call to one of the available backend gRPC servers that
      implements the actual logic for serving the call.
    - Client-Side Load Balancing
      Rather than having an intermediate proxy layer for load balancing, you can imple‐
      ment the load-balancing logic at the gRPC client level. In this method, the client is
      aware of multiple backend gRPC servers and chooses one to use for each RPC.
      - the load-balancing logic may be entirely developed as part of
        the client application (also known as thick client) or it can be implemented in a dedi‐
        cated server known as lookaside load balancer.
    - There are two load-balancing policies supported in gRPC by default: pick_first and
      round_robin. pick_first tries to connect to the first address, uses it for all RPCs if it
      connects, or tries the next address if it fails. round_robin connects to all the addresses
      it sees and sends an RPC to each backend one at a time in order.

11. Compression
    To use network bandwidth efficiently, use compression when performing RPCs
    between client and services. Using gRPC compression on the client side can be imple‐
    mented by setting a compressor when you do the RPC.

///////////////////////////////////////////////////////////////////////
// Chapter 06 /////////////////////////////////////////////////////////
// Secured gRPC ///////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. Authenticating a gRPC Channel with TLS
   Transport Level Security (TLS) aims to provide privacy and data integrity between two
   communicating applications. Here, it’s about providing a secure connection between
   gRPC client and server applications.

2. One-Way Secured Connection
   In a one-way connection, only the client validates the server to ensure that it receives
   data from the intended server. When establishing the connection between the client
   and the server, the server shares its public certificate with the client, who then vali‐
   dates the received certificate.

3. mTLS Secured Connection
   The main intent of an mTLS connection between client and server is to have control
   of clients who connect to the server. Unlike a one-way TLS connection, the server is
   configured to accept connections from a limited group of verified clients. Here both
   parties share their public certificates with each other and validate the other party. The
   basic flow of connection is as follows:

   1. Client sends a request to access protected information from the server.
   2. The server sends its X.509 certificate to the client.
   3. Client validates the received certificate through a CA for CA-signed certificates.
   4. If the verification is successful, the client sends its certificate to the server.
   5. Server also verifies the client certificate through the CA.
   6. Once it is successful, the server gives permission to access protected data.

   - server.key
     Private RSA key of the server.
   - server.crt
     Public certificate of the server.
   - client.key
     Private RSA key of the client.
   - client.crt
     Public certificate of the client.
   - ca.crt
     Public certificate of a CA used to sign all public certificates.

4. Authenticating gRPC Calls
   In order to facilitate verification of the caller, gRPC provides the capability for the cli‐
   ent to inject his or her credentials (like username and password) on every call. The
   gRPC server has the ability to intercept a request from the client and check these cre‐
   dentials for every incoming call.

- Using Basic Authentication
  is the simplest authentication mechanism. In this mechanism,
  the client sends requests with the Authorization header with a value that starts with
  the word Basic followed by a space and a base64-encoded string username:pass
  word. For example, if the username is admin and the password is admin, the header
  value looks like this:
  Authorization: Basic YWRtaW46YWRtaW4=
- Using OAuth 2.0
  In the OAuth 2.0 flow, there are four main characters:

  1. the client,
  2. the authorization server,
  3. the resource server,
  4. the resource owner.

- Using JWT
  JWT defines a container to transport identity information between the client and
  server. A signed JWT can be used as a self-contained access token, which means the
  resource server doesn’t need to talk to the authentication server to validate the client
  token. It can validate the token by validating the signature.

- Using Google Token-Based Authentication
  Identifying the users and deciding whether to let them use the services deployed on
  the Google Cloud Platform is controlled by the Extensible Service Proxy (ESP). ESP
  supports multiple authentication methods, including Firebase, Auth0, and Google ID tokens.
  In each case, the client needs to provide a valid JWT in their requests. In order to generate
  authenticating JWTs, we must create a service account for each deployed service.

5. Summary
   There are two types of credential supports in gRPC, channel and call. Channel cre‐
   dentials are attached to the channels such as TLS, etc. Call credentials are attached to
   the call, such as OAuth 2.0 tokens, basic authentication, etc. We even can apply both
   credential types to the gRPC application. For example, we can have TLS enable the
   connection between client and server and also attach credentials to each RPC call
   made on the connection.

///////////////////////////////////////////////////////////////////////
// Chapter 08 /////////////////////////////////////////////////////////
// The gRPC Ecosystem /////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////

1. gRPC Gateway // https://github.com/grpc-ecosystem/grpc-gateway
   The gRPC gateway plug-in enables the protocol buffer compiler to read the gRPC
   service definition and generate a reverse proxy server, which translates a RESTful
   JSON API into gRPC.
2. HTTP to gRPC mapping:
   • Each mapping needs to specify a URL path template and an HTTP method.
   • The path template can contain one or more fields in the gRPC request message.
   But those fields should be nonrepeated fields with primitive types.
   • Any fields in the request message that are not in the path template automatically
   become HTTP query parameters if there is no HTTP request body.
   • Fields that are mapped to URL query parameters should be either a primitive
   type or a repeated primitive type or a nonrepeated message type.
   • For a repeated field type in query parameters, the parameter can be repeated in
   the URL as ...?param=A&param=B.
   • For a message type in query parameters, each field of the message is mapped to a
   separate parameter, such as ...?foo.a=A&foo.b=B&foo.c=C.
3. HTTP/JSON Transcoding for gRPC
   Transcoding is the process of translating HTTP JSON calls to RPC calls and passing
   them to the gRPC service.
4. The gRPC Server Reflection Protocol
   Server reflection is a service defined on a gRPC server that provides information
   about publicly accessible gRPC services on that server. In simple terms, server reflection
   provides service definitions of the services registered on a server to the client
   application. So the client doesn’t need precompiled service definitions to communicate with the services.
5. gRPC Middleware
   he middleware is a software component in a distributed system that is
   used to connect different components to route requests generated by the client to the
   backend server. In gRPC Middleware, we are also talking about running code before
   and after the gRPC server or client application.
6. Health Checking Protocol
   gRPC defines a health checking protocol (Health Checking API) that allows the gRPC
   services to expose the server status so that the consumers can probe the server’s
   health information.
7. gRPC Health Probe
   The grpc_health_probe is a utility provided by the community to check the health
   status of a server that exposes its status as a service through the gRPC Health Checking Protocol.
