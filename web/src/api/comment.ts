import { Comment } from "../interfaces/comment"
import { get, post } from "./utils"

const BaseUrl = "http://localhost:8080/backend/api/v1/post"

const getCommentAll = async (
  id: number,
  token: string
): Promise<Comment[]> => {
  return await get<Comment[]>(`${BaseUrl}/${id}/comments`, token)
}

const createComment = async (
  postID: number,
  data: Comment,
  token: string
): Promise<Comment> => {
  return await post<Comment>(`${BaseUrl}/${postID}/comments`,data,token)
}

export { getCommentAll, createComment }
