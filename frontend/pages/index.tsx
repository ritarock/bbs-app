import {GetServerSideProps} from "next"
import Link from "next/link"

const BASE_URL = "http://localhost:8080/backend/topics"

type Topic = {
  id: string
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
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async () => {
  const response = await fetch(BASE_URL)
  const data = await response.json()

  return {
    props: {
      index: data
    }
  }
}
