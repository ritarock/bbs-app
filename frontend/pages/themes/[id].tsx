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
  theme, comments
}: {
  theme: Theme
  comments: Comment[]
}) {
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
      comments: data.comments
    }
  }
}
