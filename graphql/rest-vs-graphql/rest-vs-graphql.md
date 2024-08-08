# REST vs GraphQL

GraphQL is new but REST has been there for a while, let’s compare these to see which one is better. While both REST and GraphQL have their strengths and weaknesses, in this blog we will cover how GraphQL solves the problem of over and under-fetching of data and considerations for choosing the right approach for your project.

## How REST works?

![how-rest-works](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/how-rest-works.png)

The above image is a typical example of how RESTful applications work. As you can see there is a REST client and a backend server which consists of multiple services. So the client makes multiple API calls to multiple services in the backend server.
Some of the common issues a developer faces, while designing a RESTful architecture are:

* Underfetching
* Over-fetching
* Making multiple API calls
  
Let’s see this issue with the help of an example.
Below is an example of a typical blog application where people can write some posts, read a post written by other people, and may or may not like it.

![news-feed](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/news-feed.png)

As you can see in the above image, there is a newsfeed screen with a list of blog posts, where each post requires a title, a body, date, author name, and number of likes. So to fill in this data we can make an API call to __/posts__ that lists all the blog posts.
__/posts__ gave us a list of blogs. But as you can see in response we are getting the author ID but on screen, we have to show the author's name. So this is called under-feching. So to fetch an author's details we can make an API call __/author/:id__. And for each blog, we have to make another API call to fetch author details. So this is a typical example where to fill in data we would end up making N+1 queries as shown in the below image.

![news-feed-2](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/news-feed2.png)

## A Possible Solution

![news-feed-issues](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/news-feed-issues.png)

We can design a custom endpoint __/newsfeed__ or __/dashboard__ that gives us the exact fields needed for this page. Designing a new endpoint specific to our screen can solve our problem in this case. But what if we update the UI component in the future? In that case either we will have some extra fields or some lesser fields in this custom endpoint, which would result in updating the endpoint or versioning the endpoint. And this is where complexities begin.

## How does GraphQL work?

In GraphQL, we have exact or controlled fetching where data aggregates into a single query.

![how-graphql-works](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/how-graphql-works.png)

This is what a typical GraphQL query looks like.

![gql-query](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/gql-query.png)

As you can see we get what we asked for. The request and response have the same shape and the exact fields that the client asked for.
So REST is an architectural style, which gives the current state of a resource on the server. REST is centered around resources in the server where the server decides what data should come in response. While GraphQL is just a query language for reading and mutating data. So in GraphQL, we don’t specify how the data is retrieved but the kinds of data that are available. Here is a tabular comparison of both:

![rest-vs-grahql](https://github.com/goschool-dev/blog/blob/master/graphql/rest-vs-graphql/assets/rest-vs-graphql.png)

## Drawbacks of GraphQL

* The beauty of REST is that we don’t need any special libraries to consume someone else’s APIs. Requests can simply be sent using common tools like cURL or simply by a web browser. But for graphQL, we need special packages and libraries like Apollo & Gqlgen for Go to write the backend server and make client calls.
* GraphQL queries are difficult to cache. REST uses HTTP GET method for fetching resources and HTTP GET has a well-defined caching behavior that is leveraged by browsers, CDNs, proxies, etc. GraphQL has a single point of entry and it uses HTTP POST by default, which prevents full use of HTTP caching.
  
## Some use cases of GraphQL

* We can use GraphQL for apps for devices such as mobile phones, smartwatches, and IoT devices where bandwidth usage matters. GraphQL can be very helpful here as it will improve API performance by reducing the number of API calls and payload size by querying only required fields.
* Where we need to fetch nested and related resources by a single API call, as graphQL can aggregate multiple resources into a single query.
* We can make use of a graphQL gateway, which would be a layer of graphQL over the existing REST backend. So you won’t have to change the existing legacy software.

## Conclusion

Choosing between REST and GraphQL depends on various factors such as the application's complexity and the client's needs. REST with its simplicity, scalability, and wide adoption remains a solid choice for simple CRUD-based APIs. At the same time, GraphQL with its strong type system, schema flexibility, and reduced network traffic,  shines in scenarios where flexibility, efficiency, and precise data fetching are important.
