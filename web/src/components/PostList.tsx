import { usePostsList } from "../generated/api";
import { PostForm } from "./PostForm";
import { PostItem } from "./PostItem";

interface Props {
  selectedPostId: number | null;
  onSelectPost: (id: number) => void;
}

export function PostList({ selectedPostId, onSelectPost }: Props) {
  const { data, isLoading, error } = usePostsList();

  return (
    <div className="post-list">
      <PostForm />

      {isLoading && <div className="loading">Loading posts...</div>}

      {error && <div className="error">Failed to load posts</div>}

      {data?.items.map((post) => (
        <PostItem
          key={post.id}
          post={post}
          isSelected={selectedPostId === post.id}
          onClick={() => onSelectPost(post.id)}
        />
      ))}

      {data?.items.length === 0 && (
        <div className="empty">No posts yet. Create one!</div>
      )}
    </div>
  );
}
