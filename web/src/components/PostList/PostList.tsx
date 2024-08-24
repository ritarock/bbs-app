import { useEffect, useState } from "react"
import { Post } from "../../@types/post"
import { PostAPI } from "../../api"

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
    <div className="box-border h-vh w-5/6 border-2 border-slate-500 mx-auto">
      <p className="bg-blue-100 text-right pr-4">
        post
      </p>
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
    </>
  )
}

export default PostList
