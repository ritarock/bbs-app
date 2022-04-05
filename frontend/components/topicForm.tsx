import Router, { NextRouter, useRouter } from "next/router";

export default function TopicForm({ postUrl }) {
  const register = async (event) => {
    event.preventDefault();
    const res = await fetch(
      postUrl,
      {
        body: JSON.stringify({
          name: event.target.name.value,
          detail: event.target.detail.value,
        }),
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
      },
    );

    const result = await res.json();
  };

  const router = useRouter();
  return (
    <>
      <form onSubmit={register}>
        <label>name</label>
        <input id="name" type="text" required />
        <label>detail</label>
        <input id="detail" type="text" required />
        <button type="submit" onClick={() => clicked(router)}>
          CREATE
        </button>
      </form>
    </>
  );
}

const clicked = (router: NextRouter) => {
  router.push("http://localhost:3000/");
  router.reload();
};
