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

### Error Handling

Today, in the `storage` package, I'm using the `database/sql` package to interact with my database. This library will return an error, if for example there is some connectivity problem, user request problem, or development mistake.

I don't want the `storage` package to expose any of the errors that the `database/sql` library returns. But why?

Suppose:
1) The `storage` packages exposes the `sql.ErrNoRows` error that comes from the `database/sql` package. It returns this error to indicate that no rows were found.
2) My `api` package imports my `storage` package, and asks it for some data
3) My `api` package realizes that whenever my `storage` package can't find rows, it returns a `sql.ErrNoRows` error. So, whenever it asks for data, it performs a check: "do this if the `storage` package returns `sql.ErrNoRows`."
4) I one day decide to replace `database/sql` with a different package that returns a `diffpkg.ErrNoRows` error when there are no rows, instead of a `sql.ErrNoRows` error.

In this case, my `api` code would break. The check that it performs of "do this if `sql.ErrNoRows` is returned" would never execute, even if storage actually couldn't find any rows.

This case isn't so problematic because I'm the developer of the `storage` package. And, I'm also the developer of the `storage` package's client package (the `api` package). So, I could simply go in to the `api` package code and fix it. I'll make it check for `diffpkg.ErrNoRows` instead of `sql.ErrNoRows`.

But, what if I shared my `storage` package code? What if other people around the world started importing `storage`? In this case, once I made that change to the `storage` package, people all around the world that are checking for `sql.ErrNoRows`, would have their code break.

So, how could we avoid this problem?

The answer is never expose `sql.ErrNoRows` to begin with. Create a new error in your `storage` package that will never ever change.

In this case, we could create a `storage.ErrNoRows` error. This error will be returned regardless of whether the `storage` package uses `database/sql` or `diffpkg`. So, the `api` code won't check for `sql.ErrNoRows`, it will check for `storage.ErrNoRows`. And, it will be future-proof safe against all internal changes in the `storage` code.

Obviously it's pretty unrealistic I would publish the `storage` package. But, the point still stands: hide internal errors so that your clients don't break when they change.

Even if you don't publish your package, hiding internal errors will help you do less work when you change the internals of dependencies.
