# Example Go API

I built this repo to show patterns and conventions that I think are nice and helpful when building an API. I chose to use Go because it's a sweet, well-thought-out language, that is well-suited for HTTP applications.

The database models this repo uses are totally unimportant. They just exist for testing and demonstration purposes.

## Some patterns:

### Separation of concern

The database, api, and config logic are all in distinct packages. This allows for really nice abstraction.

For example, the api package can ask the database package for some data, without knowing at all what it does or uses internally. It won't know what the database query looks like, what database library is being used, or what errors that library might return.

### Dependency Injection

Problem: Many packages depend on other packages to run.

Example: Let's suppose you want your API to return a list of reviews.

Well, for that, the `api` package needs a `reviewsRepo` from the `storage` package. The `reviewsRepo` needs a `Database`. A `Database` needs a `Config`.

The question is, how do we give these packages what they need?

Answer: Initialize all the dependencies in the entry point function, `main`. Start with the most basic dependency, a `Config` object. Once you a `Config` object, you can create a `Database`. Once you create a `Database`, you can create a `reviewsRepo`. Once you create a `reviewsRepo`, you can give that `reviewsRepo` to the `api.reviewsRouter`. The `api.reviewsRouter` can then use `reviewsRepo` to get a list of reviews.

### Abstracted Error Handling

We don't want packages to expose internal errors because clients can become dependent on those errors. More on this [here](./ERROR-HANDLING.md).

### Conventions

Are [here](./CONVENTIONS.md).

### To-Dos

Are [here](./TODO.md)

### Shout Outs

#### Simple Gopher
The biggest and most helpful reference in building this project was this repo I found on Reddit: <https://github.com/doppelganger113/simple_gopher>. I've learned a lot from it.

It uses the same patterns of separation of concern and dependency injection. However, it is a little bit more complex.

Some Differences:
* My code will have a dependency chain like: `api->repo->database`. Marko's code looks like: `api->service->repo->database`.
* Each layer in his code is separated by interfaces. I use structs.
* He has fancy concurrency, CICD, and AWS Cognito Authentication stuff. I don't have these things yet. And, may not add them.
* I think he doesn't abstract errors between layers.

#### Useful links
* <https://benhoyt.com/writings/go-routing/>
* <https://stackoverflow.com/questions/26462043/how-to-disallow-direct-struct-initialization>
* <https://www.alexedwards.net/blog/organising-database-access>
* <https://www.alexedwards.net/blog/making-and-using-middleware>
* <https://go.dev/blog/go1.13-errors>
* <https://www.joeshaw.org/error-handling-in-go-http-applications/>
