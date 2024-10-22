
# Chapter 01

> Go provides a viable alternative to existing languages and platforms for developing
large-scale web applications. Large-scale web applications typically need to be
- Scalable
- Modular
- Maintainable
- High-performance

- Go scales well vertically with its excellent support for concurrent programming. A single Go web application with a single OS thread can be scheduled to run hundreds of thousands of goroutines with efficiency and performance.

- Go web applications are compiled as static binaries, without any dynamic dependencies, and can be distributed to systems that don’t have Go built in.

- Web application is a computer program that responds to an HTTP request by a client and sends HTML back to the client in an HTTP response.

- In comparison, a web application doesn’t simply return files; it processes the request and performs operations that are programmed into the application.


![[Screenshot from 2024-10-14 13-58-38 1.png]]

An application: is a software program that interacts with a user and helps the user to
perform an activity.

if a program doesn’t render HTML to a user but
instead returns data in any other format to another program, it is a web service


HTTP is a stateless, text-based, request-response protocol that uses the
client-server computing model.

In HTTP , the client is also known as the user-agent
and is often a web browser. The server is often called the web server.

### HTTP request
1- Request-line
2- Zero or more request headers
3- An empty line
4- The message body (optional)


 - GET:  Tells the server to return the specified resource.

 - HEAD:  The same as GET except that the server must not return a message
body. This method is often used to get the response headers without carrying
the weight of the rest of the message body over the network.

 - POST:  Tells the server that the data in the message body should be passed to
the resource identified by the URI. What the server does with the message body
is up to the server.

 - PUT:  Tells the server that the data in the message body should be the resource
at the given URI. If data already exists at the resource identified by the URI, that
data is replaced. Otherwise, a new resource is created at the place where the
URI is.

 - DELETE:  Tells the server to remove the resource identified by the URI.

 - TRACE:  Tells the server to return the request. This way, the client can see what
the intermediate servers did to the request.

 - OPTIONS:  Tells the server to return a list of HTTP methods that the server sup-
ports.

 - CONNECT:  Tells the server to set up a network connection with the client. This
method is used mostly for setting up SSL tunneling (to enable HTTPS).

 - PATCH:  Tells the server that the data in the message body modifies the
resource identified by the URI.

**Safe request methods**
A method is considered safe if it doesn’t change the state of the server—that is, the
server provides only information and nothing else. **GET , HEAD, OPTIONS, and TRACE**
are safe methods.

**Idempotent request methods**
A method is considered idempotent if the state of the server doesn’t change the second
time.

**Request headers**
Although the HTTP request method defines the action requested by the calling client,
other information on the request or the client is often placed in HTTP request headers. Request headers are colon-separated name-value pairs in plain text, terminated by
a carriage return (CR) and line feed (LF).

![[Screenshot from 2024-10-16 10-53-43.png]]



### HTTP response
An HTTP response message is sent every time there’s a request. Like the HTTP
request, the HTTP response consists of a few lines of plain text:
 - A status line
 - Zero or more response headers
 - An empty line
 - The message body (optional)

**Response status code**
![[Screenshot from 2024-10-16 11-14-31.png]]


**Response headers**
headers tell the server more about the request and what the client wants, the response
headers are the means for the server to tell the client more about the response and
what the server wants (from the client).
![[Screenshot from 2024-10-16 11-15-47.png]]


**URI**
This is the general form of a URI:
<scheme name> : <hierarchical part> [ ? <query> ][ # <fragment> ]

The fragment starts after the hash (#). If a URI has a query, the fragment will follow the query.
 The fragment is meant to be processed by the client, so web browsers normally strip the fragment out before sending the URI to the server.


Let’s look at an example of an HTTP scheme URI:
http://sausheong:password@www.example.com/docs/filename=sausheong&location=singapore#summary

The scheme is http, followed by the colon. The segment sausheong:password fol-
lowed by the at sign (@) is the user and password information. This is followed by the
rest of the hierarchical part, www.example.com/docs/file. The top level of the hierarchi-
cal part is the domain name of the server, www.example.com, followed on by docs and
then file, each separated by a forward slash. Next is the query, which begins after the
question mark (?). The query consists of two name-value pairs: name=sausheong and
location=singapore, joined by a single ampersand (&). Finally, the fragment follows
after the query, starting after the hash (#).


**RFC 3986** defines a set of characters that are reserved or not reserved. Everything
in the reserved list needs to be URL encoded.




---


**HTTP/2**
  is a binary protocol, unlike HTTP/1.x, which is text-based. This makes
HTTP/2 more efficient to parse, and it is more compact and less prone for errors. But
that means you can no longer send HTTP/2 messages directly through the network,
through applications such as telnet, and so it is harder to debug if you’re used to
HTTP/1.x.



**Parts of a web app**
From the previous sections you’ve seen that a web application is a piece of program
that does the following:
1. Takes input through HTTP from the client in the form of an HTTP request
message
2. Processes the HTTP request message and performs necessary work
3. Generates HTML and returns it in an HTTP response message
As a result, there are two distinct parts of a web app: the **handlers** and the **template** engine.

**Handler**
A handler receives and processes the HTTP request sent from the client. It also calls the template engine to generate the HTML and finally bundles data into the HTTP response to be sent back to the client.
Sometimes **service objects** or functions are used to manipulate the models, freeing the model from being too bloated and enabling reuse of code. In this case, service objects can be reused on different models and the same logic can be placed in single service object instead of being copied in different models.


**Template engine**
A template is code that can be converted into HTML that’s sent back to the client in an
HTTP response message. Templates can be partly in HTML or not at all.
**A template engine** generates the final HTML using templates and data. As you may recall, template engines evolved from an earlier technology, SSI.
There are two types of templates with different design philosophies:
- Static templates or logic-less templates are HTML interspersed with placeholder
tokens. A static template engine will generate the HTML by replacing these
tokens with the correct data. There’s little to no logic in the template itself. As
you can see, this is similar to the concepts from SSI. Examples of static template
engines are CTemplate and Mustache.
- Active templates often contain HTML too, but in addition to placeholder tokens,
they contain other programming language constructs like conditionals, itera-
tors, and variables. Examples of active template engines are Java ServerPages
(JSP), Active Server Pages (ASP), and Embedded Ruby (ERB). PHP started off as
a kind of active template engine and has evolved into its own programming
language.


## Chapter 02

1. **Application design**

we’ll be using the following format: http://<servername>/<handler-
name>?<parameters>
http://chitchat/thread/read?id=123.

When the request reaches the server, a **multiplexer** will inspect the URL being
requested and redirect the request to the correct handler.
![[Screenshot from 2024-10-20 17-21-47.png]]

2. **Data model**

The four data structures are:
 1. User : Representing the forum user’s information
 2. Session : Representing a user’s current login session
 3. Thread : Representing a forum thread (a conversation among forum users)
 4. Post : Representing a post (a message added by a forum user) within a thread
![[Screenshot from 2024-10-20 17-28-32.png]]

3. Receiving and processing requests
Receiving and processing requests is the heart of any web application. Let’s recap
what you’ve learned so far:
A client sends a request to a URL at the server.
The server has a multiplexer, which redirects the request to the correct handler
to process the request.
The handler processes the request and performs the necessary work.
The handler calls the template engine to generate the correct HTML to send
back to the client.
