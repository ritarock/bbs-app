import React, { useEffect, useState } from "react";
import { Comment } from "../../@types/comment";
import { CommentAPI } from "../../api";
import Box from "../Box";

interface CommentListProps {
  postID: number;
}

const CommentList: React.FC<CommentListProps> = ({ postID }) => {
  const [comments, setComments] = useState<Comment[]>([]);

  useEffect(() => {
    const fetchComments = async () => {
      try {
        const getComments = await CommentAPI.getCommentAll(postID);
        setComments(getComments);
      } catch {
        console.log("Error fetching posts: ", console.error);
      }
    };

    fetchComments();
  }, [postID]);

  return (
    <>
      {comments.length > 0 && (
        <Box
          content={
            <div>
              {
                <ul>
                  {comments.map((comment) => (
                    <li key={comment.id} className="mb-4">
                      <span className="font-bold text-base pl-1">
                        {comment.content}
                      </span>
                    </li>
                  ))}
                </ul>
              }
            </div>
          }
          offHeader={true}
        />
      )}
    </>
  );
};

export default CommentList;
