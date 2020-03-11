Obsidian
========

Object storage service dataplane

Goals
-----

1.  Provide reliable, durable, performant, and secure object storage service.
2.  Provide the building blocks of a web-scale object storage service.
3.  Provide inter-operability with other services -- do one thing well, not the kitchen sink approach.

ToDo
----

* Finish CRUD object storage APIs
    [x] GET
    [x] POST
    [] PUT
    [] DELETE
* Start UI for object storage control plane
* Design gossip protocol for node discovery
* Implement pass-through requests
* Implement write configuration such that writes will go to N nodes to ensure N replicas
* Implement nanny service to check integrity of files on disk periodically
* Implement storage information service to monitor available disk
* Use PKI for validating data plane requests
* Use TLS for web endpoint (let's encrypt)
* Design simulation of 3 endpoints via docker on a single host
