import { getPosts } from "@/api/post";
import { PostForm } from "@/component/PostForm";
import { PostList } from "@/component/PostList";
import { Post } from "@/interface/post";
import { GetServerSideProps } from "next";
import Head from "next/head";

export default function Home({ posts }: { posts: Post[] }) {
  return (
    <>
      <Head>
        <title>bbs-app</title>
      </Head>
      <div>
        <center>
          BBS
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
