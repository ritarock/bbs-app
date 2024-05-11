import { Post } from "../interfaces/post";
import { get, post } from "./utils";

const PostUrl = "http://localhost:8080/backend/api/v1/posts";

const getPostAll = async (token: string): Promise<Post[]> => {
  return await get<Post[]>(PostUrl, token);
};

const getPost = async (id: number, token: string): Promise<Post> => {
  return await get<Post>(`${PostUrl}/${id.toString()}`, token);
};

const createPost = async (data: Post, token: string): Promise<Post> => {
  return await post<Post>(PostUrl, data, token);
};

export { createPost, getPost, getPostAll };
