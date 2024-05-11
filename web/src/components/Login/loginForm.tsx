import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";
import { User } from "../../interfaces/user";
import { useForm } from "react-hook-form";
import { UserAPI } from "../../api";

const LoginForm = () => {
  const navigate = useNavigate();
  const [, setCookie] = useCookies(["token"]);
  const defaultValues: User = {
    name: "",
    password: "",
  };
  const {
    register,
    handleSubmit,
    formState: {
      errors,
      isDirty,
      isValid,
    },
  } = useForm({ defaultValues });

  const onsubmit = async (data: User) => {
    const response = await UserAPI.login(data);
    if (response.token !== undefined) {
      setCookie("token", response);
      navigate("/", { replace: true });
    } else {
      alert("something wrong");
    }
  };

  const onerror = (err: unknown) => console.log(err);

  return (
    <>
      <form onSubmit={handleSubmit(onsubmit, onerror)} noValidate>
        <div>
          <label htmlFor="name">name:</label>
          <br />
          <input
            id="name"
            type="text"
            {...register("name", {
              required: "name is required",
              maxLength: {
                value: 30,
                message: "name length <= 30",
              },
            })}
          />
          <div>{errors.name?.message}</div>
        </div>

        <div>
          <label htmlFor="password">password:</label>
          <br />
          <input
            id="password"
            {...register("password", {
              required: "password is required",
              maxLength: {
                value: 30,
                message: "8 <= password length <= 30",
              },
            })}
          />
          <div>{errors.password?.message}</div>
        </div>

        <div>
          <button type="submit" disabled={!isDirty || !isValid}>
            SEND
          </button>
        </div>
      </form>
    </>
  );
};

export default LoginForm;
