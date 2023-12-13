import { createSignal } from "solid-js";
import ShareButton from "../../buttons/share-button";
import { Game } from "../../game/model/game";
import { GameDifficulty } from "../../game/presentation/board/difficulty-selector";
import { useShareDialog } from "./context";
import { useGame } from "../../game/context/game";
import { User } from "../../game/model/user";
import { useUser } from "../../game/context/game-difficulty";

export function create_share_url(game: Game) {
  let share_url = `${import.meta.env.VITE_SHARE_URL}/share`;

  let flipped = game.flipped.join(",");

  share_url += `?flipped=${flipped}`;
  share_url += `&name=${game.user_name}`;
  if (game.selected_option) {
    share_url += `&answer_id=${game.selected_option?.id}`;
  }
  share_url += `&correct_id=${game.correct_option.id}`;

  return share_url;
}

export function generate_share_data(game: Game, user: User): ShareData {
  return {
    url: create_share_url(game),
    title: "Mural Share",
    text: `${game.user_name ? `${game.user_name}'s ` : ""}Mural${
      user.difficulty == GameDifficulty.hard ? "*" : ""
    } #${game.game_key}\nScore: ${game.score}`,
  };
}

export function ShareLink(props: { share_url: string; share_data: ShareData }) {
  return (
    <div class="flex flex-col w-full">
      <div class="flex flex-col space-y-2">
        <div class="relative w-full">
          <input
            type="share-link"
            id="share-link"
            class="block p-2 rounded-lg w-full z-20 text-sm text-river-bed-700 bg-desert-sand-50 border-2 border-river-bed-300 focus:ring-contessa-500 focus:contessa-500 "
            value={props.share_url}
          />
          <button
            type="submit"
            class="absolute shadow-lg top-0 end-0 h-full p-2 text-sm font-medium text-desert-sand-50 bg-river-bed-500 border-river-bed-500 rounded-e-lg border  hover:bg-river-bed-600 focus:ring-4 focus:outline-none focus:ring-contessa-600 "
            onclick={() => {
              navigator?.clipboard?.writeText(props.share_url);
            }}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M15.666 3.888A2.25 2.25 0 0013.5 2.25h-3c-1.03 0-1.9.693-2.166 1.638m7.332 0c.055.194.084.4.084.612v0a.75.75 0 01-.75.75H9a.75.75 0 01-.75-.75v0c0-.212.03-.418.084-.612m7.332 0c.646.049 1.288.11 1.927.184 1.1.128 1.907 1.077 1.907 2.185V19.5a2.25 2.25 0 01-2.25 2.25H6.75A2.25 2.25 0 014.5 19.5V6.257c0-1.108.806-2.057 1.907-2.185a48.208 48.208 0 011.927-.184"
              />
            </svg>
            <span class="sr-only">Copy</span>
          </button>
        </div>
        <button
          onclick={() => {
            navigator?.share(props.share_data);
          }}
          class="
      flex space-x-2
      w-full p-1 text-base text-desert-sand-100 rounded-md 
      justify-center items-center
      border-dingley-700 border-2
        hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800
    "
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-6 h-6"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M6 12L3.269 3.126A59.768 59.768 0 0121.485 12 59.77 59.77 0 013.27 20.876L5.999 12zm0 0h7.5"
            />
          </svg>

          <div>Send</div>
        </button>
      </div>
    </div>
  );
}

export function ShareDialog() {
  const [is_open, { open, close }] = useShareDialog();
  const [name, set_name] = createSignal("");
  const [share_data, set_share_data] = createSignal<ShareData>({});
  const [share_url, set_share_url] = createSignal<string>("");
  const [game, set_game] = useGame();
  const [user, _] = useUser();

  return (
    <div classList={{ hidden: !is_open(), block: is_open() }}>
      <div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div>
      <div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 right-0 md:w-128 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto">
        <div id="dialog-content" class="p-8 flex flex-col space-y-2 w-full">
          <div id="info-dialog">
            <div class="flex flex-col space-y-2">
              <div class="flex justify-between items-center">
                <div class="text-3xl">Share</div>
                <button
                  onClick={() => {
                    close();
                  }}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-6 h-6"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M5.47 5.47a.75.75 0 011.06 0L12 10.94l5.47-5.47a.75.75 0 111.06 1.06L13.06 12l5.47 5.47a.75.75 0 11-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 01-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 010-1.06z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                </button>
              </div>
              <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
              {share_data().url == undefined ? (
                <div id="create-link-form" class="flex flex-col space-y-2">
                  <input
                    id="name"
                    name="name"
                    value={name()}
                    oninput={(e) => set_name(e.target.value)}
                    placeholder="Enter your name..."
                    class="block w-full p-2 text-sm text-river-bed-700 bg-desert-sand-50 border-2 border-river-bed-300 placeholder:text-river-bed-500 rounded-lg focus:ring-river-bed-700 focus:border-river-bed-700"
                  />
                  <button
                    onclick={(e) => {
                      e.preventDefault();
                      set_game("user_name", name());
                      set_share_data(generate_share_data(game, user));
                      set_share_url(create_share_url(game));
                    }}
                    type="submit"
                    id="share-button"
                    class="w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-dingley-700 border-2 hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800"
                  >
                    Create Share Link
                  </button>
                </div>
              ) : (
                <ShareLink share_data={share_data()} share_url={share_url()} />
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
