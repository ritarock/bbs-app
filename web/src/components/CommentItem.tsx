import type { Comment } from "../generated/model";

interface Props {
  comment: Comment;
}

export function CommentItem({ comment }: Props) {
  const formattedDate = new Date(comment.commentedAt).toLocaleString("ja-JP");

  return (
    <div className="comment-item">
      <div className="comment-item-body">{comment.body}</div>
      <div className="comment-item-date">{formattedDate}</div>
    </div>
  );
}
