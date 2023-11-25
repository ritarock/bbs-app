import { Post } from "../interfaces/post";
import { get, post } from "./util";

const BaseUrl = "http://localhost:8080/backend/api/v1/posts";

const getPostsAll = async (token: string): Promise<Post[]> => {
  const response = await get<Post[]>(BaseUrl, token);
  return response;
};

const getPost = async (id: number, token: string): Promise<Post> => {
  const response = await get<Post>(`${BaseUrl}/${id.toString()}`, token);
  return response;
};

const createPost = async (data: Post, token: string): Promise<Post> => {
  const response = await post<Post>(BaseUrl, data, token);
  return response;
};

export { createPost, getPost, getPostsAll };
