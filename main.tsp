import "@typespec/http";

using Http;
@service(#{ title: "bbs-app" })
@server("http://localhost:8080", "Single server endpoint")

namespace BBSApp;

model Post {
  id: int64;
  @minLength(1) @maxLength(30) title: string;
  @minLength(1) @maxLength(255) content: string;
  postedAt: utcDateTime;
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

model Comment {
  id: int64;
  postId: int64;
  @minLength(1) body: string;
  commentedAt: utcDateTime;
}

model CommentList {
  items: Comment[];
}

model CreateCommentRequest {
  @minLength(1) body: string;
}

model UpdateCommentRequest {
  @minLength(1) body: string;
}

@route("/posts/{postId}/comments")
interface Comments {
  @post create(@path postId: int64, @body body: CreateCommentRequest): {
    @statusCode statusCode: 201;
    @body body: Comment;
  } | Error;

  @get read(@path postId: int64, @path id: int64): Comment | Error;
  @get list(@path postId: int64): CommentList | Error;

  @put update(@path postId: int64, @path id: int64, @body body: UpdateCommentRequest): Comment | Error;

  @delete delete(@path postId: int64, @path id: int64): {
    @statusCode statusCode: 204;
  } | Error;
}

// === Authentication ===
model User {
  id: int64;
  @minLength(1) @maxLength(255) email: string;
  createdAt: utcDateTime;
}

model AuthResponse {
  user: User;
  token: string;
}

model SignUpRequest {
  @minLength(1) @maxLength(255) email: string;
  @minLength(8) @maxLength(128) password: string;
}

model SignInRequest {
  @minLength(1) @maxLength(255) email: string;
  @minLength(8) @maxLength(128) password: string;
}

@route("/auth")
interface Auth {
  @route("/signup")
  @post signup(@body body: SignUpRequest): {
    @statusCode statusCode: 201;
    @body body: AuthResponse;
  } | Error;

  @route("/signin")
  @post signin(@body body: SignInRequest): AuthResponse | Error;

  @route("/me")
  @get me(@header Authorization: string): User | Error;
}
