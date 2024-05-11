import { useCookies } from "react-cookie"
import { Comment } from "../../interfaces/comment"
import { CommentAPI } from "../../api"
import { useForm } from "react-hook-form"
import { useParams } from "react-router-dom"


const CommentForm = () => {
  const { id } = useParams()
  const [cookie] = useCookies(["token"])
  const defaultValues: Comment = {
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

  const onsubmit = async (data: Comment) => {
    await CommentAPI.createComment(+id!, data, cookie.token.token)
    window.location.reload()
  }

  const onerror = (err: unknown) => console.log(err)

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

export default CommentForm
