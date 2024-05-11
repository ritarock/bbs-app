import PostForm from "./components/Post/postForm"
import PostList from "./components/Post/postList"

function App() {
  return (
    <>
      <div>BBS-APP</div>
      <hr />
      <div>
        <PostList />
        <PostForm />
      </div>
    </>
  )
}

export default App
