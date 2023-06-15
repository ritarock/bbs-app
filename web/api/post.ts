import { Post } from "../interface/post";

export async function getPosts(): Promise<Post[]> {
  return fetch("http://localhost:8080/backend/api/v1/posts").then((response) =>
    response.json()
  );
}

export async function getPost(id: number): Promise<Post> {
  return fetch(`http://localhost:8080/backend/api/v1/posts/${id}`).then(
    (response) => response.json(),
  );
}

export function postPost(body: { "title": string; "content": string }) {
  return fetch("http://localhost:8080/backend/api/v1/posts", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  })
    .then((response) => response.json());
}
