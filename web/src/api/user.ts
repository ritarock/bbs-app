import { User } from "../interfaces/user"
import { postNoToken } from "./utils"

const BaseUrl = "http://localhost:8080/backend"

type SignupSuccessResponse = {
  status: "success"
}

type SignupFailureResponse = {
  status: "error"
}

type SignupResponse = SignupSuccessResponse | SignupFailureResponse

type LoginResponse = {
  message?: "Unauthorized"
  token?: string
}

const signup = async (data: User): Promise<SignupResponse> => {
  const response = await postNoToken<User>(`${BaseUrl}/signup`, data)

  if (response as string !== "success") {
    return { status: "error" }
  }

  return { status: "success"}
}

const login = async (data: User): Promise<LoginResponse> => {
  const response: LoginResponse = await postNoToken<User>(
    `${BaseUrl}/login`,
    data,
  )

  if (response.message) {
    return { message: "Unauthorized" }
  }

  return { token: response.token }
}

export { signup, login }
