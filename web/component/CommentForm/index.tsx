import { postComment } from "../../api/comment";
import { NextRouter, useRouter } from "next/router";
import { FormEventHandler } from "react";

export const CommentForm = () => {
  const router = useRouter();

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input type="text" name="content" placeholder="content" />
        <br />
        <input type="hidden" name="post_id" value={router.query.id} />
        <br />
        <input type="submit" value="POST" onClick={() => clicked(router)} />
      </form>
    </>
  );
};

const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
  event.preventDefault();
  const { value: content } = (event.target as any).content;
  const { value: post_id } = (event.target as any).post_id;

  postComment({ "content": content, "post_id": +post_id });
};

const clicked = (router: NextRouter) => {
  router.push(`http://localhost:3000/${router.asPath}`);
  router.reload();
};
