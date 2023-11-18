import { Post } from "../interfaces/post";
import { get, post } from "./util";

const BaseUrl = "http://localhost:8080/backend/api/v1/posts";

const getPostsAll = async (): Promise<Post[]> => {
  const response = await get<Post[]>(BaseUrl);
  return response;
};

const getPost = async (id: number): Promise<Post> => {
  const response = await get<Post>(`${BaseUrl}/${id.toString()}`);
  return response;
};

const createPost = async (data: Post): Promise<Post> => {
  const response = await post<Post>(BaseUrl, data);
  return response;
};

export { createPost, getPost, getPostsAll };
