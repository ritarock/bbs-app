import { Post } from "../@types/post";
import { get, post } from "./util";

const PostBaseUrl = "http://localhost:8080/backend/api/v1/posts";

const getPostAll = async (): Promise<Post[]> => {
  return await get(PostBaseUrl);
};

const getPost = async (id: number): Promise<Post> => {
  return await get(`${PostBaseUrl}/${id}`);
};

const createPost = async (data: Post): Promise<Post> => {
  return await post(PostBaseUrl, data);
};

export { createPost, getPost, getPostAll };
