# Go API with Gorilla Mux

Specified Root URL: `http://localhost:5000`

Build and Run

```shell
go build && goapi.exe
```

Methods: `GET`, `POST`, `PUT`,`PATCH`, `DELETE`

GET all posts: `/posts`
GET a single post: `/posts/{id}`

Create a post: `/posts` method: `POST`

Request Sample:

```js
{
  "title": "post title",
  "body": "lorem ipsum",
  "author": {
    "fullname": "author",
    "username": "author",
    "email": "author@gmail.com"
  }
}
```

Update All fields: `/posts/{id}` method `PUT`

Request Sample:

```js
{
    "title":"updated post title",
    "body":"updated post body",
    "author":{
        "fullname":"updated author",
        "username":"updated author",,
        "email":"updated author@gmail.com"
    }
}
```

Update a Specific Field: `/posts/{id}/` method `PUT`

```js
{
    "title":"updated post title",
    "author":{
        "fullname":"updated author",
    }
}
```

Delete a Specific Field: `/posts/{id}/` method `DELETE`

> Nasimul Hasan Deep
