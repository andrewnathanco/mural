import { createEffect, createSignal } from "solid-js";
import { get_countdown_till_next_game } from "../../game/service";
import { useShareWarningDialog } from "./context";

export default function ShareWarningDialog() {
  const [is_open, { open, close }] = useShareWarningDialog();
  const [countdown, set_countdown] = createSignal(
    get_countdown_till_next_game()
  );

  createEffect(() => {
    setInterval(() => {
      set_countdown(get_countdown_till_next_game());
    }, 1000);
  });

  return (
    <div
      classList={{
        hidden: !is_open(),
        block: is_open(),
      }}
    >
      <div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div>
      <div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 bottom-0 right-0 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto">
        <div id="dialog-content" class="p-8 flex flex-col space-y-2 w-full">
          <div id="info-dialog">
            <div class="flex flex-col space-y-2">
              <div class="flex justify-between items-center">
                <div class="text-3xl">Warning</div>
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
              <div class="flex flex-col space-y-2">
                <div class="text-xl">
                  You are going to view someone else's board and you haven't
                  played yet today. Play today's mural{" "}
                  <a
                    class="text-underline text-contessa-500"
                    href={
                      import.meta.env.VITE_SHARE_URL ?? "mural.andrewnathan.net"
                    }
                  >
                    here
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}