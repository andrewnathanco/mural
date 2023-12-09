export default function SubmitButton(props: { disabled: boolean }) {
  const disabled = props.disabled;

  return (
    <div
      classList={{
        "bg-river-bed-400": disabled,
        "border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700":
          !disabled,
      }}
      class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center "
    >
      Submit
    </div>
  );
}
