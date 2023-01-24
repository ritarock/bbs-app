import { NextRouter, useRouter } from "next/router";
import { FormEventHandler } from "react";

const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
  event.preventDefault();
  const { value: name } = (event.target as any).name;
  const { value: detail } = (event.target as any).detail;
  fetch("http://localhost:8080/backend/api/v1/topics", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      "name": name,
      "detail": detail,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
    });
};

const clicked = (router: NextRouter) => {
  router.push("http://localhost:3000/");
  router.reload();
};

export const TopicForm = () => {
  const router = useRouter();
  return (
    <>
      <form onSubmit={handleSubmit}>
        <input type="text" name="name" placeholder="topic name" />
        <br />
        <textarea name="detail" placeholder="detail" />
        <br />
        <input type="submit" value="POST" onClick={() => clicked(router)} />
      </form>
    </>
  );
};
