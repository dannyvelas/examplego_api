# Abstracting Errors

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
