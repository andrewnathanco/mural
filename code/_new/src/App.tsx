import { type Component } from "solid-js";
import { GameProvider } from "./components/game/context";
import Index from "./view";
import { InfoDialogProvider } from "./components/dialog/info/context";
import { ShareDialogProvider } from "./components/dialog/share/context";

const App: Component = () => {
  return (
    <GameProvider>
      <InfoDialogProvider>
        <ShareDialogProvider>
          <Index />
        </ShareDialogProvider>
      </InfoDialogProvider>
    </GameProvider>
  );
};

export default App;
