import { NextRouter, useRouter } from "next/router";
import { FormEventHandler } from "react";

const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
  event.preventDefault();
  const { value: body } = (event.target as any).body;
  const { value: topic_id } = (event.target as any).topic_id;
  fetch("http://localhost:8080/backend/api/v1/comments", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      "body": body,
      "topic_id": +topic_id,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
    });
};

const clicked = (router: NextRouter) => {
  router.push(`http://localhost:3000/${router.asPath}`);
  router.reload();
};

export const CommentForm = () => {
  const router = useRouter();
  return (
    <>
      <form onSubmit={handleSubmit}>
        <textarea name="body" placeholder="body" />
        <input name="topic_id" type="hidden" value={router.query.id} />
        <br />
        <input type="submit" value="POST" onClick={() => clicked(router)} />
      </form>
    </>
  );
};
