# Go and PostgreSQL

a simple crud with go and postgre

## Routes

- http://127.0.0.1:9999 || (homepage)
- http://127.0.0.1:9999/api/articles || (method POST: add article)
- http://127.0.0.1:9999/api/articles || (method GET: get all articles)
- http://127.0.0.1:9999/api/articles/{id} || (method GET: get detail of article)
- http://127.0.0.1:9999/api/articles/{id} || (method PUT: update article)
- http://127.0.0.1:9999/api/articles/{id} || (method DELETE: delete article)

## Data JSON

```javascript
{
     "id": 1,
     "code": "ART123",
     "title": "belajar pemrograman golang",
     "desc": "belajar golang, gorm, dan postgre dengan postman",
     "content": "test123 test124"
}
```
