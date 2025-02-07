# Curl to proto.Any
```sh
curl --location --request POST 'http://localhost:9000/hack/posting-batch' \
--header 'Content-Type: application/json' \
--data-raw '{
    "instructions": [
        {
            "instruction": {
                "@type": "type.googleapis.com/main.main.mint.ThisIsMyMsg",
                "value": {

                    "amount": 1,
                    "idempotent_key": "2892c441-4bb0-48d9-9226-58078284c126"
                }
            },
            "metadata": "eyJrZXkiOiJ2YWwifQo=",
            "restrictions": 324234
        }
    ]
}'
```
