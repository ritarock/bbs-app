import Head from "next/head";
import { GetServerSideProps } from "next";
import { Topic } from "../../gen/v1/ts/topic";
import { Comment } from "../../gen/v1/ts/comment";
import { CommentList } from "@/component/CommentList";
import { CommentForm } from "@/component/Form/Comment";

export default function TopicView(
  { topic, comments }: { topic: Topic; comments: Comment[] },
) {
  return (
    <>
      <Head>
        <title>bbs-app</title>
      </Head>
      <h2>{topic.name}</h2>
      <div>{topic.detail}</div>
      <CommentList comments={comments} />
      <CommentForm />
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ params }) => {
  const topicUrl = `http://localhost:8080/backend/api/v1/topics/${params!.id}`;
  const responseTopic = await fetch(topicUrl);
  const topic: Promise<Topic> = await responseTopic.json();

  const commentUrl = `http://localhost:8080/backend/api/v1/topics/${
    params!.id
  }/comments`;
  const responseComments = await fetch(commentUrl);
  const comments: Promise<Comment> = await responseComments.json();

  return {
    props: { topic, comments },
  };
};
