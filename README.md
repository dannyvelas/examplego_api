# Example Go API

I built this repo to show patterns and conventions that I think are nice and helpful when building an production quality API. I chose to use Go because it's a sweet, well-thought-out language, that is well-suited for HTTP applications.

The database models this repo uses are totally unimportant. They just exist for testing and demonstration purposes.

I explain why I made this [here](./MOTIVATION.md).

**NOTE**: This is WORK IN PROGRESS. I plan to add and fix some endpoints. Here are some of the [TODOs](./TODO.md).

## Setup

1. `git clone https://github.com/dannyvelas/examplego_api.git && cd examplego_api`
2. `cp .env.example .env`
3. `docker-compose up -d`
4. `make run`

## Some patterns:

<details>

<summary>Slimness Within Reason</summary>

I tried to minimize the amount of size of dependencies, within reason. The most important dependency here is the routing library. This required the most thought and research.

Per my motivation, I chose not to use [Gin](https://github.com/gin-gonic/gin) even though it is probably Golang's most famous HTTP routing dependency. It seemed like it provided more features than I needed.

I could have gone to the extreme and only used `net/http` for routing, using something like Axel Wagner’s [Shift Path technique](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html). But, I felt like this was too much boilerplate.

So, I opted for [go-chi](https://github.com/go-chi/chi). This felt like a happy medium. It's routing logic is quite small (claiming ~1000LOC), yet it's still very functional and easy to use. As a bonus, it's perfect for modularity (more on that in the next section) and fast.

I was planning to use [http-router](https://github.com/julienschmidt/httprouter) because I think it's [even faster](https://gist.github.com/pkieltyka/123032f12052520aaccab752bd3e78cc) and similarly light. But I didn't because it [doesn't have support for subrouters](https://github.com/julienschmidt/httprouter/issues/141). So, it's a little bit harder to achieve modularity.
</details>

<details>

<summary>Separation of Concern</summary>

I tried to separate concerns as much as possible, keeping everything in its own isolated module.

For example, the database, API, and config logic are all in distinct packages. This means that the `api` package can ask the database package for some data, without knowing at all what it does or uses internally. It won't know what the database query looks like, what database library is being used, or what errors that library might return.

Also, I exposed some routes in the `main` file, like `/api/login` and `/api/admin/reviews`. But I chose to keep domain-specific routes in their own sub-routes. For example `/api/admin/reviews/all` and `/api/admin/reviews/active` are only listed and defined in a sub-router which is in `api/reviews_router.go`.
</details>

<details>

<summary>Dependency Injection</summary>

As I was writing this, I noticed that I needed some way of making my `Database` accessible to my routers. When I was first learning how to make API endpoints, I realized that an easy way to do this was to just make a globally scoped singleton instance of a `Database`. 

I think this works fine in NodeJS because [JS is not a multithreading language](https://deepu.tech/concurrency-in-modern-languages-js/). So, singletons in NodeJS need not be thread-safe. However in every other language, singletons are probably best to avoid if you don't want to touch [thread synchronization](https://stackoverflow.com/questions/1823286/singleton-in-go).

Aside from being unsafe, singletons also seem to be an [overused pattern in general](https://gameprogrammingpatterns.com/singleton.html).

Steering away from singletons, I came across [dependency injection](https://www.alexedwards.net/blog/organising-database-access). This was perfect! I could inject a service that interacts with the database into my routing functions.

As an example, suppose I want a routing function to get some reviews from the database. How can I do this?

In `main`, I could initialize an instance of a `Database` and pass or "inject" that into the `reviewsRepo` service. I can then inject the `reviewsRepo` service into `api.reviewsRouter`. Consequently, all the routing functions in `api.reviewsRouter` will have access to `reviewsRepo`, which will have access to the database.
</details>

<details>

<summary>Abstracted Error Handling</summary>

I'm very careful and interested in error handling. In my opinion it's a majorly important thing that often gets glossed over or put off. It's very obvious that programs generally get an input A and turn it to output B. But, it's more subtle to realize that they actually also may return a variety of other failure outputs. 

The path the program takes to returns B and not any failure output, is often called the happy path. And, the paths that return non-B outputs are called unhappy paths.[^1]

Unhappy paths are more subtle because developers are often thinking about how to get their program to return the right output. So, the happy path is where most of the focus and energy goes. The unhappy paths are often just treated as "throw an exception here. And, if you have time, make sure its error message doesn't expose internal or sensitive information."

However, after some years of using monadic functional types in Scala, Elm, and Rust, I've realized just how many unhappy paths there are. These languages had forced me to use types like `Maybe` and `Result<Left, Right>`, where `None` or `Left` represent unhappy results, and `Just` or `Right` represent happy results. Seeing these types all over my programs made me realize that error handling may be close to half of where development is spent, even though its where only a fraction of focus goes.

So, I tried my best to set up a good convention in handling errors here, taking advantage of Go's [explicit error handling approach](https://go.dev/blog/error-handling-and-go) and some of its [neat ways](https://go.dev/blog/go1.13-errors) to embed errors.

Part of this convention is to abstract errors between packages. I go into even more depth [here](./ABSTRACTING-ERRORS.md).
</details>

## Conventions

Are [here](./CONVENTIONS.md).

## Shout Outs

<details>

<summary>Simple Gopher</summary>

The biggest and most helpful reference in building this project was this repo I found on Reddit: <https://github.com/doppelganger113/simple_gopher>. I've learned a lot from it.

It uses the same patterns of separation of concern and dependency injection. However, it is a little bit more complex.

Some Differences:
* My code will have a dependency chain like: `api->repo->database`. Marko's code looks like: `api->service->repo->database`.
* Each layer in his code is separated by interfaces. I use structs.
* He has fancy concurrency, CICD, and AWS Cognito Authentication stuff. I don't have these things yet. And, may not add them.
* I think he doesn't abstract errors between packages.
</details>

<details>

<summary><h3>Additional Useful Links</h3></summary>

* Deciding what router to use: <https://benhoyt.com/writings/go-routing/>
* Deciding whether to use a framework or library: <https://stephensearles.com/framework-vs-library/>
* Whether to use getter/setter and constructor pattern in Golang: <https://stackoverflow.com/questions/26462043/how-to-disallow-direct-struct-initialization>
* Middleware patterns: <https://www.alexedwards.net/blog/making-and-using-middleware>
* Testing database: <https://faun.pub/how-to-test-database-repository-in-golang-771b59c8084e>
* Testify suite: <https://medium.com/nerd-for-tech/testing-rest-api-in-go-with-testify-and-mockery-c31ea2cc88f9>
* Testing naming conventions: <https://ieftimov.com/posts/testing-in-go-naming-conventions/#:~:text=The%20Golang%20source%20code%20itself,of%20the%20function%20under%20test.>
</details>


[^1]: I like to think that this is like the [Anna Karenina principle in statistics](https://en.wikipedia.org/wiki/Anna_Karenina_principle). A dataset may violate the null hypothesis in various ways, but there's only one way in which all the assumptions are satisfied. Similarly, a program may fail in various ways, but succeed in only one way.
