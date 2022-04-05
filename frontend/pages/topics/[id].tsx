import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import Form from "../../components/form";

import { Comment } from "../../interfaces";
import { toDateFormat } from "../../lib/util";

export default function Topic(
  { commentData, topicData }: {
    commentData: { code: number; comments: Comment[] };
    topicData: { code: number; topic: { name: string; detail: string } };
  },
) {
  const router = useRouter();
  return (
    <>
      <h2>{topicData.topic.name}</h2>
      <h3>{topicData.topic.detail}</h3>
      <div>
        {commentData.comments.map((e) => (
          <div key={e.id}>{e.body}: {toDateFormat(e.created_at)}</div>
        ))}
      </div>
      <Form
        postUrl={"http://localhost:8080/backend/api/comments"}
      />
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ params }) => {
  const commentUrl =
    `http://localhost:8080/backend/api/topics/${params.id}/comments`;
  const commentResponse = await fetch(commentUrl);
  const commentData = await commentResponse.json();

  const topicUrl = `http://localhost:8080/backend/api/topics/${params.id}`;
  const topicResponse = await fetch(topicUrl);
  const topicData = await topicResponse.json();

  return {
    props: { commentData, topicData },
  };
};
