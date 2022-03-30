# Conventions

These aren't hard and fast rules. They are agreed-upon guidelines to remove overhead when writing/sharing code. It would be good to follow them when possible.

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
* Sentinel Errors are prefixed with `Err` or `err`:
    * Examples:
        * `errUnauthorized = sentinelError{http.StatusUnauthorized, "Unauthorized"}`
        * `ErrDatabaseQuery = sentinelError{"Error querying database"}`

#### Naming

* Use the plural form for database models. Also use plural form for repos and routers when naming files and variables.
    * Examples:
        * `reviews` table
        * `reviewsRepo`, `reviews_repo`, `ReviewsRepo`
        * `reviewsRouter`, `reviews_router`, `ReviewsRouter`
* When possible, name a variable using lowercase camelCase form of its struct. If it's an error, you can name the variable `err` or suffix the variable with `Err`.
    * Example: `adminsRepo := storage.NewAdminsRepo(database)`
    * Example: `func respondError(w http.ResponseWriter, internalErr error, apiErr apiError) {`
* When possible, name a receiver variable using the lowercase camelCase form of its struct. If it's an error struct, use `e` instead.
    * Example: `func (reviewsRepo ReviewsRepo) GetActive...`
    * Example: `func (e sentinelError) Error()...`
* To avoid confusion between when to use "Repo" and when "Repository" I chose to never use the latter. FZF/Grep the code here and you won't find any case-insensitive instances of the word "repository."
* Name `.md` files using `TRAIN-CASE`, or my favorite synonym, `SCREAMING-KEBAB-CASE`.
