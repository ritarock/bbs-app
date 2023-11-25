import LoginForm from "./form";
import { useCookies } from "react-cookie";
import { useEffect } from "react";

const LoginPage = () => {
  const [, , removeCookie] = useCookies(["token"]);

  useEffect(() => {
    removeCookie("token");
  });

  return (
    <>
      <div>login</div>
      <hr />
      <LoginForm />
    </>
  );
};

export default LoginPage;
