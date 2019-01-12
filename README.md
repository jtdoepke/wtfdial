# WTF Dial

https://medium.com/wtf-dial/

> WTF is a WTF Dial? Essentially it is a real-time API for allowing members of a team to set their current “WTF level” and have it aggregated to a dashboard. It provides insight into the team’s health and allows users to easily provide their input.

## Getting Started

```bash
# Download/install wtfdial
go get github.com/jtdoepke/wtfdial

# Run the wtfdial server
wtfdial server
```

```bash
# Set your WTF level
wtfdial -u myusername 9/10; # I'm pretty confused today.
wtfdial -u myusername 75%; # Feeling a bit better.
wtfdial -u myusername 0.01; # The storm has passed.
```

## Things this doesn't do

- Multiple dials
- Dial history
- Authentication
- Persistence between server restarts
- Have a web dashboard

For some future Lunch'n'Learn I'd love to refactor
this into an example of Domain Driven Design, CQRS, Event Sourcing, and Sagas. (Possible tech stack: Kubernetes, gRPC, NATS Streaming, Kong, Badger. [Similar Example](https://github.com/shijuvar/go-distributed-sys))
