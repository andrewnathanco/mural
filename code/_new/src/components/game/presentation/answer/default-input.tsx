export default function DefaultInput(props: { disabled: boolean }) {
  const disabled = props.disabled;

  return (
    <div
      classList={{ "hover:cursor-pointer": !disabled }}
      class=" text-lg w-full rounded-lg py-2 px-4 bg-transparent text-river-bed-700 border-2 border-river-bed-700 "
    >
      Movie (2022)
    </div>
  );
}
