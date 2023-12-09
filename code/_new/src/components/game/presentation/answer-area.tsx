import SubmitButton from "../../buttons/answer-button";
import DefaultInput from "./answer/default-input";

export default function AnswerArea() {
  return (
    <main class="w-full text-river-bed-600 font-extrabold flex flex-col space-y-4">
      <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
      <ul class="w-full flex flex-col space-y-1">
        <DefaultInput disabled={false} />
        <DefaultInput disabled={false} />
        <DefaultInput disabled={false} />
        <DefaultInput disabled={false} />
      </ul>
      <div>
        <SubmitButton disabled={false} />
      </div>
    </main>
  );
}
