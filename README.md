# Sample with REST API

## Intention

Just playing with the semantics of the APIs. I've been checking the standard and some resources on the internet:
- https://www.ietf.org/rfc/rfc2616.txt.
- https://restfulapi.net/rest-put-vs-post.
- https://gearheart.io/articles/restful-api-design-best-practices.

## Disclaimer

This repo is just for playing and testing, you won't find any test, best practices are applied softly.

## Usage

1. Build it.

```
make
```

2. Run it (if you don't indicate the port it will run on 8080 by default).

```
PORT=<YOUR PORT> ./restapisample
```

3. Curl it!

```
curl -X GET http://localhost:8080/jobs/01G8XJ74WR5FPT8GAWT44ECEGA
curl -X POST http://localhost:8080/jobs  -d '{"title": "the title, "company": "the company"}'
...
```

## Results

<img width="913" alt="image" src="https://user-images.githubusercontent.com/3670816/181053559-ce1fafd7-d16c-4cb9-b8ab-f8eb91c2115c.png">
<img width="1101" alt="image" src="https://user-images.githubusercontent.com/3670816/181053614-86710491-2021-45bc-9985-b6923eeaa17b.png">


## Some thoughts

- My model has not an ID but maybe it should have it. But, it introduces a new challenge, what happens if I want to modify the ID. Should the IDs be modifiable (I believe not!).
- I am using ULIDs as ID. They have interesting properties https://blog.bitsrc.io/ulid-vs-uuid-sortable-random-id-generators-for-javascript-183400ef862c, but they are not human-friendly. Maybe a sort of "sku" should avoid this problem.
- I was trying to use another in-memory key value datastore, but there weren't easy. I love Redis, but it depends on an external software. Maybe there is a better in-memory key value.

## Conclusion

Are REST APIs easy? Probably, but sometimes reviewing the concepts is mandatory! Indeed, I was a bit confused about PUT / POST properties before this excercise.
