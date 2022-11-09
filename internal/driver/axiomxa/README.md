# RBH AxiomXa Driver

Driver for the AxiomXa product version, produced by RBH. The server is an access control system which consists of an "
Axiom Server" which provides the vendor specific UI for managing the estate. This server also hosts any API access
available.

AxiomXa provides three different methods of integration, each with benefits and drawbacks

1. HTTP API
2. Message Ports
3. Direct DB access

## HTTP API

The HTTP API has to be enabled by AxiomXa contractors with special licencing. By default it is HTTP only but is hosted
within IIS and it is the responsibility of the site administrator to manage the certificates and setup TLS for the API.
The API is quite basic and the documentation is NDA-walled, it's consists of what looks like generated API docs from
some internal representation, edited and placed into PDF (likely via Word).

The API is not quite restful, but does use JSON as the payload. While GET and POST verbs are used, you'll find urls like
this to get a single item:

```http request
# Get a single access level
POST /v1/accesslevel/one

{ "ID": "string" }
```

The API is authenticated using a simple `/v1/login` API which returns a bearer token to include in the Authorization
header of future requests.

There is no event or notification mechanism using this API.

## Message Ports

Invented by RBH, think _outgoing webhooks_. These are configured per AxiomXa install, typically by the people
maintaining that system. They are quite flexible but also limited. They use TCP not HTTP, they have a payload limit of
50 characters but can notify you of changes and events that happen in all sorts of areas within the AxiomXa system.

The payload of the messages is not defined at all. During commissioning the AxiomXa engineer sets up templates for the
payload that resembles `EVENTTIME CARDID EVENTTYPE ALLOWED` in the AxiomXa UI, they set up a trigger condition and a
target to send it to and the system replaces the template keywords and sends the data down the socket.

Here are some of the available template placeholder values (called inserts):

- `TIMESTAMP`:  Date & Time of the event, acquired from the event message.
- `EVENTID`:    Identification number associated with the event.
- `EVENTDES`:   Description of the event, acquired from the event message.
- `NETWORKID`:  Identification number associated with the network of the event.
- `NETWORKDES`: Description of the network, associated with the event message.
- `NC100ID`:    Identification number associated with the NC100 of the event.
- `NC100DES`:   Description of the NC100, associated with the event message.
- `DEVICESID`:  Identification number associated with the device (RC2, IOC16, or SafeSuiteTM panel) of the event.
- `DEVICEDES`:  The description of the device (RC2, IOC16, or SafeSuiteTM panel) associated with the event message.
- `CARDID`:     Identification number associated with the Card
- `CARDNUMBER`: Card number associated with the event.
- `CARDHOLDER`: Name of the cardholder associated with the event.
- `USAGECOUNT`: Usage count assigned to a card.

Message Ports cannot be written to.

## Direct DB Access

We try to avoid this as it's both the least secure and least robust, it does give us access to all the information
available to the AxiomXa system though.
