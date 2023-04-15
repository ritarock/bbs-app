import { Post } from "@/interface/post";
import Link from "next/link";

export const PostList = ({ posts }: { posts: Post[] }) => {
  return (
    <>
      <ul>
        {posts.map((post) => (
          <li key={post.id}>
            <Link href={`/posts/${post.id}`}>
              {post.title}
            </Link>
          </li>
        ))}
      </ul>
    </>
  );
};
