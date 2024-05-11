import { useEffect, useState } from "react"
import { useCookies } from "react-cookie"
import { useParams } from "react-router-dom"
import { CommentAPI } from "../../api"
import { Comment } from "../../interfaces/comment"

const CommentList = () => {
  const { id } = useParams()
  const [comments, setComments] = useState<Array<Comment>>([])
  const [cookie] = useCookies(["token"])

  useEffect(() => {
    const fetchComments = async () => {
      try {
        const getComments = await CommentAPI.getCommentAll(
          +id!,
          cookie.token.token,
        )
        setComments([...getComments])
      } catch (error) {
        console.log("Error fetching comments: ", error)
      }
    }

    fetchComments()
  }, [id, cookie])

  return (
    <>
      {comments.map((comment) => (
        <li key={comment.id}>
          {comment.content}
        </li>
      ))}
    </>
  )
}

export default CommentList
