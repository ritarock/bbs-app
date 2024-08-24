import { Comment } from "../@types/comment";
import { get, post } from "./util";

const CommentBaseUrl = "http://localhost:8080/backend/api/v1/post";

const getCommentAll = async (postID: number): Promise<Comment[]> => {
  return await get(`${CommentBaseUrl}/${postID}/comments`);
};

const createComment = async (
  postID: number,
  data: Comment,
): Promise<Comment> => {
  return await post(`${CommentBaseUrl}/${postID}/comments`, data);
};

export { createComment, getCommentAll };
