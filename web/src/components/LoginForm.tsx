import { useState, FormEvent } from "react";
import { useAuthSignin, useAuthSignup } from "../generated/api";
import { useAuth } from "../contexts/AuthContext";

export function LoginForm() {
  const [isSignUp, setIsSignUp] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);

  const { login } = useAuth();

  const signInMutation = useAuthSignin();
  const signUpMutation = useAuthSignup();

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      if (isSignUp) {
        const result = await signUpMutation.mutateAsync({
          data: { email, password },
        });
        login(result.token, result.user);
      } else {
        const result = await signInMutation.mutateAsync({
          data: { email, password },
        });
        login(result.token, result.user);
      }
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "An authentication error occurred"
      );
    }
  };

  const isLoading = signInMutation.isPending || signUpMutation.isPending;

  return (
    <div className="login-container">
      <div className="login-form">
        <h2>{isSignUp ? "Sign Up" : "Login"}</h2>

        {error && <div className="error-message">{error}</div>}

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              disabled={isLoading}
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              id="password"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              minLength={8}
              disabled={isLoading}
            />
          </div>

          <button type="submit" disabled={isLoading}>
            {isLoading ? "Processing..." : isSignUp ? "Sign Up" : "Login"}
          </button>
        </form>

        <button
          type="button"
          className="toggle-mode"
          onClick={() => setIsSignUp(!isSignUp)}
          disabled={isLoading}
        >
          {isSignUp ? "Already have an account? Login" : "Need an account? Sign Up"}
        </button>
      </div>
    </div>
  );
}
