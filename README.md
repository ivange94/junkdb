# JunkDB
A key-value store implemented from scratch, inspired by concepts from Designing Data-Intensive Applications.

## Phase One
- In this phase, the focus will be on storage/retrieval implementation so I'll build this as a Go library and not a standalone database server.
### Scope
- Go api to add, update and delete records
### Out of Scope
- Database servers
- Query languages
- Transactions
- Concurrency control

### API
* type JunkDB: instance of the database
* func New(filePath): creates a new database connection. If file doesn't exist, a new one is created.
* *JunkDB.Update(key, value): inserts or updates a key
* *JunkDB.Read(key): gets the value stored at key or 

