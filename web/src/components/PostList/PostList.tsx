import { useEffect, useState } from "react"
import { Post } from "../../@types/post"
import { PostAPI } from "../../api"
import Box from "../Box"

const PostList = () => {
  const [posts, setPosts] = useState<Post[]>([])

  useEffect(() => {
    const fetchPost = async () => {
      try {
        const getPostAll = await PostAPI.getPostAll()
        setPosts(getPostAll)
      } catch {
        console.log("Error fetching posts: ", console.error)
      }
    }

    fetchPost()
  }, [])

  return (
    <>
      <Box
        header={
          <a href="/">
            post
          </a>
        }
        content={
          <div>
      {posts.length > 0 && (
        <ul>
          {posts.map(post => (
            <li key={post.id} className="mb-4">
              <span className="font-bold text-lg text-sky-600">
                <a href={`/posts/${post.id}`}>
                  {post.title}
                </a>
              </span>
              <br />
              <span>{post.content}</span>
            </li>
          ))}
        </ul>
      )}
          </div>
        }
      />
    </>
  )
}

export default PostList
