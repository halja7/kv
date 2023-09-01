# Key Value store


1. Make a key value store
2. Make it persist
3. (future) There may be some tasks that execute

## How do we persist this data

1. Probably initially persist to a file (maybe postgres irl)
2. The store shouldn't have to know that it's connected to a disk file


Read operations

0. On open, read from file and hydrate in-memory map `OpenStore(filename string)`
    - If not there, it's a new store
1. First, attempt to read from in-memory map
    1. If not there, read from file
    2. Else, return value from map

Write operations

0. Need a persist to disk operation (Flush) (user shouldn't have to call it)
1. When?
    - Not when open or read
    - Only `Set` can drift the in-memory db and file
    - Set always calls `Flush`, leaves the store always valid


## Next steps

Writing a test for OpenStore
- Asserting against values in a file 
- Could just be `testdata.json` already on the fs

Hang on... we can use Flush for this!

Set
Flush
OpenStore from flushed file
Get
Assert


Can also write more granular individual tests after e2e ...
(covers what we care about)

## Let's do a CLI (2 weeks)

1. How could users use this package?
2. What's the simplest CLI around this store that would be of any use? 


