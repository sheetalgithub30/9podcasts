import "./App.css";
import LoginPage from "./auth/LoginPage";
import HomePage from "./home/HomePage";

const loggedIn = true;

function App() {
  return <>{loggedIn ? <HomePage /> : <LoginPage />}</>;
}

export default App;
