import { useEffect, useState } from "react";
import { PostAPI } from "../../api";
import { Post } from "../../interfaces/post";
import PostForm from "./form";
import { useCookies } from "react-cookie";

const PostList = () => {
  const [posts, setPosts] = useState<Array<Post>>([]);
  const [cookie] = useCookies(["token"]);
  useEffect(() => {
    const fetchPostsList = async () => {
      try {
        const getPostsList = await PostAPI.getPostsAll(cookie.token.token);
        setPosts([...getPostsList]);
      } catch (error) {
        console.log("Error fetching posts: ", error);
      }
    };
    fetchPostsList();
  }, [cookie]);

  return (
    <>
      {Object.keys(cookie).length === 0
        ? (
          <>
            <p>
              <a href="/login">login</a>
            </p>
            <p>
              <a href="/signup">signup</a>
            </p>
          </>
        )
        : (
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
        )}
    </>
  );
};

export default PostList;
