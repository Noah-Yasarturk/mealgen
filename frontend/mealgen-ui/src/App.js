import './App.css';
import AppHeader from './header/header.tsx';
import Placemat from './placemat/placemat.tsx';

function App() {
  return (
    <div id="AppContainer" className="globalAppContainer">
      <AppHeader pageName="Manage Meals"></AppHeader>
      <Placemat></Placemat>
    </div>
  );
}

export default App;
