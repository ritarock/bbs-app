import { useForm } from "react-hook-form"
import { User } from "../../interfaces/user"
import { useNavigate } from "react-router-dom"
import { UserAPI } from "../../api"

const SignupForm = () => {
  const defaultValues: User = {
    name: "",
    password: ""
  }

  const {
    register,
    handleSubmit,
    formState: {
      errors,
      isDirty,
      isValid
    }
  } = useForm({ defaultValues })

  const navigate = useNavigate()

  const onsubmit = async (data: User) => {
    const response = await UserAPI.signup(data)
    if (response.status === "success") {
      navigate("/login", { replace: true })
    } else {
      alert("already exists")
    }
  }

  const onerror = (err: unknown) => console.log(err)

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
              }
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
              }
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
  )
}

export default SignupForm
