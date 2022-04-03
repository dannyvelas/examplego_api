# Conventions

These aren't hard and fast rules. They are guidelines to remove overhead when writing/sharing code. It would be good to follow them when possible. These are followed in addition to the [Go Conventions Here](https://github.com/golang/go/wiki/CodeReviewComments).

## Packages
* Make everything private by default
* Try to prevent nesting packages unless necessary.

## Git Commits
* Commit messages are "\<topic\>: \<what-you-did\>". \<topic\> can be a package, a file, or a general concept like "errors". The only exception are commits that concern the entire repository.
    * Examples:
        * `README: added conventions`
        * `models: updated review model`
        * `errors: moved apierror/apierror.go to routing/`
        * `upgraded to Go 1.18`

## Errors

* Define all the errors of your package in an `errors.go` file. Here is where you can define your sentinel errors and error constructing functions.
    * Examples: `api/errors.go`, `storage/errors.go`.
* When returning error messages use this format: "file\_name: function\_name: embedded\_error." Use `%v` and not `%w` per [Abstracted Error Handling](#abstracted-error-handling)
    * Example: `err = fmt.Errorf("login_router: Error decoding credentials body: %v", err)`
* Sentinel Errors are prefixed with `Err` or `err`:
    * Examples:
        * `errUnauthorized = sentinelError{http.StatusUnauthorized, "Unauthorized"}`
        * `ErrDatabaseQuery = sentinelError{"Error querying database"}`

## Naming

* Use the plural form for database models. Also use plural form for repos and routers when naming files and variables.
    * Examples:
        * `reviews` table
        * `reviewsRepo`, `reviews_repo`, `ReviewsRepo`
        * `reviewsRouter`, `reviews_router`, `ReviewsRouter`
* Try to avoid abbreviations most of the time.
    * Examples:
        * `defaultValue` not `defaultVal`
        * `database` not `db`
* When possible, name a variable using lowercase camelCase form of its struct. If it's an error, you can name the variable `err` or suffix the variable with `Err`.
    * Examples:
        * `adminsRepo := storage.NewAdminsRepo(database)`
        * `func respondError(w http.ResponseWriter, internalErr error, apiErr apiError) {`
* When possible, name a receiver variable using the lowercase camelCase form of its struct. If it's an error struct, use `e` instead.
    * Examples:
        * `func (reviewsRepo ReviewsRepo) GetActive...`
        * `func (e sentinelError) Error()...`
* To avoid confusion between when to use "Repo" and when "Repository" I chose to never use the latter. FZF/Grep the code here and you won't find any case-insensitive instances of the word "repository."
* Name `.md` files using `TRAIN-CASE`, or my favorite synonym, `SCREAMING-KEBAB-CASE`.
