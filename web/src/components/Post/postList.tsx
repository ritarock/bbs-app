import { useEffect, useState } from "react";
import { PostAPI } from "../../api";
import { Post } from "../../interfaces/post";
import PostForm from "./form";

const PostList = () => {
  const [posts, setPosts] = useState<Array<Post>>([]);
  useEffect(() => {
    const fetchPostsList = async () => {
      try {
        const getPostsList = await PostAPI.getPostsAll();
        setPosts([...getPostsList]);
      } catch (error) {
        console.log("Error fetching posts: ", error);
      }
    };
    fetchPostsList();
  }, []);

  return (
    <>
      <div>
        {posts.length > 0 &&
          (
            <ul>
              {posts.map((post) => (
                <li key={post.id}>
                  <a href={`/posts/${post.id}`}>
                    {post.title}
                  </a>
                </li>
              ))}
            </ul>
          )}
      </div>
      <hr />
      <div>
        <PostForm />
      </div>
    </>
  );
};

export default PostList;
