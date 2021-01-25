import { GetServerSideProps } from "next"
import Link from "next/link"
import { useForm } from 'react-hook-form'

const BACKEND_API_THEMES_BASE_PATH = "http://localhost:8080/service/v1/themes"

type Theme = {
    id: number
    name: string
    detail: string
}

export default function Home({
  indexData
}: {
  indexData : Theme[]
}) {

  const { register, handleSubmit } = useForm()
  const onSubmit = (data: any) => CreateTheme(data) 

  return (
    <div>
      <h1>Theme</h1>
      <div>
        <ul>
          {indexData.map(data => {
            return (
              <li>
                <Link href={`/themes/${data.id}`}>
                  <a>{data.name}</a>
                </Link>
                &nbsp;{data.detail}
              </li>
            )
          })}
        </ul>
      </div>

      <hr />
      <div>
        <h3>Create Theme</h3>
        <form onSubmit={handleSubmit(onSubmit)}>
          Theme:
          <br />
          <input name="name" ref={register} />
          <br />
          detail:
          <br />
          <textarea name="detail" ref={register} />
          <br />
          <input type="submit" value="CREATE" />
        </form>
      </div>
    </div>
  )
}

export const getServerSideProps: GetServerSideProps = async () => {
  const response = await fetch(BACKEND_API_THEMES_BASE_PATH)
  const data = await response.json()

  console.log(data.themes)

  return {
    props: {
      indexData: data.themes
    }
  }
}

export const CreateTheme = (data: any) => {
  const headers = new Headers()
  headers.append("Content-Type", "application/json")

  const requestOptions = {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  }

  fetch(BACKEND_API_THEMES_BASE_PATH, requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
}