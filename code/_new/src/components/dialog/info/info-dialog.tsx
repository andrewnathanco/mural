import { createEffect, createSignal } from "solid-js";
import { useInfoDialog } from "./context";
import { get_countdown_till_next_game } from "../../game/service";

export default function InfoDialog() {
  const [is_open, { open, close }] = useInfoDialog();
  const [countdown, set_countdown] = createSignal(
    get_countdown_till_next_game()
  );

  createEffect(() => {
    setInterval(() => {
      set_countdown(get_countdown_till_next_game());
    }, 1000);
  });

  return (
    <div classList={{ hidden: !is_open(), block: is_open() }}>
      <div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div>
      <div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 bottom-0 right-0 md:w-128 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto">
        <div id="dialog-content" class="p-8 flex flex-col space-y-2 w-full">
          <div id="info-dialog">
            <div class="flex flex-col space-y-2">
              <div class="flex justify-between items-center">
                <div class="text-3xl">Mural</div>
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
                <div class="text-xl">What is Mural</div>
                <div class="text-md">
                  Mural is a daily puzzle game where you have to guess a movie
                  poster by flipping over tiles. The puzzle refreshes daily at
                  11:59 EST.
                </div>
                <div>
                  Every day has a different theme. On normal weeks the themes go
                  by decade.
                </div>
                <ul>
                  <li>
                    <strong>Monday</strong>: 2020s
                  </li>
                  <li>
                    <strong>Tuesday</strong>: 2010s
                  </li>
                  <li>
                    <strong>Wednesday</strong>: 2000s
                  </li>
                  <li>
                    <strong>Thursday</strong>: 1990s
                  </li>
                  <li>
                    <strong>Friday</strong>: 1980s
                  </li>
                  <li>
                    <strong>Saturday</strong>: 1970s
                  </li>
                  <li>
                    <strong>Sunday</strong>: Random
                  </li>
                </ul>
                <div class="text-xl">How to Play</div>
                <div>
                  Flip over tiles one at a time. Each tile has a penalty. The
                  outer tiles have the lowest penalty and the inner ones the
                  highest.
                </div>
                <div class="text-xl">Penalty Distribution</div>
              </div>
              <div class="flex flex-col space-y-2">
                <div class="text-xl">Ways to Support</div>
                <div class="text-md">
                  <div>
                    I make these projects in my free time and so any support
                    would be great. If you liked the project and want to help
                    with building it, check out the{" "}
                    <a
                      target="_blank"
                      class="text-contessa-600 underline hover:underline-offset-2 hover:cursor-pointer"
                      href="https://github.com/andrewnathanco/mural"
                    >
                      source code
                    </a>
                  </div>
                  <br />
                  <div>
                    If you like what I do here and want to check out other
                    projects, you can checkout my{" "}
                    <a
                      target="_blank"
                      class="text-contessa-600 underline hover:underline-offset-2 hover:cursor-pointer"
                      href="https://andrewnathan.net"
                    >
                      portfolio site
                    </a>{" "}
                    or you can check out the community focused startup I'm
                    working on called{" "}
                    <a
                      target="_blank"
                      class="text-contessa-600 underline hover:underline-offset-2 hover:cursor-pointer"
                      href="https://joinpickup.com"
                    >
                      Pickup
                    </a>
                    . We are building a platform to help you connect with people
                    in your local community.
                  </div>
                  <br />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
