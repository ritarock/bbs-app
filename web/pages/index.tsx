import Head from "next/head";
import { TopicList } from "../component/TopicList";
import { GetServerSideProps } from "next";
import { Topic } from "../gen/v1/ts/topic";
import { TopicForm } from "@/component/Form/Topic";

export default function Home({ topics }: { topics: Topic[] }) {
  return (
    <>
      <Head>
        <title>bbs-app</title>
      </Head>
      <TopicList topics={topics} />
      <TopicForm />
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  const url = "http://localhost:8080/backend/api/v1/topics";
  const response = await fetch(url);
  const topics: Promise<Topic> = await response.json();

  return {
    props: { topics },
  };
};
