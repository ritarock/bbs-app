import { useAuth } from "../contexts/AuthContext";

export function Header() {
  const { user, isAuthenticated, logout } = useAuth();

  return (
    <header className="app-header">
      <h1>BBS App</h1>

      {isAuthenticated && user ? (
        <div className="user-info">
          <span>{user.email}</span>
          <button onClick={logout} className="logout-button">
            Logout
          </button>
        </div>
      ) : null}
    </header>
  );
}
