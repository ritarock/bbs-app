import { useCommentsList } from "../generated/api";
import { CommentItem } from "./CommentItem";
import { CommentForm } from "./CommentForm";

interface Props {
  postId: number;
}

export function CommentList({ postId }: Props) {
  const { data, isLoading, error } = useCommentsList(postId);

  return (
    <div className="comment-section">
      <h3 className="comment-section-title">Comments</h3>

      {isLoading && <div className="loading">Loading comments...</div>}

      {error && <div className="error">Failed to load comments</div>}

      <div className="comment-list">
        {data?.items.map((comment) => (
          <CommentItem key={comment.id} comment={comment} />
        ))}

        {data?.items.length === 0 && (
          <div className="empty">No comments yet.</div>
        )}
      </div>

      <CommentForm postId={postId} />
    </div>
  );
}
