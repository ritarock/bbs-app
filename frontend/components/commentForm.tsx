import { NextRouter, useRouter } from "next/router";

export default function TopicForm({ postUrl }) {
  const router = useRouter();
  const register = async (event) => {
    event.preventDefault();
    const res = await fetch(
      postUrl,
      {
        body: JSON.stringify({
          body: event.target.body.value,
          topic_id: +router.query.id,
        }),
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
      },
    );

    const result = await res.json();
  };

  return (
    <>
      <form onSubmit={register}>
        <label>comment</label>
        <input id="body" type="text" required />
        <button type="submit" onClick={() => clicked(router)}>
          CREATE
        </button>
      </form>
    </>
  );
}

const clicked = (router: NextRouter) => {
  router.push(`http://localhost:3000/topics/${router.query.id}`);
  router.reload();
};
