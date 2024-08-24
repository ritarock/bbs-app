import { useForm } from "react-hook-form";
import { Comment } from "../../@types/comment";
import { CommentAPI } from "../../api";
import Box from "../Box";
import React from "react";

interface CommentFormProps {
  postID: number;
}

const CommentForm: React.FC<CommentFormProps> = ({ postID }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Comment>();

  const onSubmit = async (data: Comment) => {
    await CommentAPI.createComment(postID, data);
    window.location.reload();
  };

  return (
    <>
      <Box
        content={
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex flex-col items-center"
          >
            <textarea
              className="border border-blue-400 w-5/6 mt-4 pl-1 h-32 resize-none"
              {...register("content", { required: true, maxLength: 255 })}
              placeholder="content"
            />
            {errors.content && <p>This field is required</p>}
            <br />

            <input
              type="submit"
              value="COMMENT"
            />
          </form>
        }
        offHeader={true}
      />
    </>
  );
};

export default CommentForm;
