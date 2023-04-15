import { type Comment } from "@/interface/comment";

export async function getCommentByPost(id: number): Promise<Comment[]> {
  return fetch(`http://localhost:8080/backend/api/v1/posts/${id}/comments`)
    .then((x) => x.json());
}

export function postComment(body: { "content": string; "post_id": number }) {
  return fetch("http://localhost:8080/backend/api/v1/comments", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  })
    .then((response) => response.json());
}
