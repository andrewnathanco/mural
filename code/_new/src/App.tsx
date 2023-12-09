import { type Component } from "solid-js";
import { GameProvider } from "./components/game/context";
import Index from "./view";

const App: Component = () => {
  return (
    <GameProvider>
      <Index />
    </GameProvider>
  );
};

export default App;
