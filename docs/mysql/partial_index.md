## partial index 


### Example

```
  Create index idx_non_null_email ON users(email) where email is not null
```

Best practice:

 - match exacly, index only relevant data
 - when many rows is NULL


