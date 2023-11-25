import { useEffect, useState } from "react";
import { CommentAPI } from "../../api";
import { Comment } from "../../interfaces/comment";
import { useParams } from "react-router-dom";
import { useCookies } from "react-cookie";

const CommentList = () => {
  const { id } = useParams();
  const [comments, setComments] = useState<Array<Comment>>([]);
  const [cookie] = useCookies(["token"]);

  useEffect(() => {
    const fetchCommentsList = async () => {
      try {
        const getCommentsList = await CommentAPI.getCommentsAll(
          +id!,
          cookie.token.token,
        );
        setComments([...getCommentsList]);
      } catch (error) {
        console.log("Error fetching comments: ", error);
      }
    };
    fetchCommentsList();
  }, [id, cookie]);

  return (
    <>
      {comments.length > 0 &&
        (
          <ul>
            {comments.map((comment) => (
              <li key={comment.id}>
                {comment.content}
              </li>
            ))}
          </ul>
        )}
    </>
  );
};

export default CommentList;
