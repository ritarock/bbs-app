import { useCookies } from "react-cookie";
import PostList from "./components/PostList";
import Box from "./components/Box";

function App() {
  const [cookie] = useCookies(["token"]);

  return (
    <>
      {Object.keys(cookie).length === 0
        ? (
          <Box
            content={
              <>
                <div>
                  <a href="/signup">
                    signup
                  </a>
                </div>
                <div>
                  <a href="/login">
                    login
                  </a>
                </div>
              </>
            }
          />
        )
        : <PostList />}
    </>
  );
}

export default App;
