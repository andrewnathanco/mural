import { Accessor, createContext, createSignal, useContext } from "solid-js";

type HintDialog = [
  Accessor<boolean>,
  {
    close: () => void;
    open: () => void;
  }
];

const HintDialogContext = createContext<HintDialog>();

export function HintDialogProvider(props: { children: any }) {
  const [dialog_status, set_dialog_status] = createSignal(false);
  const dialog: HintDialog = [
    dialog_status,
    {
      close() {
        set_dialog_status(false);
        document.body.style.overflowY = "auto";
      },
      open() {
        set_dialog_status(true);
        document.body.style.position = "relative";
        document.body.style.overflowY = "hidden";
      },
    },
  ];

  return (
    <HintDialogContext.Provider value={dialog}>
      {props.children}
    </HintDialogContext.Provider>
  );
}

export function useHintDialog(): HintDialog {
  return useContext(HintDialogContext) as HintDialog;
}
