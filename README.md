# data-struct

#  Scenario

  - mysql: index, real-use case, locking 
  - partner open td:
      client_id, account_id, req_id, amount, status

  - 
    SELECT * from task where done=false
    SELECT * FROM users WHERE PHONE=a OR document_id=b

  - redis: 
    lock, 
    ranking (same score but diff time)

  - pub/sub:
      - retension (7d, exceeded 7d)
      - deadletter queue
      ---
      - ack deadline secs
      - retry policy
      - ordering key
      - pull vs push
      - stuck msg
  - http:
    what happend when click login button
    http call to payment system:
    Err handling


