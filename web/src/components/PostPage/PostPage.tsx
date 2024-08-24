import { useEffect, useState } from "react";
import { Post } from "../../@types/post";
import { PostAPI } from "../../api";
import { useParams } from "react-router-dom";
import Box from "../Box";
import CommentForm from "../CommentForm";
import CommentList from "../CommentList";

const PostPage = () => {
  const { id } = useParams();
  const [post, setPost] = useState<Post>();

  useEffect(() => {
    const fetchPost = async () => {
      try {
        const getPost = await PostAPI.getPost(+id!);
        setPost(getPost);
      } catch {
        console.log("Error fetching post: ", console.error);
      }
    };

    fetchPost();
  }, [id]);

  return (
    <>
      <Box
        content={
          <>
            <div className="font-bold text-lg text-sky-600 pl-1">
              {post?.title}
            </div>
            <div className="pl-1">{post?.content}</div>
          </>
        }
      />
      <br />
      <CommentList postID={+id!} />
      <br />
      <CommentForm postID={+id!} />
    </>
  );
};

export default PostPage;
