import { getCommentByPost } from "@/api/comment";
import { getPost } from "@/api/post";
import { CommentList } from "@/component/CommentList";
import { Post } from "@/interface/post";
import { Comment } from "@/interface/comment";
import { GetServerSideProps } from "next";
import Head from "next/head";
import { CommentForm } from "@/component/CommentForm";
import Link from "next/link";

export default function Posts(
  { post, comments }: { post: Post; comments: Comment[] },
) {
  return (
    <>
      <Head>
        <title>{post.title}</title>
      </Head>
      <div>
        <center>
          {post.title}
          <br />
          {post.content}
          <hr />
          <CommentList comments={comments} />
          <br />
          <CommentForm />

          <Link href={"/"}>
            back
          </Link>
        </center>
      </div>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const post = await getPost(+context.query.id!);
  const comments = await getCommentByPost(+context.query.id!);

  return {
    props: {
      post,
      comments,
    },
  };
};
