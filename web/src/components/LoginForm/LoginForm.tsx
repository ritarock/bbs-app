import { useForm } from "react-hook-form";
import { User } from "../../@types/user";
import { AuthAPI } from "../../api";
import Box from "../Box";
import { useCookies } from "react-cookie";

const LoginForm = () => {
  const [, setCookie] = useCookies(["token"]);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<User>();

  const onSubmit = async (data: User) => {
    const response = await AuthAPI.login(data);
    if (response.token !== undefined) {
      setCookie("token", response.token);
      window.location.href = "/";
    } else {
      alert("something wrong");
    }
  };

  return (
    <>
      <Box
        content={
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex flex-col items-center"
          >
            <input
              className="border border-blue-400 w-5/6 mt-4 pl-1"
              {...register("name", { required: true, maxLength: 30 })}
              placeholder="name"
            />
            {errors.name && <p>This field is required</p>}
            <br />

            <input
              type="password"
              className="border border-blue-400 w-5/6 mt-4 pl-1"
              {...register("password", { required: true, maxLength: 30 })}
              placeholder="password"
            />
            {errors.password && <p>This field is required</p>}
            <br />

            <input
              type="submit"
              value="login"
            />
          </form>
        }
      />
    </>
  );
};

export default LoginForm;
