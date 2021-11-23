
const BASE_TOPIC_URL = "http://localhost:8080/backend/topics"
const BASE_COMMENT_URL = "http://localhost:8080/backend/topics"

type Topic = {
  id: string
  title: string
  detail: string
}

type Comment = {
  id: string
  topic_id: string
  body: string
}

export default function Topic({
  topic, comments
}: {
  topic: Topic
  comments: string[]
}) {
  console.log(comments)
  return (
    <>
      <h1>{topic.title}</h1>
      <h2>{topic.detail}</h2>
      <ul>
        {comments.map(comment => {
          return (
            <li>
              {comment}
            </li>
          )
        })}
      </ul>
    </>
  )
}

export async function getServerSideProps({params}) {
  const topicResponse = await fetch(`${BASE_TOPIC_URL}/${params.id}`)
  const topicData = await topicResponse.json()

  const commentResponse = await fetch(`${BASE_TOPIC_URL}/${params.id}/comments`)
  const commentData = await commentResponse.json()

  return {
    props: {
      topic: topicData.data[0],
      comments: commentData.data.map(data => data.body)
    }
  }
}
