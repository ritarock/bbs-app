import { GetServerSideProps } from "next";
import { getPosts } from "../api/post";
import Head from "next/head";
import { PostList } from "@/component/PostList";
import { Post } from "../interface/post";
import { PostForm } from "@/component/PostForm";

export default function Home({ posts }: { posts: Post[] }) {
  return (
    <>
      <Head>
        <title>bbs-app</title>
      </Head>
      <div>
        <center>
          <h1>BBS</h1>
          <PostList posts={posts} />
          <br />
          <PostForm />
        </center>
      </div>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  const posts = await getPosts();

  return {
    props: {
      posts,
    },
  };
};
