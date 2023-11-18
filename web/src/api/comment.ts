import { Comment } from "../interfaces/comment";
import { get, post } from "./util";

const BaseUrl = "http://localhost:8080/backend/api/v1/post";

const getCommentsAll = async (id: number): Promise<Comment[]> => {
  const response = await get<Comment[]>(`${BaseUrl}/${id}/comments`);
  return response;
};

const createComment = async (
  postId: number,
  data: Comment,
): Promise<Comment> => {
  const response = await post<Comment>(`${BaseUrl}/${postId}/comments`, data);
  return response;
};

export { createComment, getCommentsAll };
