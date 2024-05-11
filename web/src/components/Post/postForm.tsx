import { useCookies } from "react-cookie"
import { Post } from "../../interfaces/post"
import { PostAPI } from "../../api"
import { useForm } from "react-hook-form"

const PostForm = () => {
  const [cookie] = useCookies(["token"])
  const defaultValues: Post = {
    title: "",
    content: "",
  }
  const {
    register,
    handleSubmit,
    formState: {
      errors,
      isDirty,
      isValid
    }
  } = useForm({ defaultValues })

  const onsubmit = async (data: Post) => {
    await PostAPI.createPost(data, cookie.token.token)
    window.location.reload()
  }

  const onerror = (err: unknown) => console.log(err)

  return (
    <>
      <form onSubmit={handleSubmit(onsubmit, onerror)} noValidate>
        <div>
          <label htmlFor="title">title:</label>
          <br />
          <input
            id="title"
            type="text"
            {...register("title", {
              required: "title is required",
              maxLength: {
                value: 30,
                message: "title length <= 30",
              }
            })}
          />
          <div>{errors.title?.message}</div>
        </div>
        
        <div>
          <label htmlFor="content">content:</label>
          <br />
          <textarea
            id="content"
            {...register("content", {
              required: "content is required",
              maxLength: {
                value: 255,
                message: "content length <= 255",
              }
            })}
          />
          <div>{errors.content?.message}</div>
        </div>

        <div>
          <button type="submit" disabled={!isDirty || !isValid}>
            SEND
          </button>
        </div>
      </form>
    </>
  )
}

export default PostForm
