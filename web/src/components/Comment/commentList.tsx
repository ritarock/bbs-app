import { useEffect, useState } from "react";
import { CommentAPI } from "../../api";
import { Comment } from "../../interfaces/comment";
import { useParams } from "react-router-dom";

const CommentList = () => {
  const { id } = useParams();
  const [comments, setComments] = useState<Array<Comment>>([]);

  useEffect(() => {
    const fetchCommentsList = async () => {
      try {
        const getCommentsList = await CommentAPI.getCommentsAll(+id!);
        setComments([...getCommentsList]);
      } catch (error) {
        console.log("Error fetching comments: ", error);
      }
    };
    fetchCommentsList();
  }, [id]);

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
