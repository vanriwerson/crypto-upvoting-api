import './App.css';
import Router from './routes';
import NavLinks from './components/NavLinks';

function App() {
  return (
    <div className="App">
      <NavLinks />
      <h1>Users Favorite Crypto</h1>
      <Router />
    </div>
  );
}

export default App;
