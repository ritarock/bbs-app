import { useEffect } from "react"
import { useCookies } from "react-cookie"
import LoginForm from "./loginForm"

const LoginPage = () => {
  const [, , removeCookie] = useCookies(["token"])

  useEffect(() => {
    removeCookie("token")
  })

  return (
    <>
      <div>login</div>
      <hr />
      <LoginForm />
    </>
  )
}

export default LoginPage
