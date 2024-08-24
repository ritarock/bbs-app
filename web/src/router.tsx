import { createBrowserRouter } from "react-router-dom";
import App from "./App";
import PostForm from "./components/PostForm";
import PostPage from "./components/PostPage";

const route = createBrowserRouter([
  { path: "/", element: <App /> },
  { path: "/post", element: <PostForm /> },
  { path: "/posts/:id", element: <PostPage /> },
]);

export default route;
