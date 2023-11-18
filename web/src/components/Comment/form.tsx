import { useForm } from "react-hook-form";
import { Post } from "../../interfaces/post";
import { CommentAPI } from "../../api";
import { Comment } from "../../interfaces/comment";
import { useParams } from "react-router-dom";

const CommentForm = () => {
  const defaultValues: Comment = {
    content: "",
  };
  const { id } = useParams();

  const { register, handleSubmit, formState: { errors, isDirty, isValid } } =
    useForm({
      defaultValues,
    });

  const onsubmit = async (data: Post) => {
    await CommentAPI.createComment(+id!, data);
    window.location.reload();
  };

  const onerror = (err: unknown) => console.log(err);

  return (
    <>
      <form onSubmit={handleSubmit(onsubmit, onerror)} noValidate>
        <div>
          <label htmlFor="content">content:</label>
          <br />
          <textarea
            id="content"
            {...register("content", {
              required: "content is required",
              maxLength: {
                value: 255,
                message: "title length <= 255",
              },
            })}
          />
          <div>{errors.content?.message}</div>
        </div>
        <div>
          <button type="submit" disabled={!isDirty || !isValid}>send</button>
        </div>
      </form>
    </>
  );
};

export default CommentForm;
