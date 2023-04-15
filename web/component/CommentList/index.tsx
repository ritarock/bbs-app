import { Comment } from "@/interface/comment";

export const CommentList = ({ comments }: { comments: Comment[] }) => {
  return (
    <>
      <ul>
        {comments.map((comment) => (
          <li key={comment.id}>
            {comment.content}
          </li>
        ))}
      </ul>
    </>
  );
};
