import { usePostsRead } from "../generated/api";
import { CommentList } from "./CommentList";

interface Props {
  postId: number | null;
}

export function PostDetail({ postId }: Props) {
  const { data: post, isLoading, error } = usePostsRead(postId ?? 0, {
    query: {
      enabled: postId !== null,
    },
  });

  if (postId === null) {
    return (
      <div className="post-detail">
        <div className="post-detail-empty">Select a post to view details</div>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="post-detail">
        <div className="loading">Loading post...</div>
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="post-detail">
        <div className="error">Failed to load post</div>
      </div>
    );
  }

  const formattedDate = new Date(post.postedAt).toLocaleString("ja-JP");

  return (
    <div className="post-detail">
      <h2 className="post-detail-title">{post.title}</h2>
      <div className="post-detail-date">{formattedDate}</div>
      <p className="post-detail-content">{post.content}</p>

      <CommentList postId={post.id} />
    </div>
  );
}
