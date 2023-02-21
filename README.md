# go-chi-blog-api
Go Chi api for simple blog application

1. Get all posts:
curl http://localhost:8080/posts

2. Get a post by ID:
curl http://localhost:8080/posts/1

3. Create a new post:
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Post","content":"This is a new post"}' http://localhost:8080/posts

4. Update an existing post by ID:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Post","content":"This is an updated post"}' http://localhost:8080/posts/1

5. Delete a post by ID:
curl -X DELETE http://localhost:8080/posts/1

6. Like a post by ID:
curl -X POST http://localhost:8080/posts/1/like

7. Comment on a post by ID:
curl -X POST -H "Content-Type: application/json" -d '{"content":"This is a new comment"}' http://localhost:8080/posts/1/comment
