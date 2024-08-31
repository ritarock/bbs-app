import { User } from "../@types/user";
import { postNoToken } from "./util";

const SignupUrl = "http://localhost:8080/backend/signup";
const LoginUrl = "http://localhost:8080/backend/login";

type loginResponse = {
  token: string;
};

const signup = async (data: User): Promise<User> => {
  return await postNoToken(SignupUrl, data);
};

const login = async (data: User): Promise<loginResponse> => {
  return await postNoToken(LoginUrl, data);
};

export { login, signup };
