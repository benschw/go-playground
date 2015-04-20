
### ugly

- shows straight forward way to discover an address
- can't test `Discover` without actually resolving a SRV record

### better

- adds `AddressGetter` interface so that tests can use a stub implementation
- didn't have to include unnecessary interface in core library
- allows implementation to only depend on the parts of the library we are using


### server

- example client/server implementation where client discovers server's address
- out `AddressGetter` interface allows load balancer to be injected, 
  so when component testing the local instance can be leveraged.

