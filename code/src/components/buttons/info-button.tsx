import { useInfoDialog } from "../dialog/info/context";

export default function InfoButton() {
  const [_, { open }] = useInfoDialog();
  return (
    <button
      onclick={() => {
        window.scrollTo(0, 0);
        open();
      }}
      id="info-button"
      class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 "
    >
      Info
    </button>
  );
}
