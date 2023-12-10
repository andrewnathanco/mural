import { type Component } from "solid-js";
import { GameProvider } from "./components/game/context";
import Index from "./view";
import { InfoDialogProvider } from "./components/dialog/info/context";
import { ShareDialogProvider } from "./components/dialog/share/context";
import { HintDialogProvider } from "./components/dialog/hint/context";

const App: Component = () => {
  return (
    <GameProvider>
      <InfoDialogProvider>
        <ShareDialogProvider>
          <HintDialogProvider>
            <Index />
          </HintDialogProvider>
        </ShareDialogProvider>
      </InfoDialogProvider>
    </GameProvider>
  );
};

export default App;
