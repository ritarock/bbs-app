import { createBrowserRouter } from "react-router-dom";
import App from "./App";
import PostForm from "./components/PostForm";
import PostPage from "./components/PostPage";
import SignupForm from "./components/SignupForm";
import LoginForm from "./components/LoginForm";

const route = createBrowserRouter([
  { path: "/", element: <App /> },
  { path: "/post", element: <PostForm /> },
  { path: "/posts/:id", element: <PostPage /> },
  { path: "/signup", element: <SignupForm /> },
  { path: "/login", element: <LoginForm /> },
]);

export default route;
