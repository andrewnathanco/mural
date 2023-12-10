import { useHintDialog } from "../dialog/hint/context";

export default function HintButton() {
  const [_, { open }] = useHintDialog();
  return (
    <button
      onclick={() => {
        window.scrollTo(0, 0);
        open();
      }}
      id="hint-button"
      class="w-full p-1 text-base text-river-bed-700 rounded-md flex justify-center items-center hover:cursor-pointer bg-transparent hover:bg-desert-sand-300 border-river-bed-700 border-2"
    >
      Hints
    </button>
  );
}
