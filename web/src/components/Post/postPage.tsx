import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { PostAPI } from "../../api";
import { Post } from "../../interfaces/post";
import CommentForm from "../Comment/form";
import CommentList from "../Comment/commentList";
import { useCookies } from "react-cookie";

const PostPage = () => {
  const { id } = useParams();
  const [post, setPost] = useState<Post>();
  const [cookie] = useCookies(["token"]);

  useEffect(() => {
    const fetchPost = async () => {
      try {
        const getPost = await PostAPI.getPost(+id!, cookie.token.token);
        setPost(getPost);
      } catch (error) {
        console.log("Error fetching post: ", error);
      }
    };
    fetchPost();
  }, [id, cookie]);

  return (
    <>
      <h1>{post?.title}</h1>
      <h2>{post?.content}</h2>
      <hr />
      <CommentList />
      <hr />
      <CommentForm />
      <a href="/">TOP</a>
    </>
  );
};

export default PostPage;
