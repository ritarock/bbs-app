import { postPost } from "../../api/post";
import { NextRouter, useRouter } from "next/router";
import { FormEventHandler } from "react";

export const PostForm = () => {
  const router = useRouter();

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input type="text" name="title" placeholder="title" />
        <br />
        <input type="text" name="content" placeholder="content" />
        <br />
        <input type="submit" value="POST" onClick={() => clicked(router)} />
      </form>
    </>
  );
};

const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
  event.preventDefault();
  const { value: title } = (event.target as any).title;
  const { value: content } = (event.target as any).content;

  postPost({ "title": title, "content": content });
};

const clicked = (router: NextRouter) => {
  router.push("http://localhost:3000/");
  router.reload();
};
