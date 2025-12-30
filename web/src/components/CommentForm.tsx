import { useState, type FormEvent } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useCommentsCreate, getCommentsListQueryKey } from "../generated/api";

interface Props {
  postId: number;
}

export function CommentForm({ postId }: Props) {
  const [body, setBody] = useState("");
  const queryClient = useQueryClient();

  const createComment = useCommentsCreate({
    mutation: {
      onSuccess: () => {
        setBody("");
        queryClient.invalidateQueries({
          queryKey: getCommentsListQueryKey(postId),
        });
      },
    },
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (!body.trim()) return;

    createComment.mutate({ postId, data: { body } });
  };

  return (
    <form className="comment-form" onSubmit={handleSubmit}>
      <textarea
        className="comment-form-textarea"
        placeholder="Add a comment..."
        value={body}
        onChange={(e) => setBody(e.target.value)}
        required
      />
      <button
        type="submit"
        className="comment-form-button"
        disabled={createComment.isPending}
      >
        {createComment.isPending ? "Sending..." : "Comment"}
      </button>
    </form>
  );
}
