import {GetServerSideProps} from "next"
import Link from "next/link"
import {SubmitHandler, useForm} from "react-hook-form"

const BASE_TOPIC_URL = "http://localhost:8080/backend/topics"

type Topic = {
  id: string
  title: string
  detail: string
}

type Inputs = {
  title: string
  detail: string
}

export default function Home({
  index
}: {
  index: {
    code: number
    data: Topic[]
  }
}) {
  const {register, handleSubmit, watch, formState: {errors}} = useForm<Inputs>()
  const onSubmit: SubmitHandler<Inputs> = data => createTopic(data)
  return (
    <>
      <div>
        <h1>Topic</h1>
        <ul>
          {index.data.map(data => {
            return (
              <li>
                <Link href={`/topics/${data.id}`}>
                  <a>{data.title}</a>
                </Link>
              </li>
            )
          })}
        </ul>
      </div>
      <hr />
      <div>
        <h3>Create Topic</h3>
          <form onSubmit={handleSubmit(onSubmit)}>
            Title :
            <input {...register("title")} />
            <br />
            detail :
            <input {...register("detail")} />
            <br />
            <input type="submit" value="create" />
          </form>
      </div>
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async () => {
  const response = await fetch(BASE_TOPIC_URL)
  const data = await response.json()

  return {
    props: {
      index: data
    }
  }
}

export const createTopic = (data: Inputs) => {
  const method = "POST"
  const headers = {
    "Content-type": "application/json"
  }
  const body = JSON.stringify(data)

  fetch(BASE_TOPIC_URL, {method, headers, body})
    .then(response => response.text())
    .then(result => console.log(result))
}
