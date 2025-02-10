```bash
curl -X POST "http://127.0.0.1:3003/login" \
-H "Content-Type: application/json" \
-d '{"email": "test@mail.com", "password": "test12345"}'
```

```bash
curl -X POST "http://127.0.0.1:3003/protected" \
-H "Content-Type: application/json
-H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```


