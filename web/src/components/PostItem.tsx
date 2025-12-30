import type { Post } from "../generated/model";

interface Props {
  post: Post;
  isSelected: boolean;
  onClick: () => void;
}

export function PostItem({ post, isSelected, onClick }: Props) {
  const formattedDate = new Date(post.postedAt).toLocaleString("ja-JP");

  return (
    <div
      className={`post-item ${isSelected ? "selected" : ""}`}
      onClick={onClick}
    >
      <div className="post-item-title">{post.title}</div>
      <div className="post-item-date">{formattedDate}</div>
    </div>
  );
}
