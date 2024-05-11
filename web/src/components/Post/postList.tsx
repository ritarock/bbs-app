import { useEffect, useState } from "react"
import { Post } from "../../interfaces/post"
import { useCookies } from "react-cookie"
import { PostAPI } from "../../api"

const PostList = () => {
  const [posts, setPosts] = useState<Array<Post>>([])
  const [cookie] = useCookies(["token"])

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const getPosts = await PostAPI.getPostAll(cookie.token.token)
        setPosts([...getPosts])
      } catch {
        console.log("Error fetching posts: ", console.error);
      }
    }

    fetchPosts()
  }, [cookie])

  console.log(cookie)
  return (
    <>
      {Object.keys(cookie).length === 0
        ? (
          <>
            <p>
              <a href="/login">login</a>
            </p>
            <p>
              <a href="/signup">signup</a>
            </p>
          </>
        ) : (
          <>
            <div>
              {posts.length > 0 &&
                (
                  <ul>
                    {posts.map((post) => (
                      <li key={post.id}>
                        <a href={`/posts/${post.id}`}>
                          {post.title}
                        </a>
                      </li>
                    ))}
                  </ul>
                )
              }
            </div>
          </>
        )
      }
    </>
  )
}

export default PostList
