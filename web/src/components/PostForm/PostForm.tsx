import { useForm } from "react-hook-form";
import Box from "../Box";
import { Post } from "../../@types/post";
import { PostAPI } from "../../api";

const PostForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Post>();

  const onSubmit = async (data: Post) => {
    await PostAPI.createPost(data);
    window.location.href = "/";
  };

  return (
    <>
      <Box
        content={
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex flex-col items-center"
          >
            <input
              className="border border-blue-400 w-5/6 mt-4 pl-1"
              {...register("title", { required: true, maxLength: 30 })}
              placeholder="title"
            />
            {errors.title && <p>This field is required</p>}
            <br />

            <textarea
              className="border border-blue-400 w-5/6 mt-4 pl-1 h-32 resize-none"
              {...register("content", { required: true, maxLength: 255 })}
              placeholder="content"
            />
            {errors.content && <p>This field is required</p>}
            <br />

            <input
              type="submit"
              value="POST"
            />
          </form>
        }
      />
    </>
  );
};

export default PostForm;
