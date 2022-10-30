import './App.css';
import Router from './routes';
import NavLinks from './components/NavLinks/NavLinks';

function App() {
  return (
    <div className="App">
      <NavLinks />
      <Router />
    </div>
  );
}

export default App;
