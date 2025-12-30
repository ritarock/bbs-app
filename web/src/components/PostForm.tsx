import { useState, type FormEvent } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { usePostsCreate, getPostsListQueryKey } from "../generated/api";

interface Props {
  onPostCreated?: () => void;
}

export function PostForm({ onPostCreated }: Props) {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const queryClient = useQueryClient();

  const createPost = usePostsCreate({
    mutation: {
      onSuccess: () => {
        setTitle("");
        setContent("");
        queryClient.invalidateQueries({ queryKey: getPostsListQueryKey() });
        onPostCreated?.();
      },
    },
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!title.trim() || !content.trim()) return;

    createPost.mutate({ data: { title, content } });
  };

  return (
    <form className="post-form" onSubmit={handleSubmit}>
      <input
        type="text"
        className="post-form-input"
        placeholder="Title (max 30 chars)"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        maxLength={30}
        required
      />
      <textarea
        className="post-form-textarea"
        placeholder="Content (max 255 chars)"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        maxLength={255}
        required
      />
      <button
        type="submit"
        className="post-form-button"
        disabled={createPost.isPending}
      >
        {createPost.isPending ? "Posting..." : "Post"}
      </button>
    </form>
  );
}
