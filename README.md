# JunkDB
A key-value store implemented from scratch, inspired by concepts from Designing Data-Intensive Applications.

## Phase One
- In this phase, the focus is an append-only log for storage and retrieval.
### Scope
- Append new records to the log
- Update a key by appending a newer value for the same key
- Read the latest value for a key by scanning the log from the end
### Out of Scope
- Delete support
- Query languages
- Transactions
- Concurrency control

### API
- `put <key> <value>` appends a record to the log
- `get <key>` returns the newest value for that key
