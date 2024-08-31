import { useForm } from "react-hook-form";
import { User } from "../../@types/user";
import { AuthAPI } from "../../api";
import Box from "../Box";

const SignupForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<User>();

  const onSubmit = async (data: User) => {
    await AuthAPI.signup(data);
    window.location.href = "/";
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
              value="signup"
            />
          </form>
        }
      />
    </>
  );
};

export default SignupForm;
