# Go Scaffold Repositories

This package provides a set of conventions and default implementations for interacting with database instances and tables in a Go application. Each database instance is represented by `***DB`, and tables are represented by `***Repository`.

## Conventions

- **Database Instances**: Represented by `***DB`.
- **Tables**: Represented by `***Repository`.

## Repository Methods

Each repository includes the following default method signatures:

- **Create**: Adds a new record to the table.
- **Get**: Retrieves records based on specified criteria.
- **FindByID**: Retrieves a record by its unique identifier.
- **Update**: Updates a record with all mandatory fields. Non-supplied fields will be set to `null`.
- **Patch**: Updates a record with only the supplied fields.
- **DeleteByID**: Deletes a record by its unique identifier.

### Method Details

#### Create

```go
func (r *YourRepository) Create(entity YourEntity) error
```

#### Get

```go
func (r *YourRepository) Get(criteria YourCriteria) ([]YourEntity, error)
```

#### FindByID

```go
func (r *YourRepository) FindByID(id string) (YourEntity, error)
```

#### Update

```go
func (r *YourRepository) Update(entity YourEntity) error
```

#### Patch

```go
func (r *YourRepository) Patch(id string, updates map[string]interface{}) error
```

#### DeleteByID

```go
func (r *YourRepository) DeleteByID(id string) error
```

## Example Usage

```go
type UserDB struct {
    // Database connection and other fields
}

type UserRepository struct {
    db *UserDB
}

type User struct {
    ID    string
    Name  string
    Email string
}

// Implement the repository methods for UserRepository
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

For any questions or inquiries, please contact [i.am@ferrysutanto.com].
