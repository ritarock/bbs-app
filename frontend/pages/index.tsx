import { GetServerSideProps } from "next"
import Link from "next/link"

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
