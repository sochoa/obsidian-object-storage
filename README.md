Obsidian
========

Object storage service dataplane

Goals
-----

1.  Provide an object storage service that does not suck.
2.  Provide the building blocks of a web-scale object storage service.
3.  Provide inter-operability with other services -- do one thing well, not the kitchen sink approach.

ToDo
----

* [X] Migrate existing endpoints and web app to use Gorilla Mux in order to use the same URLs for GET/POST/PUT/DELETE methods.
* [ ] Finish CRUD object storage APIs
  - [x] GET
  - [x] POST
  - [ ] PUT -- Not gonna do this for now.  It doesn't seem useful.  
  - [X] DELETE
* [ ] Start UI for object storage control plane
* [ ] Design gossip protocol for node discovery
* [ ] Implement pass-through requests
* [ ] Implement write configuration such that writes will go to N nodes to ensure N replicas
* [ ] Implement nanny service to check integrity of files on disk periodically
* [ ] Implement storage information service to monitor available disk
* [ ] Use PKI for validating data plane requests
* [ ] Use TLS for web endpoint (let's encrypt)
* [ ] Design simulation of 3 endpoints via docker on a single host


Architecture
------------


https://localhost:8080/obsidian/ui

CREATE|GET|DELETE interaction via UI

http://localhost:8080/obsidian/api/{bucket}/{path/to/thing}

CREATE|GET|DELETE api call

