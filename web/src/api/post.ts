import { Post } from "../interfaces/post"
import { get, post } from "./utils"

const BaseUrl = "http://localhost:8080/backend/api/v1/posts"

const getPostAll = async (token: string): Promise<Post[]> => {
  return await get<Post[]>(BaseUrl, token)
}

const getPost = async (id: number, token: string): Promise<Post> => {
  return await get<Post>(`${BaseUrl}/${id.toString()}`, token)
}

const createPost = async (data: Post, token: string): Promise<Post> => {
  return await post<Post>(BaseUrl, data, token)
}

export { getPostAll, getPost, createPost }
