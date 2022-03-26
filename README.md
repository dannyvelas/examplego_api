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

We don't want packages to expose internal errors because clients can become dependent on those errors. More on this [here](./ERROR-HANDLING.md)

## Some Conventions

### Go conventions this code follows:

* Try for package names to not have underscores or dashes and have meaningful names unlike "util" or "helpers"
* File names with more than one word should be separated by underscores
* Use camelCase for variable names
* Don't automatically make structs only able to be initialized with constructors (aka regular functions in Go).

### Conventions I Decided:

#### Packages
* Make everything private by default
* Try to prevent nesting packages unless necessary.

#### Git Commits
* Commit messages are "\<topic\>: \<what-you-did\>". \<topic\> can be a package, a file, or a general concept like "errors".
    * Examples:
        * `README: added conventions`
        * `models: updated review model`
        * `errors: moved apierror/apierror.go to routing/`
    

#### Errors

* Define all the errors of your package in an `errors.go` file. Here is where you can define your sentinel errors and error constructing functions.
    * Examples: `api/errors.go`, `storage/errors.go`.
* When returning error messages use this format: "file\_name: function\_name: embedded\_error." Use `%v` and not `%w` per [Abstracted Error Handling](#abstracted-error-handling)
    * Example: `err = fmt.Errorf("login_router: Error decoding credentials body: %v", err)`

#### Naming

* Use the plural form for database models. Also use plural form for repos and routers when naming files and variables.
    * Examples:
        * `reviews` table
        * `reviewsRepo`, `reviews_repo`, `ReviewsRepo`
        * `reviewsRouter`, `reviews_router`, `ReviewsRouter`
* Name a variable using lowercase camelCase form of its struct.
    * Example: `adminsRepo := storage.NewAdminsRepo(database)`
* Name a receiver variable using the lowercase camelCase form of its struct.
    * Example: `func (reviewsRepo ReviewsRepo) GetActive...`
* To avoid confusion between when to use "Repo" and when "Repository" I chose to never use the latter. FZF/Grep the code here and you won't find any case-insensitive instances of the word "repository."
