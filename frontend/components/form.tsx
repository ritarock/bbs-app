import { NextRouter, Router, useRouter } from "next/router";
import CommentForm from "./commentForm";
import TopicForm from "./topicForm";

export default function Form({ postUrl }) {
  return (
    <>
      <FormCheck
        postUrl={postUrl}
      />
    </>
  );
}

function FormCheck({ postUrl }) {
  const router = useRouter();
  if (router.pathname === "/") {
    return (
      <>
        <TopicForm postUrl={postUrl} />
      </>
    );
  }
  if (router.pathname.includes("/topics/")) {
    return (
      <>
        <CommentForm postUrl={postUrl} />
      </>
    );
  }
}
