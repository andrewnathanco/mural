import InfoButton from "../components/buttons/info-button";
import StatsButton from "../components/buttons/stats-button";
import GameArea from "../components/game/presentation/game-area";

export default function Index() {
  return (
    <div class="flex flex-col items-center justify-center">
      <div class="flex flex-col items-center space-y-4">
        <div class="flex flex-col space-y-4 w-full">
          <div class="text-5xl flex space-x-2 items-center">
            <div>Mural #5</div>
            <div
              id="game-version"
              class="font-semibold w-min h-min text-gray-600 text-xs border-2 px-1 border-river-bed-700 rounded-lg"
            >
              v0.1.1
            </div>
          </div>
          <div class="flex justify-between">
            <div class="flex flex-col space-y-1 items-start">
              <div id="game-theme" class="text-contessa-500 text-4xl">
                1970
              </div>
              <div class="text-md">Today's Theme</div>
            </div>
            <div class="flex flex-col space-y-1 items-start">
              <div id="game-sessions" class="text-contessa-500 text-4xl">
                5
              </div>
              <div class="text-md">Have Played</div>
            </div>
          </div>

          <div class="flex flex-col space-y-1">
            <div class="w-full">
              <InfoButton />
            </div>

            <div class="w-full">
              <StatsButton />
            </div>
          </div>

          <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>

          <div class="text-3xl flex space-x-2 items-center flex-col">
            <div id="game-area">
              <GameArea />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
