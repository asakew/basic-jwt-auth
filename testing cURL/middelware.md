It looks like you have a snippet of Go code from a web application set up using the Gin framework for API testing with cURL. 🎉 I'll help you create cURL commands for the API endpoints you defined! Here’s how you can structure the cURL commands for each of the admin routes:

### 1. Create a Blog
Assuming your endpoint requires some JSON data to create a new blog post, you can use the following cURL command:

```bash
curl -X POST http://localhost:8080/admin/createBlog \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
-d '{
    "title": "My New Blog Admin Post",
    "content": "This is the content of my new blog post."
}'
```

### 2. Edit a Blog
To edit a blog post, you would specify the blog ID in the URL. Here’s an example command:

```bash
curl -X PUT http://localhost:8080/admin/editBlog/123 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
-d '{
    "title": "Updated Blog Post Title",
    "content": "This is the updated content."
}'
```

### 3. Delete a Blog
For deleting a blog post, the command would look like this:

```bash
curl -X DELETE http://localhost:8080/admin/deleteBlog/123 \
-H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

### Notes:
- Replace `http://localhost:8080` with your actual server URL and port.
- Ensure you replace `123` with the actual blog ID you want to edit or delete.
- Replace `YOUR_ADMIN_TOKEN` with a valid token if you're using authorization.

Let me know if you need help with anything else! 😄
