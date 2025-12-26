import "@typespec/http";

using Http;
@service(#{ title: "bbs-app" })
@server("http://localhost:8080", "Single server endpoint")

namespace BBSApp;

model Post {
  id: int64;
  @minLength(1) @maxLength(30) title: string;
  @minLength(1) @maxLength(255) content: string;
  createdAt: utcDateTime;
}

model PostList {
  items: Post[]
}

model CreatePostRequest {
  @minLength(1) @maxLength(30) title: string;
  @minLength(1) @maxLength(255) content: string;
}

model UpdatePostRequest {
  @minLength(1) @maxLength(30) title: string;
  @minLength(1) @maxLength(255) content: string;
}

@error
model Error {
  code: int32;
  message: string;
}

@route("/posts")
interface Posts {
  @post create(@body body: CreatePostRequest): {
    @statusCode statusCode: 201;
    @body body: Post;
  } | Error;

  @get read(@path id: int64): Post | Error;
  @get list(): PostList | Error;

  @put update(@path id: int64, @body body: UpdatePostRequest): Post | Error;

  @delete delete(@path id: int64): {
    @statusCode statusCode: 204;
  } | Error;
}
