import React, { useState, useEffect } from "react";
import { Navigate } from "react-router-dom";
import { requestLogin, setToken } from "../../services/requestAPI";
import './LoginForm.css'

function LoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLogged, setIsLogged] = useState(false);
  const [failedTryLogin, setFailedTryLogin] = useState(false);

  const login = async (e) => {
    e.preventDefault();
    try {
      const token = await requestLogin("/login", { email, password });
      setToken(token);

      localStorage.setItem("token", token);

      setIsLogged(true);
    } catch (error) {
      setFailedTryLogin(true);
      setIsLogged(false);
    }
  };

  useEffect(() => {
    setFailedTryLogin(false);
  }, [email, password]);

  if (isLogged) return <Navigate to="/vote" />;

  return (
    <form className="login-form">
      <label htmlFor="email">
        <span>Email:</span>
        <input
          type="text"
          id="email"
          value={email}
          onChange={({ target: { value } }) => setEmail(value)}
          placeholder="Type your e-mail here"
        />
      </label>

      <label htmlFor="password">
        <span>Password:</span>
        <input
          type="password"
          id="password"
          value={password}
          onChange={({ target: { value } }) => setPassword(value)}
          placeholder="Type your e-mail here"
        />
      </label>

      {failedTryLogin ? (
        <p className="login-failed">
          Incorrect e-mail or password. Please, try again.
        </p>
      ) : null}

      <button
        type="submit"
        className="btn"
        onClick={(e) => login(e)}
      >
        Login
      </button>
    </form>
  );
}

export default LoginForm;
