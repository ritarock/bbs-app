const BASE_TOPIC_URL = 'http://localhost:8080/backend/topics'
import { SubmitHandler, useForm } from 'react-hook-form'
import { useRouter } from 'next/router'
import { Topic } from '../../interfaces'

type Inputs = {
  body: string
}

export default function Topic({
  topic,
  comments,
}: {
  topic: Topic
  comments: string[]
}) {
  const router = useRouter()
  const id = router.query.id[0]
  const {
    register,
    handleSubmit,
  } = useForm<Inputs>()
  const onSubmit: SubmitHandler<Inputs> = (data) => createComent(id, data)

  return (
    <>
      <h1>{topic.title}</h1>
      <h2>{topic.detail}</h2>
      <ul>
        {comments.map((comment) => {
          return (
            // eslint-disable-next-line react/jsx-key
            <li>{comment}</li>
          )
        })}
      </ul>
      <hr />
      <div>
        <h3>Create Topic</h3>
        <form onSubmit={handleSubmit(onSubmit)}>
          Comment :
          <input {...register('body')} />
          <br />
          <input type="submit" value="create" />
        </form>
      </div>
    </>
  )
}

export async function getServerSideProps({ params }) {
  const topicResponse = await fetch(`${BASE_TOPIC_URL}/${params.id}`)
  const topicData = await topicResponse.json()

  const commentResponse = await fetch(`${BASE_TOPIC_URL}/${params.id}/comments`)
  const commentData = await commentResponse.json()

  console.log(commentData.data === null)
  return {
    props: {
      topic: topicData.data[0],
      comments:
        commentData.data !== null
          ? commentData.data.map((data) => data.body)
          : [],
    },
  }
}

export const createComent = (id: string, data: Inputs) => {
  const method = 'POST'
  const headers = {
    'Content-type': 'application/json',
  }
  const body = JSON.stringify(data)

  fetch(`${BASE_TOPIC_URL}/${id}/comments`, { method, headers, body })
    .then((response) => response.text())
    .then((result) => console.log(result))
}
