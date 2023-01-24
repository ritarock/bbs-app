import { Comment } from "../../gen/v1/ts/comment";

export const CommentList = ({ comments }: { comments: Comment[] }) => {
  return (
    <>
      <div>
        <ul>
          {comments.map((comment) => (
            <li key={comment.id}>
              {comment.body}
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};
