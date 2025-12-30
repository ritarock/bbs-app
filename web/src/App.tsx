import "./index.css";
import { BBS } from "./components/BBS";
import { Header } from "./components/Header";
import { LoginForm } from "./components/LoginForm";
import { useAuth } from "./contexts/AuthContext";

export function App() {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return <div className="loading">Loading...</div>;
  }

  return (
    <div className="app">
      <Header />
      <main>{isAuthenticated ? <BBS /> : <LoginForm />}</main>
    </div>
  );
}

export default App;
