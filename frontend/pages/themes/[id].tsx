import { useForm } from 'react-hook-form'
import Link from "next/link"

type Theme = {
  id: number
  name: string
  detail: string
}

type Comment = {
  id: number
  body: string
  theme_id: number
}

export default function Theme({
  theme, comments, params_id
}: {
  theme: Theme
  comments: Comment[]
  params_id: string
}) {
  const { register, handleSubmit } = useForm()
  const onSubmit = (data: any) => CreateComment(data, params_id) 

  return (
    <div>
      <h1>{theme.name}</h1>
      <div>{theme.detail}</div>
      <hr/>
      {comments.map(comment => {
        return (
          <div>{comment.body}</div>
        )
      })}

      <hr />
      <div>
        <h3>Create Comment</h3>
        <form onSubmit={handleSubmit(onSubmit)}>
          comment:
          <br />
          <textarea name="body" ref={register} />
          <br />
          <input type="submit" value="CREATE" />
        </form>
      </div>
      <Link href="http://localhost:3000/">TOP</Link>
    </div>
  )
}

const BACKEND_API_THEMES_BASE_PATH = "http://localhost:8080/service/v1/themes"

export async function getServerSideProps({ params }) {
  const response = await fetch(`${BACKEND_API_THEMES_BASE_PATH}/${params.id}`)
  const data = await response.json()


  return {
    props: {
      theme: data.theme,
      comments: data.comments,
      params_id: params.id
    }
  }
}

export const CreateComment = (data: any, params_id: string) => {
  const headers = new Headers()
  headers.append("Content-Type", "application/json")

  const requestOptions = {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data),

  }

  fetch(`${BACKEND_API_THEMES_BASE_PATH}/${params_id}/comments`, requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
}