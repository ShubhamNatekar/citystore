# Kind of E-Commerce - Example of Microservices in Go with Docker and MongoDB

### Overview

Dream-Project is an example project which demonstrates the use of microservices for a kind of E-Commerce site.
The Dream-Project backend is powered by 4 microservices, all of witch happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

 * Address Service: Provides information like addresses of all user whether they customer of buyer etc.
 * Product Service: Provides all details about all products.
 * Customer : Provides details about all customer. 
 * Shops Service: Provides information about all Shops.

## Deployment

The application can be deployed in both environments: **local machine** or in a **kubernetes cluster**. You can find the appropriate documentation for each case in the following links:

* [localhost](./docs/localhost.md)
* [kubernetes](./docs/kubernetes.md)

## How To Use dream-project Services

* [endpoints](./docs/testing.md)

### Significant Revisions

* [Microservices - Martin Fowler](http://martinfowler.com/articles/microservices.html)
* [Web Development with Go](http://www.apress.com/9781484210536)
* [Umer Mansoor - Cinema](https://github.com/umermansoor/microservices)
* [Traefik](https://traefik.io/)
